package temporal

import (
	"context"
	"time"

	"go.temporal.io/api/workflowservice/v1"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/log"
	"go.temporal.io/sdk/worker"
	"google.golang.org/protobuf/types/known/durationpb"
)

type Config struct {
	TemporalHost      string `json:"temporalHost" envconfig:"TEMPORAL_HOST"`
	TemporalNameSpace string `json:"temporalNameSpace" envconfig:"TEMPORAL_NAMESPACE"`
}

const (
	// Temporal offers configurable retention periods for Workflow Execution data.
	// This retention period is adjustable per Namespace through the Temporal Web UI,
	// ranging from 1 to 90 days.
	_DefaultWorkflowExecutionRetention = 30 * 24 * time.Hour // one month
)

func NewClient(ctx context.Context, cfg Config, logger log.Logger) (client.Client, error) {
	clientOptions := client.Options{
		HostPort:  cfg.TemporalHost,
		Namespace: cfg.TemporalNameSpace,
		Logger:    logger,
	}

	err := registerNamespace(ctx, cfg.TemporalNameSpace, clientOptions)
	if err != nil {
		return nil, err
	}

	temporalClient, err := client.Dial(clientOptions)
	if err != nil {
		return nil, err
	}

	return temporalClient, nil
}

func RunWorkflow(
	cli client.Client,
	taskQueue string,
	options worker.Options,
	workflowFunc any,
	activityFuncs []any,
) error {
	w := worker.New(cli, taskQueue, options)
	if workflowFunc != nil {
		w.RegisterWorkflow(workflowFunc)
	}

	for i := range activityFuncs {
		w.RegisterActivity(activityFuncs[i])
	}

	return w.Run(worker.InterruptCh())
}

func registerNamespace(ctx context.Context, serviceName string, options client.Options) error {
	namespaceClient, err := client.NewNamespaceClient(options)
	if err != nil {
		return err
	}

	err = namespaceClient.Register(
		ctx, &workflowservice.RegisterNamespaceRequest{
			Namespace:                        serviceName,
			WorkflowExecutionRetentionPeriod: durationpb.New(_DefaultWorkflowExecutionRetention),
		},
	)
	switch {
	case err == nil:
	case err.Error() == "Namespace already exists.":
	default:
		return err
	}

	return nil
}
