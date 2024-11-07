package rpc

import "errors"

var (
	ErrInvalidChainID           = errors.New("invalid chain id")
	ErrFailedToCallAllUpstreams = errors.New("failed to call all upstreams")
)
