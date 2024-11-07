package entity

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

const (
	SearchAttrKeySubAccountAddress = "subaccountAddress"
	SearchAttrKeyExecutorAddress   = "executorAddress"
	SearchAttrKeyChainID           = "chainID"
	SearchAttrKeyExecutorID        = "executorID"
)

type ExecuteWorkflowParams struct {
	Nonce    uint64                   `json:"nonce"`
	Params   OrchestratorParams       `json:"params"`
	Schedule *ScheduledWorkflowConfig `json:"schedule"`
}

type OrchestratorParams struct {
	ExecutorAddress   common.Address     `json:"executorAddress"`
	SubAccountAddress common.Address     `json:"subAccountAddress"`
	ExecutorID        string             `json:"executorID"`
	ChainID           int64              `json:"chainID"`
	Subscription      ClientSubscription `json:"subscription"`
}

func (o OrchestratorParams) ID() string {
	var hashInput []byte
	hashInput = append(hashInput, o.SubAccountAddress.Bytes()...)
	hashInput = append(hashInput, o.ExecutorAddress.Bytes()...)
	binary.LittleEndian.PutUint64(hashInput, uint64(o.ChainID))
	return fmt.Sprintf("%x", sha256.Sum256(hashInput))
}

type ScheduledWorkflowConfig struct {
	Every time.Duration `json:"every"`
	ID    string        `json:"ID"`
}
