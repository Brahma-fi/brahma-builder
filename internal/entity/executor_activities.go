package entity

import (
	"time"
)

const (
	BaseTaskQueue    = "base-task-queue"
	DefaultNamespace = "brahma-builder"
)

func ExecutionHandler(_ ExecCtx) error {
	return nil
}

type ExecCtx struct {
	ScheduleCtx
	ExecuteWorkflowParams
	TriggeredAt time.Time
}
