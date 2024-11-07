package config

import (
	"fmt"

	"github.com/Brahma-fi/brahma-builder/internal/entity"
	"github.com/Brahma-fi/brahma-builder/pkg/rpc"
	"github.com/Brahma-fi/brahma-builder/pkg/temporal"
	"github.com/Brahma-fi/brahma-builder/pkg/vault"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	TemporalConfig         temporal.Config
	ConsoleBaseURL         string                 `json:"consoleBaseURL" envconfig:"CONSOLE_BASE_URL"`
	ChainID2RPCURLs        rpc.ChainID2RpcURLs    `json:"chainID2RPCURLs" envconfig:"CHAIN_ID_2_RPC"`
	ExecutorConfig         entity.ExecutorConfigs `json:"executorConfig" envconfig:"EXECUTOR_CONFIG"`
	SyncSubscriptionsEvery string                 `json:"syncSubscriptionsEvery" envconfig:"SYNC_SUBSCRIPTIONS_EVERY"`
	ExecutorPluginAddress  string                 `json:"executorPluginAddress" envconfig:"EXECUTOR_PLUGIN_ADDRESS"`
	ServiceName            string                 `json:"serviceName" envconfig:"SERVICE_NAME"`
	HostPort               string                 `json:"hostPort" envconfig:"HOST_PORT"`
}

func (c Config) NewExecutorConfigRepo() entity.ExecutorConfigRepo {
	return entity.NewExecutorConfigRepo(c.ExecutorConfig)
}

type Option func(*envConfig)

type envConfig struct {
	envFilePath string
}

func defaultEnvConfig() *envConfig {
	return &envConfig{
		envFilePath: ".env",
	}
}

func WithEnvPath(path string) Option {
	return func(cfg *envConfig) {
		cfg.envFilePath = path
	}
}

func NewConfigFromVault(client vault.Client) (*Config, error) {
	cfg := &Config{}
	err := vault.LoadConfig(cfg, client)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func LoadEnvConfig(cfg any, options ...Option) error {

	defaultCfg := defaultEnvConfig()
	for _, option := range options {
		option(defaultCfg)
	}

	err := godotenv.Overload(defaultCfg.envFilePath)
	if err != nil {
		return fmt.Errorf("failed to read env file, %w", err)
	}

	err = envconfig.Process("", cfg)
	if err != nil {
		return fmt.Errorf("failed to fill config structure.en, %w", err)
	}

	return nil
}

func NewConfigFromEnv() (*Config, error) {
	cfg := &Config{}
	err := LoadEnvConfig(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
