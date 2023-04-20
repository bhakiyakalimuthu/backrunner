package runner

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/bhakiyakalimuthu/backrunner/runner/abi/isushipair"
	"github.com/bhakiyakalimuthu/backrunner/runner/abi/iunipool"
	btypes "github.com/bhakiyakalimuthu/backrunner/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/metachris/flashbotsrpc"
	"go.uber.org/zap"
)

var _ Trade = (*UniTrade)(nil)

// UniTrade extends the Trade interface to perform backrun on Uniswap trades
type UniTrade struct {
	logger                *zap.Logger
	clients               *Clients
	contracts             *Contracts
	senderKey, bundlerKey *ecdsa.PrivateKey
}

func NewUniTrade(logger *zap.Logger, clients *Clients, contracts *Contracts, senderKey, bundlerKey *ecdsa.PrivateKey) *UniTrade {
	return &UniTrade{
		logger:     logger,
		clients:    clients,
		contracts:  contracts,
		senderKey:  senderKey,
		bundlerKey: bundlerKey,
	}
}

func (u *UniTrade) ExecuteBackrun(ctx context.Context, txHash common.Hash) {
	tx, isPending, err := u.clients.EthClient.TransactionByHash(ctx, txHash)
	u.logger.Debug("received pending tx", zap.String("txHash", txHash.Hex()))
	if !isPending || err != nil {
		return
	}
	if tx.To() == nil {
		return // omit contract creation
	}

	// Back run only UniSwapRouter02 trades for simplicity
	if tx.To().Hash() == _UniSwapRouter02.Hash() {
		u.logger.Debug("Found uniswap trade", zap.String("txHash", txHash.Hex()))
		trade, err := u.decodeTx(tx)
		if err != nil {
			u.logger.Error("decode transaction failed", zap.Error(err), zap.String("txHash", txHash.Hex()))
			return
		}

		// Validate whether transaction landed on-chain
		receipt, err := u.clients.EthClient.TransactionReceipt(ctx, txHash)
		if err != nil {
			u.logger.Error("failed to check transaction receipt", zap.Error(err), zap.String("txHash", txHash.Hex()))
			return
		}
		if trade != nil && receipt != nil && (trade.TradeParams.Params.TokenOut) == _WETHAddress.String() {
			// user is trading X tokens for WETH
			// buy X tokens
			_ = u.execute(ctx, trade, tx)
		}
	}
}

func (u *UniTrade) execute(ctx context.Context, trade *btypes.UniswapTrade, mempoolTx *types.Transaction) error {
	// Get pool address from each exchanges
	uniPoolAddress, sushiPoolAddress, err := u.getPoolAddresses(trade)
	if err != nil {
		return err
	}

	unipool, err := iunipool.NewIunipool(*uniPoolAddress, u.clients.EthClient)
	if err != nil {
		return fmt.Errorf("failed to create unipool instance %v", err)
	}
	sushipair, err := isushipair.NewIsushipair(*sushiPoolAddress, u.clients.EthClient)
	if err != nil {
		return fmt.Errorf("failed to create sushipair instance %v", err)
	}

	uniToken0, err := unipool.Token0(&bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("failed to get unipool token0 %v", err)
	}
	sushiToken0, err := sushipair.Token0(&bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("failed to get sushipool token0 %v", err)
	}

	buyAmount := trade.TradeParams.Params.AmountIn // exactInputSingle
	if buyAmount == nil {
		buyAmount = trade.TradeParams.Params.AmountInMaximum // ExactOutputSingle
	}
	auth, err := bind.NewKeyedTransactorWithChainID(u.senderKey, big.NewInt(1))
	if err != nil {
		return fmt.Errorf("failed to create sender auth %v", err)
	}

	// build transaction meta data
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(800000)
	auth.GasPrice = big.NewInt(42000000000) // maxFeePerGas
	auth.Context = context.Background()
	gasFeeCap, err := u.clients.EthClient.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get suggestedGasPrice %v", err)
	}
	suggestedTip, err := u.clients.EthClient.SuggestGasTipCap(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get suggestedTip %v", err)
	}
	nonce, err := u.clients.EthClient.PendingNonceAt(ctx, crypto.PubkeyToAddress(u.senderKey.PublicKey))
	if err != nil {
		return fmt.Errorf("failed to get pending nonce %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasFeeCap = gasFeeCap
	auth.GasTipCap = suggestedTip

	tx, err := u.contracts.Executor.Execute(auth,
		*sushiPoolAddress,
		*uniPoolAddress,
		common.HexToAddress(trade.Params.TokenIn),
		sushiToken0 == common.HexToAddress(trade.Params.TokenIn), // will always be tokenIn bcz we filter by (tokenOut == WETH)
		uniToken0 == _WETHAddress,
		buyAmount,
		trade.Params.SqrtPriceLimitX96,
	)
	if err != nil {
		return fmt.Errorf("failed to execute contract executor transaction %v", err)
	}

	config, block := params.MainnetChainConfig, params.MainnetChainConfig.LondonBlock
	signer := types.MakeSigner(config, block)
	signedBackrunTx, err := types.SignTx(tx, signer, u.senderKey)
	if err != nil {
		return fmt.Errorf("failed to sign transaction %v", err)
	}

	rawMempoolTx, err := mempoolTx.MarshalBinary()
	if err != nil {
		return fmt.Errorf("failed to marshal tx %v", err)
	}
	rawBackrunTx, err := signedBackrunTx.MarshalBinary()
	if err != nil {
		return fmt.Errorf("failed to marshal tx %v", err)
	}
	currentBlock, err := u.clients.EthClient.BlockNumber(ctx)
	if err != nil {
		return fmt.Errorf("failed to get block number %v", err)
	}

	// simulate bundle via Flashsbots
	callBundleOpts := flashbotsrpc.FlashbotsCallBundleParam{
		Txs:              []string{hexutil.Encode(rawMempoolTx), hexutil.Encode(rawBackrunTx)},
		BlockNumber:      fmt.Sprintf("0x%x", currentBlock+1),
		StateBlockNumber: "latest",
	}

	result, err := u.clients.FbClient.FlashbotsCallBundle(u.bundlerKey, callBundleOpts)
	if err != nil {
		return err
	}
	u.logger.Info("bundle simulation succeeded", zap.Any("simResult", result))
	// TODO:determine the profit to decide whether to send onchain or not

	// send bundle via Flashsbots
	sendBundleOpts := flashbotsrpc.FlashbotsSendBundleRequest{
		Txs:         []string{hexutil.Encode(rawMempoolTx), hexutil.Encode(rawBackrunTx)},
		BlockNumber: fmt.Sprintf("0x%x", currentBlock+1),
	}
	_result, err := u.clients.FbClient.FlashbotsSendBundle(u.bundlerKey, sendBundleOpts)
	if err != nil {
		return err
	}
	u.logger.Info("send bundle  succeeded", zap.Any("sendBundleResult", _result))
	return nil
}

func (u *UniTrade) getPoolAddresses(trade *btypes.UniswapTrade) (*common.Address, *common.Address, error) {
	uniPoolAddress, err := u.contracts.UniswapFactory.GetPool(&bind.CallOpts{}, common.HexToAddress(trade.TradeParams.Params.TokenIn), common.HexToAddress(trade.TradeParams.Params.TokenOut), trade.Params.Fee)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get uniswap  pool %v", zap.Error(err))
	}
	sushiPoolAddress, err := u.contracts.SushiswapFactory.GetPair(&bind.CallOpts{}, common.HexToAddress(trade.TradeParams.Params.TokenIn), common.HexToAddress(trade.TradeParams.Params.TokenOut))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get sushi pool %v", zap.Error(err))
	}
	return &uniPoolAddress, &sushiPoolAddress, nil
}

func (u *UniTrade) decodeTx(tx *types.Transaction) (*btypes.UniswapTrade, error) {
	if !(len(tx.Data()) >= scMethodSizeInBytes) {
		return nil, fmt.Errorf("unknown transaction data len")
	}
	abi, err := abi.JSON(bytes.NewReader(uniswapRouter02Bytes))
	if err != nil {
		return nil, err
	}
	scMethod := tx.Data()[:scMethodSizeInBytes]

	if hexutil.Encode(scMethod) == scMethodMultiCall {
		// found multicall trade
		multicall, err := decodeTxDataField(abi, scMethod, tx.Data())
		if err != nil {
			return nil, err
		}
		deadline := multicall["deadline"]
		data := multicall["data"].([][]byte)
		swapcallMethod := hexutil.Encode(data[0][:scMethodSizeInBytes])
		swapcall, err := decodeTxDataField(abi, data[0][:scMethodSizeInBytes], data[0])
		if err != nil {
			return nil, err
		}

		switch swapcallMethod {
		case scMethodSwapExactTokensForTokens, scMethodSwapTokensForExactTokens:
			log.Warn("swapExactTokensForTokens & swapTokensForExactTokens unimplemented")
			return nil, nil
		case scMethodExactInputSingle, scMethodExactOutputSingle:
			return swapcallToTrade(swapcall, deadline)
		}
	}

	return nil, nil
}

func swapcallToTrade(swapcall map[string]interface{}, deadline interface{}) (*btypes.UniswapTrade, error) {
	var _deadline time.Duration
	f, err := strconv.ParseFloat(deadline.(*big.Int).String(), 64)
	if err == nil {
		_deadline = time.Duration(f * float64(time.Second))
	}
	params, err := mapToJson(swapcall)
	if err != nil {
		return nil, err
	}
	params.Params.TokenIn = strings.ToLower(params.Params.TokenIn)
	params.Params.TokenOut = strings.ToLower(params.Params.TokenOut)
	trade := &btypes.UniswapTrade{Deadline: _deadline, TradeParams: params}
	return trade, nil
}

func mapToJson(params map[string]interface{}) (*btypes.TradeParams, error) {
	var p *btypes.TradeParams
	data, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal trade params %v", err)
	}
	if err := json.Unmarshal(data, &p); err != nil {
		return nil, fmt.Errorf("failed to unmarshal trade params %v", err)
	}
	return p, nil
}

func decodeTxDataField(abi abi.ABI, scMethodBytes, data []byte) (map[string]interface{}, error) {
	method, err := abi.MethodById(scMethodBytes)
	if err != nil {
		return nil, err
	}
	inputsSigData := data[scMethodSizeInBytes:]
	inputsMap := make(map[string]interface{})
	if err := method.Inputs.UnpackIntoMap(inputsMap, inputsSigData); err != nil {
		return nil, err
	}
	return inputsMap, nil
}
