package runner

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	_ "embed"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/bhakiyakalimuthu/backrunner/runner/abi/executor"
	"github.com/bhakiyakalimuthu/backrunner/runner/abi/isushipair"
	"github.com/bhakiyakalimuthu/backrunner/runner/abi/iunipool"
	sushiswapfactory "github.com/bhakiyakalimuthu/backrunner/runner/abi/sushifactory"
	uniswapfactory "github.com/bhakiyakalimuthu/backrunner/runner/abi/unifactory"
	btypes "github.com/bhakiyakalimuthu/backrunner/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/metachris/flashbotsrpc"
	"go.uber.org/zap"
)

var (
	_UniSwapRouter02 = common.HexToAddress("0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45")
	_WETHAddress     = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")

	//go:embed abi/iunipool/iuniswapv3pool.json
	uniswapRouter02Bytes []byte
)

const (
	scMethodSizeInBytes       = 4 // first 4 bytes of data field
	scMethodMultiCall         = "0x5ae401dc"
	scMethodExactInputSingle  = "0x04e45aaf"
	scMethodExactOutputSingle = "0x5023b4df"
)

type Runner struct {
	logger                *zap.Logger
	clients               *Clients
	contracts             *Contracts
	senderKey, bundlerKey *ecdsa.PrivateKey
}

type Contracts struct {
	UniswapFactory   *uniswapfactory.Uniswapfactory
	SushiswapFactory *sushiswapfactory.Sushiswapfactory
	Executor         *executor.Executor
}

type Clients struct {
	BaseClient *rpc.Client
	Subscriber *gethclient.Client
	EthClient  *ethclient.Client
	FbClient   *flashbotsrpc.FlashbotsRPC
}

func NewRunner(logger *zap.Logger, clients *Clients, contracts *Contracts, senderKey, bundlerKey *ecdsa.PrivateKey) *Runner {
	return &Runner{
		logger:     logger,
		clients:    clients,
		contracts:  contracts,
		senderKey:  senderKey,
		bundlerKey: bundlerKey,
	}
}

func (r *Runner) MonitorMempool(ctx context.Context) error {
	pendingTxs := make(chan common.Hash)
	subscription, err := r.clients.Subscriber.SubscribePendingTransactions(ctx, pendingTxs)
	if err != nil {
		return fmt.Errorf("failed to subscribe pending txs: %v", err)
	}

	for {
		select {
		case txHash := <-pendingTxs:
			tx, isPending, err := r.clients.EthClient.TransactionByHash(ctx, txHash)
			r.logger.Info("received pending tx", zap.String("txHash", txHash.Hex()))
			if !isPending || err != nil {
				continue
			}
			if tx.To().Hash() == _UniSwapRouter02.Hash() {
				r.logger.Info("=============================")
				r.logger.Info("Found uniswap trade", zap.String("txHash", txHash.Hex()))
				r.logger.Info("=============================")
				trade, err := r.decodeTxFinal(ctx, tx)
				if err != nil {
					r.logger.Error("decode transaction failed", zap.Error(err), zap.String("txHash", txHash.Hex()))
					continue
				}

				receipt, err := r.clients.EthClient.TransactionReceipt(ctx, txHash)
				if err != nil {
					r.logger.Error("failed to check transaction receipt", zap.Error(err), zap.String("txHash", txHash.Hex()))
					continue
				}
				if trade != nil && receipt != nil && (strings.ToLower(trade.TokenOut) == strings.ToLower(_WETHAddress.String())) {
					r.executeBackrun(ctx, trade, tx)
				}
			}
		case err := <-subscription.Err():
			return fmt.Errorf("error occured in pending tx subscription: %v", err)
		case <-ctx.Done():
			close(pendingTxs)
			return nil
		}
	}
}
func (r *Runner) executeBackrun(ctx context.Context, trade *btypes.UniswapTrade, mempoolTx *types.Transaction) error {
	uniPoolAddress, sushiPoolAddress, err := r.getPoolAddresses(trade)
	if err != nil {
		return err
	}
	unipool, err := iunipool.NewIunipool(*uniPoolAddress, r.clients.EthClient)
	if err != nil {
		return fmt.Errorf("failed to create unipool instance %v", err)
	}
	uniToken0, err := unipool.Token0(&bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("failed to get unipool token0 %v", err)

	}
	sushipair, err := isushipair.NewIsushipair(*sushiPoolAddress, r.clients.EthClient)
	if err != nil {
		return fmt.Errorf("failed to create sushipair instance %v", err)
	}
	sushiToken0, err := sushipair.Token0(&bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("failed to get sushipool token0 %v", err)

	}
	buyAmount := trade.AmountInMax
	if buyAmount == nil {
		buyAmount = trade.AmountIn
	}
	auth, err := bind.NewKeyedTransactorWithChainID(r.senderKey, big.NewInt(1))
	if err != nil {
		return fmt.Errorf("failed to create sender auth %v", err)

	}
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(800000)
	auth.GasPrice = big.NewInt(42000000000) // maxFeePerGas
	auth.Context = context.Background()
	suggestedGasPrice, err := r.clients.EthClient.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println("error getting gas price", err)
		return fmt.Errorf("failed to get suggestedGasPrice %v", err)
	}
	suggestedTip, err := r.clients.EthClient.SuggestGasTipCap(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get suggestedTip %v", err)
	}
	nonce, err := r.clients.EthClient.PendingNonceAt(ctx, crypto.PubkeyToAddress(r.senderKey.PublicKey))
	if err != nil {
		log.Error("failed to get pending nonce", zap.Error(err))
		return nil
	}
	gasPrice := suggestedGasPrice.Mul(suggestedGasPrice, big.NewInt(1.1))
	gasTipCap := suggestedTip.Mul(suggestedTip, big.NewInt(1.1))
	gasFeeCap := new(big.Int).Add(gasPrice, gasTipCap)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasFeeCap = gasFeeCap

	tx, err := r.contracts.Executor.Execute(auth,
		*sushiPoolAddress,
		*uniPoolAddress,
		common.HexToAddress(trade.TokenIn),
		sushiToken0 == common.HexToAddress(trade.TokenIn),
		uniToken0 == _WETHAddress,
		buyAmount,
		trade.SqrtPriceLimitX96,
	)
	if err != nil {
		return fmt.Errorf("failed to execute backrun transaction %v", err)
	}
	config, block := params.MainnetChainConfig, params.MainnetChainConfig.LondonBlock
	signer := types.MakeSigner(config, block)
	signedBackrunTx, err := types.SignTx(tx, signer, r.senderKey)
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
	currentBlock, err := r.clients.EthClient.BlockNumber(ctx)
	if err != nil {
		return fmt.Errorf("failed to get block number %v", err)
	}
	opts := flashbotsrpc.FlashbotsCallBundleParam{
		Txs:              []string{hexutil.Encode(rawMempoolTx), hexutil.Encode(rawBackrunTx)},
		BlockNumber:      fmt.Sprintf("0x%x", currentBlock),
		StateBlockNumber: "latest",
	}
	// simulate bundle
	result, err := r.clients.FbClient.FlashbotsCallBundle(r.bundlerKey, opts)
	if err != nil {
		return err
	}
	r.logger.Info("bundle simulation succeeded", zap.Any("simResult", result))
	// send bundle
	return nil

}

func (r *Runner) getPoolAddresses(trade *btypes.UniswapTrade) (*common.Address, *common.Address, error) {
	uniPoolAddress, err := r.contracts.UniswapFactory.GetPool(&bind.CallOpts{}, common.HexToAddress(trade.TokenIn), common.HexToAddress(trade.TokenIn), trade.Fee)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get uniswap  pool %v", zap.Error(err))
	}
	sushiPoolAddress, err := r.contracts.SushiswapFactory.GetPair(&bind.CallOpts{}, common.HexToAddress(trade.TokenIn), common.HexToAddress(trade.TokenIn))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get sushi pool %v", zap.Error(err))
	}
	return &uniPoolAddress, &sushiPoolAddress, nil
}

func (r *Runner) decodeTx(ctx context.Context, tx *types.Transaction) error {

	if !(len(tx.Data()) >= scMethodSizeInBytes) {
		return fmt.Errorf("unknown transaction data len")
	}
	scMethod := tx.Data()[:scMethodSizeInBytes]
	r.logger.Info("method from data", zap.String("method", string(scMethod)))
	fmt.Println("method from data ", hexutil.Encode(scMethod))
	abi, err := abi.JSON(bytes.NewReader(uniswapRouter02Bytes))
	if err != nil {
		return err
	}

	method, err := abi.MethodById(scMethod)
	if err != nil {
		return err
	}
	inputsSigData := tx.Data()[scMethodSizeInBytes:]

	//err = abi.Unpack(tx.Data(), "swapExactETHForTokens", tx.Data())
	//if err != nil {
	//	panic(err)
	//}
	r.logger.Info("method from data", zap.String("method", string(inputsSigData)))
	fmt.Println("method from data ", hexutil.Encode(inputsSigData))
	inputsMap := make(map[string]interface{})
	if err := method.Inputs.UnpackIntoMap(inputsMap, inputsSigData); err != nil {
		return err
	}
	r.logger.Info("inputs map", zap.Any("inputmap", inputsMap))
	fmt.Println("inputs map", inputsMap)

	val := inputsMap["data"]
	sVal := val.([][]byte)
	scMethod2 := sVal[0][:scMethodSizeInBytes]
	fmt.Println("method2 from data ", hexutil.Encode(sVal[0][:scMethodSizeInBytes]))
	fmt.Println("method2 from data ", hexutil.Encode(sVal[0][scMethodSizeInBytes:]))

	method2, err := abi.MethodById(scMethod2)
	if err != nil {
		return err
	}
	inputsSigData2 := sVal[0][scMethodSizeInBytes:]

	//err = abi.Unpack(tx.Data(), "swapExactETHForTokens", tx.Data())
	//if err != nil {
	//	panic(err)
	//}
	r.logger.Info("method from data", zap.String("method", string(inputsSigData2)))
	fmt.Println("method from data ", hexutil.Encode(inputsSigData2))
	inputsMap2 := make(map[string]interface{})
	if err := method2.Inputs.UnpackIntoMap(inputsMap2, inputsSigData2); err != nil {
		return err
	}
	r.logger.Info("inputs map", zap.Any("inputmap2", inputsMap2))
	fmt.Println("inputs map2", inputsMap2)

	str := mapToJson(inputsMap)
	r.logger.Info("input fields", zap.String("map", str))
	fmt.Println("inputs map", str)
	return nil
}

func (r *Runner) decodeTxFinal(ctx context.Context, tx *types.Transaction) (*btypes.UniswapTrade, error) {

	if !(len(tx.Data()) >= scMethodSizeInBytes) {
		return nil, fmt.Errorf("unknown transaction data len")
	}
	abi, err := abi.JSON(bytes.NewReader(uniswapRouter02Bytes))
	if err != nil {
		return nil, err
	}
	scMethod := tx.Data()[:scMethodSizeInBytes]

	if hexutil.Encode(scMethod) == scMethodMultiCall {
		multicall, err := decodeTxDataField(abi, scMethod, tx.Data())
		if err != nil {
			return nil, err
		}
		deadline := multicall["deadline"]
		data := multicall["data"].([][]byte)
		swapcallMethod := hexutil.Encode(data[0][:scMethodSizeInBytes])
		swapcall, err := decodeTxDataField(abi, data[0][:scMethodSizeInBytes], data[0][scMethodSizeInBytes:])
		if err != nil {
			return nil, err
		}
		if swapcallMethod == scMethodExactInputSingle {
			return &btypes.UniswapTrade{
				Deadline:          deadline.(*big.Int),
				TokenIn:           swapcall["tokenIn"].(string),
				TokenOut:          swapcall["tokenOut"].(string),
				Fee:               swapcall["fee"].(*big.Int),
				Recipient:         swapcall["recipient"].(string),
				AmountIn:          swapcall["amountIn"].(*big.Int),
				AmountOutMin:      swapcall["amountOutMinimum"].(*big.Int),
				SqrtPriceLimitX96: swapcall["sqrtPriceLimitX96"].(*big.Int),
			}, nil
		}
		if swapcallMethod == scMethodExactOutputSingle {
			return &btypes.UniswapTrade{
				Deadline:          deadline.(*big.Int),
				TokenIn:           swapcall["tokenIn"].(string),
				TokenOut:          swapcall["tokenOut"].(string),
				Fee:               swapcall["fee"].(*big.Int),
				Recipient:         swapcall["recipient"].(string),
				AmountOut:         swapcall["amountOut"].(*big.Int),
				AmountInMax:       swapcall["amountInMaximum"].(*big.Int),
				SqrtPriceLimitX96: swapcall["sqrtPriceLimitX96"].(*big.Int),
			}, nil
		}
	}

	return nil, nil
}

func mapToJson(params map[string]interface{}) string {
	data, _ := json.Marshal(params)
	dataStr := string(data)
	return dataStr
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
