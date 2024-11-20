package cli

import (
	"context"

	"github.com/Brahma-fi/brahma-builder/app/scheduler"
	"github.com/Brahma-fi/brahma-builder/app/worker/base"
	"github.com/Brahma-fi/brahma-builder/app/worker/morpho"
	"github.com/Brahma-fi/brahma-builder/internal/entity"
	"github.com/urfave/cli/v3"
)

func BuildCLI() *cli.Command {
	var executorID string
	return &cli.Command{
		Commands: []*cli.Command{
			{
				Name:  "scheduler",
				Usage: "Runs sync scheduler",
				Action: func(_ context.Context, _ *cli.Command) error {
					return scheduler.Run()
				},
			},
			{
				Name:    "base-worker",
				Aliases: []string{"base"},
				Usage:   "Runs base server",
				Action: func(_ context.Context, _ *cli.Command) error {
					return base.Run()
				},
			},
			{
				Name:    "morpho-worker",
				Aliases: []string{"morpho"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "id",
						Destination: &executorID,
						OnlyOnce:    true,
						Value:       entity.StrategyIDMorphoRebalancerMainnet,
					},
				},
				Usage: "Runs morpho worker",
				Action: func(_ context.Context, _ *cli.Command) error {
					return morpho.Run(executorID)
				},
			},
		},
	}
}
