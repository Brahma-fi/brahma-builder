package scheduler

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Brahma-fi/brahma-builder/config"
	"github.com/Brahma-fi/brahma-builder/internal/entity"
	"github.com/Brahma-fi/brahma-builder/internal/repo"
	integration "github.com/Brahma-fi/brahma-builder/internal/usecase/integrations"
	"github.com/Brahma-fi/brahma-builder/internal/usecase/services"
	"github.com/Brahma-fi/brahma-builder/pkg/log"
	"github.com/Brahma-fi/brahma-builder/pkg/temporal"
	"github.com/Brahma-fi/brahma-builder/pkg/vault"
)

func Run() error {
	var err error
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := log.NewLogger("sync-scheduler", "info")

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
		return err
	}
	defer temporalClient.Close()

	console := integration.NewConsoleClient(cfg.ConsoleBaseURL)

	executors := entity.NewExecutorConfigRepo(cfg.ExecutorConfig)

	schedulesRepo := repo.NewSchedulesRepo(temporalClient)

	scheduler, err := services.NewScheduler(ctx, temporalClient, console, executors, schedulesRepo)
	if err != nil {
		return err
	}

	sync, err := time.ParseDuration(cfg.SyncSubscriptionsEvery)
	if err != nil {
		return err
	}

	sch := time.Tick(sync)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	if err = scheduler.Sync(ctx); err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			fmt.Println("manager - Run - context canceled")
			return nil
		case s := <-interrupt:
			fmt.Println("manager - Run - signal: " + s.String())
			return nil
		case <-sch:
			if err = scheduler.Sync(ctx); err != nil {
				fmt.Println("failed to call sync", err)
			}
		}
	}

}
