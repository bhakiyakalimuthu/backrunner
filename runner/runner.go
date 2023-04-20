package runner

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/bhakiyakalimuthu/backrunner/runner/abi/executor"
	sushiswapfactory "github.com/bhakiyakalimuthu/backrunner/runner/abi/sushifactory"
	uniswapfactory "github.com/bhakiyakalimuthu/backrunner/runner/abi/unifactory"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/metachris/flashbotsrpc"
	"go.uber.org/zap"
)

var (
	_UniSwapRouter02 = common.HexToAddress("0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45")
	_WETHAddress     = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
	_UnknownAddress  = common.HexToAddress("0x0000000000000000000000000000000000000000")
	//go:embed abi/unifactory/uniswaprouter02.json
	uniswapRouter02Bytes []byte
)

const (
	scMethodSizeInBytes              = 4 // first 4 bytes of data field
	scMethodMultiCall                = "0x5ae401dc"
	scMethodExactInputSingle         = "0x04e45aaf"
	scMethodExactOutputSingle        = "0x5023b4df"
	scMethodSwapExactTokensForTokens = "0x472b43f3"
	scMethodSwapTokensForExactTokens = "0x42712a67"
)

type Runner struct {
	logger  *zap.Logger
	clients *Clients
	trade   Trade
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

func NewRunner(logger *zap.Logger, clients *Clients, trade Trade) *Runner {
	return &Runner{
		logger:  logger,
		clients: clients,
		trade:   trade,
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
			go func() {
				r.trade.ExecuteBackrun(ctx, txHash)
			}()
		case err := <-subscription.Err():
			return fmt.Errorf("error occured in pending tx subscription: %v", err)
		case <-ctx.Done():
			close(pendingTxs)
			return nil
		}
	}
}
