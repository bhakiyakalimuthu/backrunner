package config

import (
	"github.com/caarlos0/env"
	"github.com/go-playground/validator/v10"
)

// Config contains environment values
type Config struct {
	DebugLog          bool   `env:"DEBUG_LOG" envDefault:"false"`
	LogJSON           bool   `env:"LOG_JSON" envDefault:"true"`
	AppName           string `env:"APP_NAME" envDefault:"backrunner"`
	EthClientURL      string `env:"ETH_CLIENT_URL" envDefault:""`
	FlashbotsRelayURL string `env:"FLASHBOTS_RELAY_URL" envDefault:"https://relay.flashbots.net"`
	BundleSigningKey  string `env:"BUNDLE_SINGING_KEY" envDefault:""`
	SenderSigningKey  string `env:"SENDER_SINGING_KEY" envDefault:""`
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
