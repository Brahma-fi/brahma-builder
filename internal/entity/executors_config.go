package entity

import (
	"errors"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

type ExecutorConfig struct {
	ActivityTimeout      string         `json:"activityTimeout"`
	TaskQueue            string         `json:"taskQueue"`
	RetryAttempts        uint           `json:"retryAttempts"`
	MaximumRetryInterval string         `json:"maximumRetryInterval"`
	ChainID              int64          `json:"chainId"`
	Address              string         `json:"address"`
	Signer               string         `json:"signer"`
	Every                string         `json:"every"`
	StrategyConfig       map[string]any `json:"strategyConfig"`
	ID                   string         `json:"Id"`
}

func (e ExecutorConfig) ActivityOptions() (workflow.ActivityOptions, error) {
	interval, err := time.ParseDuration(e.MaximumRetryInterval)
	if err != nil {
		return workflow.ActivityOptions{}, err
	}

	timeout, err := time.ParseDuration(e.ActivityTimeout)
	if err != nil {
		return workflow.ActivityOptions{}, err
	}

	return workflow.ActivityOptions{
		TaskQueue: e.TaskQueue,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumInterval: interval,
			MaximumAttempts: int32(e.RetryAttempts),
		},
		StartToCloseTimeout: timeout,
	}, nil
}

type ExecutorConfigs []ExecutorConfig

type ExecutorConfigRepo struct {
	storage map[common.Address]ExecutorConfig
}

func (e ExecutorConfigRepo) Config(executor common.Address) (*ExecutorConfig, error) {
	if v, ok := e.storage[executor]; !ok {
		return nil, errors.New("executor not found")
	} else {
		return &v, nil
	}
}

func (e ExecutorConfigRepo) List() []ExecutorConfig {
	executors := make([]ExecutorConfig, 0)
	for _, cfg := range e.storage {
		executors = append(executors, cfg)
	}

	return executors
}

func (e ExecutorConfigRepo) ByID(id string) (*ExecutorConfig, error) {
	for _, v := range e.storage {
		if v.ID == id {
			return &v, nil
		}
	}

	return nil, errors.New("executor not found")
}

func NewExecutorConfigRepo(cfg ExecutorConfigs) ExecutorConfigRepo {
	storage := make(map[common.Address]ExecutorConfig)
	for _, c := range cfg {
		storage[common.HexToAddress(c.Address)] = c
	}

	return ExecutorConfigRepo{
		storage: storage,
	}
}
