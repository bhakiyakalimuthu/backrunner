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
	"github.com/mattn/go-colorable"
	"github.com/metachris/flashbotsrpc"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// Included in the build process
	_BuildVersion string
	_AppName      string

	_SushiswapFactory = common.HexToAddress("0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac")
	_UniswapV3Factory = common.HexToAddress("0x1F98431c8aD98523631AE4a59f267346ea31F984")
	_BackrunExecutor  = common.HexToAddress("0x0000000000000000000000000000000000000000")
)

func main() {
	cfg := config.NewConfig()
	l := newLogger(_AppName, _BuildVersion)

	ctx, cancel := context.WithCancel(context.Background())

	// init clients
	ethClient, err := ethclient.DialContext(ctx, cfg.EthClientURL)
	if err != nil {
		l.Fatal("failed to dial eth client", zap.Error(err))
	}

	baseClient, err := rpc.DialContext(ctx, cfg.EthClientURL)
	if err != nil {
		l.Fatal("failed to dial rpc client", zap.Error(err))
	}
	subscriber := gethclient.New(baseClient)
	fbClient := flashbotsrpc.New(cfg.FlashbotsRelayURL)

	uniFactory, err := uniswapfactory.NewUniswapfactory(_UniswapV3Factory, ethClient)
	if err != nil {
		l.Fatal("failed to create unipool instance", zap.Error(err))
	}
	sushiFactory, err := sushiswapfactory.NewSushiswapfactory(_SushiswapFactory, ethClient)
	if err != nil {
		l.Fatal("failed to create unipool instance", zap.Error(err))
	}
	executor, err := executor.NewExecutor(_BackrunExecutor, ethClient)

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
	senderKey, bundlerKey := createSigningKeys(l, cfg)
	runner := runner.NewRunner(l, clients, contracts, senderKey, bundlerKey)

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
	if version == "dev" {
		_cfg := zap.NewDevelopmentEncoderConfig()
		_cfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
		zapCore = zapcore.NewCore(
			zapcore.NewConsoleEncoder(_cfg),
			zapcore.AddSync(colorable.NewColorableStdout()),
			zapcore.DebugLevel,
		)
	} else {
		level := zap.NewAtomicLevel()
		level.SetLevel(logLevel)
		encoderCfg := zap.NewProductionEncoderConfig()
		encoder := zapcore.NewConsoleEncoder(encoderCfg)
		encoder = zapcore.NewJSONEncoder(encoderCfg)
		zapCore = zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), level)
	}
	logger := zap.New(zapCore, zap.AddCaller(), zap.ErrorOutput(zapcore.Lock(os.Stderr)))
	logger = logger.With(zap.String("app", appName), zap.String("_BuildVersion", version))
	return logger
}

func createSigningKeys(l *zap.Logger, cfg *config.Config) (bundleSigningKey *ecdsa.PrivateKey, senderSigningKey *ecdsa.PrivateKey) {
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
