package main

import (
	"context"
	"crypto/ecdsa"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bhakiyakalimuthu/backrunner/config"
	"github.com/bhakiyakalimuthu/backrunner/runner"
	"github.com/bhakiyakalimuthu/backrunner/runner/abi/executor"
	sushiswapfactory "github.com/bhakiyakalimuthu/backrunner/runner/abi/sushifactory"
	uniswapfactory "github.com/bhakiyakalimuthu/backrunner/runner/abi/unifactory"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/metachris/flashbotsrpc"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// Included in the build process
	_BuildVersion string
	_AppName      string

	_SushiswapFactory = common.HexToAddress("0xc0aee478e3658e2610c5f7a4a2e1777ce9e4f2ac")
	_UniswapV3Factory = common.HexToAddress("0x1f98431c8ad98523631ae4a59f267346ea31f984")
	_BackrunExecutor  = common.HexToAddress("0x0000000000000000000000000000000000000000")
)

func main() {
	cfg := config.NewConfig()
	l := newLogger(_AppName, _BuildVersion)
	ctx, cancel := context.WithCancel(context.Background())

	// init clients
	ethClient, err := ethclient.DialContext(ctx, cfg.EthClientWebSocketURL)
	if err != nil {
		l.Fatal("failed to dial eth client", zap.Error(err))
	}

	baseClient, err := rpc.DialContext(ctx, cfg.EthClientWebSocketURL)
	if err != nil {
		l.Fatal("failed to dial rpc client", zap.Error(err))
	}
	subscriber := gethclient.New(baseClient)
	fbClient := flashbotsrpc.New(cfg.FlashbotsRelayURL)

	// init contracts
	uniFactory, err := uniswapfactory.NewUniswapfactory(_UniswapV3Factory, ethClient)
	if err != nil {
		l.Fatal("failed to create unipool instance", zap.Error(err))
	}
	sushiFactory, err := sushiswapfactory.NewSushiswapfactory(_SushiswapFactory, ethClient)
	if err != nil {
		l.Fatal("failed to create unipool instance", zap.Error(err))
	}
	executor, err := executor.NewExecutor(_BackrunExecutor, ethClient)
	if err != nil {
		l.Fatal("failed to create  executor instance", zap.Error(err))
	}
	clients := &runner.Clients{
		BaseClient: baseClient,
		Subscriber: subscriber,
		EthClient:  ethClient,
		FbClient:   fbClient,
	}
	contracts := &runner.Contracts{
		UniswapFactory:   uniFactory,
		SushiswapFactory: sushiFactory,
		Executor:         executor,
	}
	bundlerKey, senderKey := createSigningKeys(l, cfg)

	uniswapTrade := runner.NewUniTrade(l, clients, contracts, senderKey, bundlerKey)
	runner := runner.NewRunner(l, clients, uniswapTrade)

	exit := make(chan struct{})
	go func() {
		shutdown := make(chan os.Signal, 1)
		signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
		l.Warn("shutting down")
		<-shutdown
		signal.Stop(shutdown)
		cancel()
		close(exit)
	}()
	if err := runner.MonitorMempool(ctx); err != nil {
		l.Fatal("failed to monitor mempool", zap.Error(err))
	}
	<-exit
}

func newLogger(appName, version string) *zap.Logger {
	logLevel := zap.DebugLevel
	var zapCore zapcore.Core
	level := zap.NewAtomicLevel()
	level.SetLevel(logLevel)
	encoderCfg := zap.NewProductionEncoderConfig()
	encoder := zapcore.NewJSONEncoder(encoderCfg)
	zapCore = zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), level)

	logger := zap.New(zapCore, zap.AddCaller(), zap.ErrorOutput(zapcore.Lock(os.Stderr)))
	logger = logger.With(zap.String("app", appName), zap.String("_BuildVersion", version))
	return logger
}

func createSigningKeys(l *zap.Logger, cfg *config.Config) (bundleSigningKey, senderSigningKey *ecdsa.PrivateKey) {
	var err error
	if cfg.BundleSigningKey == "" {
		l.Fatal("Cannot use relay without a bundle signing key.")
	}
	if cfg.SenderSigningKey == "" {
		l.Fatal("Cannot use backrunner without sender signing key.")
	}
	bundleSigningKey, err = crypto.HexToECDSA(strings.TrimPrefix(cfg.BundleSigningKey, "0x"))
	if err != nil {
		l.Fatal("Error creating bundle signing key", zap.Error(err))
	}
	senderSigningKey, err = crypto.HexToECDSA(strings.TrimPrefix(cfg.SenderSigningKey, "0x"))
	if err != nil {
		l.Fatal("Error creating backrunner sender signing key", zap.Error(err))
	}
	return
}
