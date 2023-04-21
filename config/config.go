package config

import (
	"github.com/caarlos0/env"
	"github.com/go-playground/validator/v10"
)

// Config contains environment values
type Config struct {
	DebugLog bool   `env:"DEBUG_LOG" envDefault:"false"`
	LogJSON  bool   `env:"LOG_JSON" envDefault:"true"`
	AppName  string `env:"APP_NAME" envDefault:"backrunner"`
	// Websocket client needed to get pending transaction
	// subscription doesn't work in RPC
	EthClientWebSocketURL string `env:"ETH_CLIENT_WSS_URL" envDefault:""`
	// Flashbots relay URL for simulating and sending bundles to builder
	FlashbotsRelayURL string `env:"FLASHBOTS_RELAY_URL" envDefault:"https://relay.flashbots.net"`
	// Bundle signing key for signing Flashbots bundles
	// This key will be identified by Flashbots reputation system
	BundleSigningKey string `env:"BUNDLE_SINGING_KEY" envDefault:""`
	// Sender signing key for signing backrunner transaction
	SenderSigningKey string `env:"SENDER_SINGING_KEY" envDefault:""`
}

// LoadFromEnv parses environment variables into a given struct and validates
// its fields' values.
func LoadFromEnv(config interface{}) error {
	if err := env.Parse(config); err != nil {
		return err
	}
	if err := validator.New().Struct(config); err != nil {
		return err
	}
	return nil
}

func NewConfig() *Config {
	var cfg Config
	if err := LoadFromEnv(&cfg); err != nil {
		panic(err)
	}
	return &cfg
}
