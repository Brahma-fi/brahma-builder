package cli

import (
	"context"

	"github.com/Brahma-fi/brahma-builder/app/scheduler"
	"github.com/Brahma-fi/brahma-builder/app/worker/base"
	"github.com/urfave/cli/v3"
)

func BuildCLI() *cli.Command {
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
		},
	}
}
