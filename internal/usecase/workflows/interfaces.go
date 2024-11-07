package workflows

import (
	"context"

	"github.com/Brahma-fi/brahma-builder/internal/entity"
	"github.com/ethereum/go-ethereum/common"
	"go.temporal.io/sdk/workflow"
)

type SubscriptionStateReader interface {
	ActiveSubscriptions(
		ctx context.Context,
		registryID string,
	) ([]entity.ClientSubscription, error)
	ExecutorByAddress(
		ctx context.Context,
		address common.Address,
		chainID uint64,
	) (*entity.ExecutorMetadata, error)
	Subscriptions(ctx context.Context, registryID string) ([]entity.ClientSubscription, error)
}

type Executor interface {
	Execute(ctx context.Context, req *entity.ExecuteTaskReq) (*entity.ExecuteTaskResp, error)
}

type activityOptions interface {
	Options() workflow.ActivityOptions
}

type contextActivity interface {
	GetExecutionContext(ctx context.Context, scheduleID string) (*entity.ScheduleCtx, error)
	activityOptions
}

type configRepo interface {
	Config(executor common.Address) (*entity.ExecutorConfig, error)
}
