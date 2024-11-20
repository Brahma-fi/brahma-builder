package morpho

import (
	"context"
	"fmt"

	"github.com/Brahma-fi/brahma-builder/config"
	"github.com/Brahma-fi/brahma-builder/internal/usecase/integrations"
	"github.com/Brahma-fi/brahma-builder/internal/usecase/services"
	"github.com/Brahma-fi/brahma-builder/internal/usecase/workflows/activities/morpho"
	"github.com/Brahma-fi/brahma-builder/pkg/log"
	"github.com/Brahma-fi/brahma-builder/pkg/rpc"
	"github.com/Brahma-fi/brahma-builder/pkg/temporal"
	"github.com/Brahma-fi/brahma-builder/pkg/vault"
	"github.com/ethereum/go-ethereum/common"
	"go.temporal.io/sdk/worker"
)

func Run(id string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := log.NewLogger("morpho-worker", "info")

	vaultCli, err := vault.New(ctx)
	if err != nil {
		return err
	}

	if err := vaultCli.RunLifetimeWatcher(logger); err != nil {
		return err
	}

	defer vaultCli.StopTokenRenew()
	cfg := &config.Config{}
	if err = vault.LoadConfig(cfg, vaultCli); err != nil {
		return err
	}

	temporalClient, err := temporal.NewClient(ctx, cfg.TemporalConfig, log.NewTemporalLoggerFromExisting(logger))
	if err != nil {
		return fmt.Errorf("failed to create temporal client: %w", err)
	}
	defer temporalClient.Close()

	rpcWithFallback, err := rpc.NewRPC(cfg.ChainID2RPCURLs)
	if err != nil {
		return fmt.Errorf("failed to connect to base client: %w", err)
	}

	console := integrations.NewConsoleClient(cfg.ConsoleBaseURL)
	executorConfig, err := cfg.NewExecutorConfigRepo().ByID(id)
	if err != nil {
		return fmt.Errorf("failed to fetch morpho config: %w", err)
	}

	strategyConfig, err := morpho.ParseConfig(executorConfig.StrategyConfig)
	if err != nil {
		return fmt.Errorf("failed to parse morpho startegy config: %w", err)
	}

	baseClient, err := rpcWithFallback.RetryableClient(executorConfig.ChainID)
	if err != nil {
		return fmt.Errorf("failed to connect to base client: %w", err)
	}

	morphoClient := integrations.NewMorphoClient(strategyConfig.BaseURL, baseClient)
	executor, err := services.NewConsoleExecutor(ctx,
		common.HexToAddress(executorConfig.Address),
		executorConfig.ChainID,
		rpcWithFallback,
		vault.NewKeyManager(vaultCli, "console-kernel"),
		console,
		common.HexToAddress(executorConfig.Signer),
		common.HexToAddress(cfg.ExecutorPluginAddress),
	)
	if err != nil {
		return fmt.Errorf("failed to create console executor: %w", err)
	}

	activity, err := morpho.NewReBalancingStrategy(
		morphoClient,
		executor,
		baseClient,
		// logs repo, check interface for impl
		nil,
		strategyConfig,
		// pricing oracle, check interface for impl
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to create morpho activity: %w", err)
	}

	return temporal.RunWorkflow(
		temporalClient,
		executorConfig.TaskQueue,
		worker.Options{},
		nil,
		[]any{activity.ExecutionHandler},
	)
}
