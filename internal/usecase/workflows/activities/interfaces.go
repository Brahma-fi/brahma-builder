package activities

import (
	"github.com/Brahma-fi/brahma-builder/pkg/rpc"
)

type rpcFactory interface {
	RetryableClient(chainID int64) (*rpc.Clients, error)
}
