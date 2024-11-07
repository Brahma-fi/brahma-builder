package entity

import (
	"encoding/json"
	"time"
)

type ExecutorMetadata struct {
	Config struct {
		FeeInBPS          int      `json:"feeInBPS"`
		FeeReceiver       string   `json:"feeReceiver"`
		LimitPerExecution bool     `json:"limitPerExecution"`
		FeeToken          string   `json:"feeToken"`
		InputTokens       []string `json:"inputTokens"`
		HopAddresses      []string `json:"hopAddresses"`
	} `json:"config"`
	Executor  string `json:"executor"`
	Signature string `json:"signature"`
	ChainId   int    `json:"chainId"`
	Timestamp int    `json:"timestamp"`
	Metadata  struct {
		Id       string `json:"id"`
		Name     string `json:"name"`
		Logo     string `json:"logo"`
		Metadata struct {
		} `json:"metadata"`
	} `json:"executorMetadata"`
	Id     string `json:"id"`
	Status int    `json:"status"`
}

type GetExecutorMetadataResp struct {
	Data  ExecutorMetadata `json:"data"`
	Error string           `json:"error"`
}

type ClientSubscription struct {
	ChainId           int               `json:"chainId"`
	CommitHash        string            `json:"commitHash"`
	CreatedAt         time.Time         `json:"createdAt"`
	Duration          int               `json:"duration"`
	FeeAmount         string            `json:"feeAmount"`
	FeeToken          string            `json:"feeToken"`
	Id                string            `json:"id"`
	Metadata          json.RawMessage   `json:"metadata"`
	RegistryId        string            `json:"registryId"`
	Status            int               `json:"status"`
	SubAccountAddress string            `json:"subAccountAddress"`
	TokenInputs       map[string]string `json:"tokenInputs"`
	TokenLimits       map[string]string `json:"tokenLimits"`
}

type GetClientSubscriptionsResp struct {
	Data  []ClientSubscription `json:"data"`
	Error string               `json:"error"`
}

type VerifyExecutableReq struct {
	Task struct {
		Subaccount        string `json:"subaccount"`
		Executor          string `json:"executor"`
		ExecutorSignature string `json:"executorSignature"`
		Executable        struct {
			CallType string `json:"callType"`
			To       string `json:"to"`
			Value    int    `json:"value"`
			Data     string `json:"data"`
		} `json:"executable"`
	} `json:"task"`
}

type VerifyExecutableResp struct {
	Data struct {
		Expiry          int    `json:"expiry"`
		PolicySignature string `json:"policySignature"`
		Error           string `json:"error"`
	} `json:"data"`
	Error string `json:"error"`
}

type Executable struct {
	CallType uint8  `json:"callType"`
	To       string `json:"to"`
	Value    string `json:"value"`
	Data     string `json:"data"`
}

type Task struct {
	Subaccount        string     `json:"subaccount"`
	Executor          string     `json:"executor"`
	ExecutorSignature string     `json:"executorSignature"`
	Executable        Executable `json:"executable"`
}

type ExecuteTaskReq struct {
	ChainID int64  `json:"-"`
	Task    Task   `json:"task"`
	Webhook string `json:"webhook"`
}

type ExecuteTaskResp struct {
	Data struct {
		Data struct {
			TaskId string `json:"taskId"`
		} `json:"data"`
		Error  string `json:"error"`
		Errors string `json:"errors"`
	} `json:"data"`
	Error string `json:"error"`
}
