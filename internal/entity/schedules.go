package entity

import "time"

const (
	StrategyIDMorphoRebalancerBase    = "morpho-rebalancer-base"
	StrategyIDMorphoRebalancerMainnet = "morpho-rebalancer-mainnet"
)

type ExecutionStatus string

const (
	ExecutionStatusRunning   ExecutionStatus = "Running"
	ExecutionStatusCompleted                 = "Completed"
	ExecutionStatusFailed                    = "Failed"
	ExecutionStatusCanceled                  = "Canceled"
)

type Schedule struct {
	Config     ExecuteWorkflowParams
	ScheduleID string
	CreatedAt  time.Time
}
