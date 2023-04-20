package runner

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/bhakiyakalimuthu/backrunner/runner/abi/executor"
	sushiswapfactory "github.com/bhakiyakalimuthu/backrunner/runner/abi/sushifactory"
	uniswapfactory "github.com/bhakiyakalimuthu/backrunner/runner/abi/unifactory"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-playground/assert/v2"
	"github.com/metachris/flashbotsrpc"
	"go.uber.org/zap/zaptest"
)

var (
	_SushiswapFactory = common.HexToAddress("0xc0aee478e3658e2610c5f7a4a2e1777ce9e4f2ac")
	_UniswapV3Factory = common.HexToAddress("0x1f98431c8ad98523631ae4a59f267346ea31f984")
)

func TestUniTrade_Decode(t *testing.T) {
	ctx := context.Background()
	ethClient, err := ethclient.DialContext(ctx, "wss://mainnet.infura.io/ws/v3/77582faac1cc4bd7be282243cb13afa1")
	if err != nil {
		t.Errorf("failed to dial eth client %v", err)
		return
	}
	fbClient := flashbotsrpc.New("https://relay.flashbots.net")

	uniFactory, err := uniswapfactory.NewUniswapfactory(_UniswapV3Factory, ethClient)
	if err != nil {
		t.Errorf("failed to create unipool instance %v", err)
		return
	}
	sushiFactory, err := sushiswapfactory.NewSushiswapfactory(_SushiswapFactory, ethClient)
	if err != nil {
		t.Errorf("failed to create unipool instance %v", err)
		return
	}
	executor, err := executor.NewExecutor(common.HexToAddress("0x4951a1c579039ebfcba0be33d2cd3a6d30b0f802"), ethClient)
	if err != nil {
		t.Errorf("failed to create unipool instance %v", err)
		return
	}

	clients := &Clients{
		EthClient: ethClient,
		FbClient:  fbClient,
	}
	contracts := &Contracts{
		UniswapFactory:   uniFactory,
		SushiswapFactory: sushiFactory,
		Executor:         executor,
	}
	bundlerKey, senderKey := createSigningKeys()
	u := &UniTrade{
		logger:     zaptest.NewLogger(t),
		clients:    clients,
		contracts:  contracts,
		senderKey:  senderKey,
		bundlerKey: bundlerKey,
	}
	// txHash := common.HexToHash("0xe1a8229ba40433afd22bfe766c028cbb9a23930ea705cb54022d0fc0da79926d")
	txHash := common.HexToHash("0x6b059739e1c53523982a687f1b998103377a72e0a65ea0598ceb177ade022da7")
	tx, _, err := ethClient.TransactionByHash(ctx, txHash)
	if err != nil {
		t.Errorf("failed to TransactionByHash %v", err)
		return
	}
	trade, err := u.decodeTx(ctx, tx)
	if err != nil {
		t.Errorf("failed to create unipool instance %v", err)
		return
	}

	assert.Equal(t, "0x626e8036deb333b408be468f951bdb42433cbf18", trade.TradeParams.Params.TokenIn)
	assert.Equal(t, "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2", trade.TradeParams.Params.TokenOut)
	assert.Equal(t, big.NewInt(10000), trade.TradeParams.Params.Fee)
	expectedInMax, _ := new(big.Int).SetString("1025556186883934588453", 0)
	assert.Equal(t, expectedInMax, trade.TradeParams.Params.AmountInMaximum)
}

func TestUniTrade_ExecuteBackrun(t *testing.T) {
	ctx := context.Background()
	ethClient, err := ethclient.DialContext(ctx, "wss://mainnet.infura.io/ws/v3/77582faac1cc4bd7be282243cb13afa1")
	if err != nil {
		t.Errorf("failed to dial eth client %v", err)
		return
	}
	fbClient := flashbotsrpc.New("https://relay.flashbots.net")

	uniFactory, err := uniswapfactory.NewUniswapfactory(_UniswapV3Factory, ethClient)
	if err != nil {
		t.Errorf("failed to create unipool instance %v", err)
		return
	}
	sushiFactory, err := sushiswapfactory.NewSushiswapfactory(_SushiswapFactory, ethClient)
	if err != nil {
		t.Errorf("failed to create unipool instance %v", err)
		return
	}
	executor, err := executor.NewExecutor(common.HexToAddress("0x4951a1c579039ebfcba0be33d2cd3a6d30b0f802"), ethClient)
	if err != nil {
		t.Errorf("failed to create unipool instance %v", err)
		return
	}
	clients := &Clients{
		EthClient: ethClient,
		FbClient:  fbClient,
	}
	contracts := &Contracts{
		UniswapFactory:   uniFactory,
		SushiswapFactory: sushiFactory,
		Executor:         executor,
	}
	bundlerKey, senderKey := createSigningKeys()
	u := &UniTrade{
		logger:     zaptest.NewLogger(t),
		clients:    clients,
		contracts:  contracts,
		senderKey:  senderKey,
		bundlerKey: bundlerKey,
	}
	// txHash := common.HexToHash("0xe1a8229ba40433afd22bfe766c028cbb9a23930ea705cb54022d0fc0da79926d")
	txHash := common.HexToHash("0x6b059739e1c53523982a687f1b998103377a72e0a65ea0598ceb177ade022da7")
	tx, _, err := ethClient.TransactionByHash(ctx, txHash)
	if err != nil {
		t.Errorf("failed to TransactionByHash %v", err)
		return
	}
	trade, err := u.decodeTx(ctx, tx)
	if err != nil {
		t.Errorf("failed to create unipool instance %v", err)
		return
	}
	_ = u.execute(ctx, trade, tx)
}

func createSigningKeys() (bundleSigningKey, senderSigningKey *ecdsa.PrivateKey) {
	bundleSigningKey, _ = crypto.GenerateKey()
	senderSigningKey, _ = crypto.GenerateKey()
	return
}
