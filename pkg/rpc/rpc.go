package rpc

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/ethereum/go-ethereum/ethclient"
)

type ChainID2RpcURLs map[string][]string

func chainID2Int(in string) (int64, error) {
	chainID, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%w, invalid chain_id", err)
	}

	return chainID, nil
}

type RPC struct {
	clients map[uint64]*Clients
	cfg     ChainID2RpcURLs
}

type RawClientWithID struct {
	*ethclient.Client
	id string
}

func (r *RawClientWithID) ID() string {
	return r.id
}

func NewRPC(cfg ChainID2RpcURLs) (*RPC, error) {
	r := &RPC{clients: make(map[uint64]*Clients), cfg: cfg}
	for chain, rpcURLs := range cfg {
		if len(rpcURLs) == 0 {
			return nil, fmt.Errorf("no clients found for chainID %s", chain)
		}

		chainID, err := chainID2Int(chain)
		if err != nil {
			return nil, err
		}

		c := &Clients{
			fallbacks: make([]RawClient, 0),
			rawURLS:   make([]string, 0),
			client:    &http.Client{},
			chainID:   chainID,
		}

		for i, rpcURL := range rpcURLs {
			uri, err := url.Parse(rpcURL)
			if err != nil {
				return nil, err
			}

			c.rawURLS = append(c.rawURLS, rpcURL)
			client, err := ethclient.Dial(rpcURL)
			if err != nil {
				return nil, err
			}

			// provider[0] is the primary provider by default
			if i == 0 {
				c.primary = &RawClientWithID{
					Client: client,
					id:     uri.Host,
				}
				continue
			}

			c.fallbacks = append(c.fallbacks, &RawClientWithID{
				Client: client,
				id:     uri.Host,
			})
		}

		r.clients[uint64(chainID)] = c
	}

	return r, nil
}

func (r *RPC) configured(chainId uint64) error {
	if v, ok := r.clients[chainId]; !ok || v.primary == nil {
		return ErrInvalidChainID
	}

	return nil
}

func (r *RPC) Client(chainID int64) (*ethclient.Client, error) {
	if err := r.configured(uint64(chainID)); err != nil {
		return nil, err
	}

	return ethclient.Dial(r.cfg[strconv.FormatUint(uint64(chainID), 10)][0])
}

func (r *RPC) RetryableClient(chainID int64) (*Clients, error) {
	if err := r.configured(uint64(chainID)); err != nil {
		return nil, err
	}

	return r.clients[uint64(chainID)], nil
}
