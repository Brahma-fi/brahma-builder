package base

import (
	"context"
	"fmt"

	"github.com/Brahma-fi/brahma-builder/config"
	"github.com/Brahma-fi/brahma-builder/internal/entity"
	"github.com/Brahma-fi/brahma-builder/internal/usecase/workflows"
	"github.com/Brahma-fi/brahma-builder/internal/usecase/workflows/activities"
	"github.com/Brahma-fi/brahma-builder/pkg/log"
	"github.com/Brahma-fi/brahma-builder/pkg/temporal"
	"github.com/Brahma-fi/brahma-builder/pkg/vault"
	"go.temporal.io/sdk/worker"
)

func Run() error {
	var err error
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := log.NewLogger("base-worker", "info")

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

	ctxActivity := activities.NewContextActivity(temporalClient.ScheduleClient())

	orchestratorActivity := workflows.NewOrchestrator(
		ctxActivity,
		cfg.NewExecutorConfigRepo(),
	)

	return temporal.RunWorkflow(
		temporalClient,
		entity.BaseTaskQueue,
		worker.Options{},
		orchestratorActivity.OrchestratorWorkflow,
		[]any{
			ctxActivity.GetExecutionContext,
		},
	)
}
