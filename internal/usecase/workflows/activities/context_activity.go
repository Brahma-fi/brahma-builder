package activities

import (
	"context"
	"time"

	"github.com/Brahma-fi/brahma-builder/internal/entity"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

type ScheduleHandlerRepo interface {
	GetHandle(ctx context.Context, scheduleID string) client.ScheduleHandle
}

type ContextActivity struct {
	repo ScheduleHandlerRepo
}

func NewContextActivity(repo ScheduleHandlerRepo) *ContextActivity {
	return &ContextActivity{repo: repo}
}

func (c *ContextActivity) GetExecutionContext(ctx context.Context, scheduleID string) (*entity.ScheduleCtx, error) {
	schedule, err := c.repo.GetHandle(ctx, scheduleID).Describe(ctx)
	if err != nil {
		return nil, err
	}

	execCtx := &entity.ScheduleCtx{
		RunningExecutionWorkflowIDs: make([]string, 0),
		ExecutionCount:              uint(schedule.Info.NumActions + 1),
	}

	for i, w := range schedule.Info.RunningWorkflows {
		// first id is always the currently running workflow
		if i == 0 {
			continue
		}

		execCtx.RunningExecutionWorkflowIDs = append(execCtx.RunningExecutionWorkflowIDs, w.WorkflowID)
	}

	if len(schedule.Info.RecentActions) != 0 {
		latest := schedule.Info.RecentActions[len(schedule.Info.RecentActions)-1]
		execCtx.PrevExecutionID = latest.StartWorkflowResult.WorkflowID
		execCtx.PrevExecutionAt = latest.ActualTime
	}

	return execCtx, nil
}

func (c *ContextActivity) Options() workflow.ActivityOptions {
	return workflow.ActivityOptions{
		TaskQueue: entity.BaseTaskQueue,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumInterval:        time.Second * 10,
			MaximumAttempts:        1,
			NonRetryableErrorTypes: []string{},
		},
		StartToCloseTimeout: time.Second * 10,
	}

}
