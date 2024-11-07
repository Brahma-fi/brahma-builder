package entity

import "time"

type ScheduleCtx struct {
	ExecutionCount              uint      `json:"executionCount"`
	PrevExecutionAt             time.Time `json:"prevExecutionAt"`
	PrevExecutionID             string    `json:"prevExecutionID"`
	RunningExecutionWorkflowIDs []string  `json:"runningExecutionWorkflowIDs"`
}

type CustomIntervalOptions struct {
	Every string `json:"Every"` // Duration which can be parsed using time.ParseDuration or number in seconds
}
