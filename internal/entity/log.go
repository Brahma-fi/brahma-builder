package entity

import (
	"database/sql/driver"
	"time"

	"github.com/google/uuid"
)

type Jsonb driver.Value

type Log struct {
	ID                uuid.UUID `db:"id" json:"id"`
	SubscriptionID    uuid.UUID `db:"sub_id" json:"subscriptionId"`
	ChainID           int64     `db:"chain_id" json:"chainId"`
	Metadata          Jsonb     `db:"metadata" json:"metadata"`
	SubAccountAddress string    `db:"subaccount_address" json:"subAccountAddress"`
	Message           string    `db:"message" json:"message"`
	OutputTxnHash     string    `db:"output_txn" json:"OutputTxnHash"`
	CreatedAt         time.Time `db:"created_at" json:"createdAt"`
}
