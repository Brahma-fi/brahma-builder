package services

import (
	"context"

	"github.com/Brahma-fi/brahma-builder/internal/entity"
	"github.com/ethereum/go-ethereum/common"
	"go.temporal.io/sdk/workflow"
)

type orchestrator interface {
	OrchestratorWorkflow(ctx workflow.Context, config entity.ExecuteWorkflowParams) error
}

type keyManager interface {
	Sign(ctx context.Context, hash string, _ common.Address) ([]byte, error)
}

type console interface {
	ActiveSubscriptions(
		ctx context.Context,
		registryID string,
	) ([]entity.ClientSubscription, error)
	Subscriptions(ctx context.Context, registryID string) ([]entity.ClientSubscription, error)
	ExecutorByAddressAndChainID(
		ctx context.Context,
		address common.Address,
		chainID uint64,
	) (*entity.ExecutorMetadata, error)
	Execute(ctx context.Context, req *entity.ExecuteTaskReq) (*entity.ExecuteTaskResp, error)
}

type executors interface {
	List() []entity.ExecutorConfig
	Config(executor common.Address) (*entity.ExecutorConfig, error)
}

type schedulesRepo interface {
	BySubAccountAddressChainIDAndStatus(
		ctx context.Context,
		subAccount common.Address,
		chainID int64,

	) ([]entity.Schedule, error)
	BySubAccountAddressesChainIDAndStatus(
		ctx context.Context,
		subAccounts []common.Address,
		chainID int64,

	) ([]entity.Schedule, error)
}
