package workflows

import (
	"errors"

	"github.com/Brahma-fi/brahma-builder/internal/entity"
	"github.com/Brahma-fi/brahma-builder/pkg/log"
	"go.temporal.io/sdk/workflow"
)

type Orchestrator struct {
	ctxActivity contextActivity
	config      configRepo
}

func NewOrchestrator(
	ctxActivity contextActivity,
	cfg configRepo,
) *Orchestrator {
	return &Orchestrator{
		ctxActivity: ctxActivity,
		config:      cfg,
	}
}

func (o *Orchestrator) OrchestratorWorkflow(ctx workflow.Context, config entity.ExecuteWorkflowParams) error {
	logger := workflow.GetLogger(ctx)
	workflowInfo := workflow.GetInfo(ctx)
	logger.Info("starting orchestratorWorkflow", log.Str("workflowID", workflowInfo.WorkflowExecution.ID))

	if config.Schedule == nil {
		return errors.New("workflow is not scheduled")
	}

	fetchCtxActivityCtx := workflow.WithActivityOptions(ctx, o.ctxActivity.Options())
	scheduleCtx := &entity.ScheduleCtx{}
	if err := workflow.ExecuteActivity(
		fetchCtxActivityCtx,
		o.ctxActivity.GetExecutionContext,
		config.Schedule.ID,
	).Get(ctx, scheduleCtx); err != nil {
		logger.Error(
			"failed to get execution context",
			log.Err(err),
			log.Str("workflowID", workflowInfo.WorkflowExecution.ID),
		)
		return err
	}

	executorConfig, err := o.config.Config(config.Params.ExecutorAddress)
	if err != nil {
		logger.Error(
			"failed to get executor config",
			log.Err(err),
			log.Str("workflowID", workflowInfo.WorkflowExecution.ID),
		)
		return err
	}

	activityOpts, err := executorConfig.ActivityOptions()
	if err != nil {
		logger.Error(
			"failed to get activity options",
			log.Err(err),
			log.Str("workflowID", workflowInfo.WorkflowExecution.ID),
		)
		return err
	}

	execActivityOptions := workflow.WithActivityOptions(ctx, activityOpts)
	err = workflow.ExecuteActivity(execActivityOptions, entity.ExecutionHandler, entity.ExecCtx{
		ScheduleCtx:           *scheduleCtx,
		ExecuteWorkflowParams: config,
		TriggeredAt:           workflowInfo.WorkflowStartTime,
	}).Get(ctx, nil)
	if err != nil {
		logger.Error(
			"executionHandler failed",
			log.Err(err),
			log.Str("workflowID", workflowInfo.WorkflowExecution.ID),
		)
		return err
	}

	logger.Info(
		"orchestratorWorkflow completed successfully",
		log.Str("workflowID", workflowInfo.WorkflowExecution.ID),
	)
	return nil
}
