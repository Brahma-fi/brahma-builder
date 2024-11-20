package morpho

import (
	"github.com/Brahma-fi/brahma-builder/internal/entity"
	"github.com/ethereum/go-ethereum/common"
)

type AutomationState struct {
	// the target vault to which money was deposited to
	TargetVault common.Address `json:"targetVault"`
	// amount that was deposited in this vault
	InputAmount string `json:"inputAmount"`
	// amount which was transferred as fees
	FeesAmount string `json:"feesAmount"`
	// amount which was generated as yield
	GeneratedYield string `json:"generatedYield"`
}

type TransitionState struct {
	Current AutomationState  `json:"current"`
	Prev    *AutomationState `json:"prev"`
}

type ExecutionMetadata struct {
	Req             *entity.SignAndExecuteRequest `json:"req"`
	TransitionState TransitionState               `json:"transitionState"`
	TaskID          string                        `json:"taskID"`
}

type ExecutionLog struct {
	Message  string
	Metadata ExecutionMetadata
}
