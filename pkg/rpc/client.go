package rpc

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"net/http"

	"github.com/Brahma-fi/brahma-builder/pkg/log"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	geth "github.com/ethereum/go-ethereum/core/types"
	"github.com/labstack/echo/v4"
)

type RawClient interface {
	CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error)
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*geth.Receipt, error)
	TransactionByHash(ctx context.Context, txHash common.Hash) (*geth.Transaction, bool, error)
	ChainID(ctx context.Context) (*big.Int, error)
	BlockNumber(ctx context.Context) (uint64, error)
	EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error)
	SuggestGasPrice(ctx context.Context) (*big.Int, error)
	SuggestGasTipCap(ctx context.Context) (*big.Int, error)
	FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]geth.Log, error)
	BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error)
	StorageAt(ctx context.Context, account common.Address, key common.Hash, blockNumber *big.Int) ([]byte, error)
	CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error)
	NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error)
	BlockByHash(ctx context.Context, hash common.Hash) (*geth.Block, error)
	BlockByNumber(ctx context.Context, number *big.Int) (*geth.Block, error)
	HeaderByHash(ctx context.Context, hash common.Hash) (*geth.Header, error)
	HeaderByNumber(ctx context.Context, number *big.Int) (*geth.Header, error)
	TransactionCount(ctx context.Context, blockHash common.Hash) (uint, error)
	TransactionInBlock(ctx context.Context, blockHash common.Hash, index uint) (*geth.Transaction, error)
	ID() string
}

type RawJSONRpcRequest struct {
	ID      int           `json:"id"`
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

type RawJSONRpcResponse struct {
	ID      int             `json:"id"`
	Jsonrpc string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result"`
	Error   any             `json:"error"`
}

type HTTPRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Clients is a retry-based ethclient caller that retries a call over multiple upstreams if an error is received.
// Implements: TransactionReader, ChainStateReader, ContractCaller, GasPricer, GasPricer1559, GasEstimator, BlockNumberReader, ChainIDReader.
// Partially implements: ChainReader, LogFilterer.
type Clients struct {
	primary   RawClient
	fallbacks []RawClient
	rawURLS   []string
	client    HTTPRequestDoer
	chainID   int64
}

func callFirstSuccessful[T any](
	ctx context.Context,
	clients *Clients,
	call func(RawClient) (T, error),
) (T, error) {
	var zero T
	result, err := call(clients.primary)
	if err == nil {
		return result, nil
	}

	logger := log.GetLogger(ctx)
	logger.Warn("failed to call primary upstream rpc", log.Str("provider", clients.primary.ID()), log.Int("chainID", int(clients.chainID)))

	for _, i := range rand.Perm(len(clients.fallbacks)) {
		result, err = call(clients.fallbacks[i])
		if err == nil {
			return result, nil
		}

		logger.Warn("failed to call fallback upstream rpc", log.Str("provider", clients.fallbacks[i].ID()), log.Int("chainID", int(clients.chainID)))
	}

	logger.Error("failed to call all upstream rpc", log.Int("chainID", int(clients.chainID)))
	return zero, ErrFailedToCallAllUpstreams
}

func (c *Clients) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	return callFirstSuccessful(ctx, c, func(client RawClient) ([]byte, error) {
		return client.CallContract(ctx, call, blockNumber)
	})
}

func (c *Clients) TransactionReceipt(ctx context.Context, txHash common.Hash) (*geth.Receipt, error) {
	return callFirstSuccessful(ctx, c, func(client RawClient) (*geth.Receipt, error) {
		return client.TransactionReceipt(ctx, txHash)
	})
}

func (c *Clients) TransactionByHash(ctx context.Context, txHash common.Hash) (*geth.Transaction, bool, error) {
	tx, isPending, err := c.primary.TransactionByHash(ctx, txHash)
	if err == nil {
		return tx, isPending, nil
	}

	logger := log.GetLogger(ctx)
	logger.Warn("failed to call primary upstream rpc", log.Str("provider", c.primary.ID()), log.Int("chainID", int(c.chainID)))

	for _, i := range rand.Perm(len(c.fallbacks)) {
		tx, isPending, err = c.fallbacks[i].TransactionByHash(ctx, txHash)
		if err == nil {
			return tx, isPending, nil
		}

		logger.Warn("failed to call fallback upstream rpc", log.Str("provider", c.fallbacks[i].ID()), log.Int("chainID", int(c.chainID)))
	}

	logger.Error("failed to call all upstream rpc", log.Int("chainID", int(c.chainID)))
	return nil, false, ErrFailedToCallAllUpstreams
}

func (c *Clients) ChainID(ctx context.Context) (*big.Int, error) {
	return callFirstSuccessful(ctx, c, func(client RawClient) (*big.Int, error) {
		return client.ChainID(ctx)
	})
}

func (c *Clients) BlockNumber(ctx context.Context) (uint64, error) {
	return callFirstSuccessful(ctx, c, func(client RawClient) (uint64, error) {
		return client.BlockNumber(ctx)
	})
}

func (c *Clients) EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error) {
	return callFirstSuccessful(ctx, c, func(client RawClient) (uint64, error) {
		return client.EstimateGas(ctx, call)
	})
}

func (c *Clients) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return callFirstSuccessful(ctx, c, func(client RawClient) (*big.Int, error) {
		return client.SuggestGasPrice(ctx)
	})
}

func (c *Clients) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	return callFirstSuccessful(ctx, c, func(client RawClient) (*big.Int, error) {
		return client.SuggestGasTipCap(ctx)
	})
}

func (c *Clients) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]geth.Log, error) {
	return callFirstSuccessful(ctx, c, func(client RawClient) ([]geth.Log, error) {
		return client.FilterLogs(ctx, q)
	})
}

func (c *Clients) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return callFirstSuccessful(ctx, c, func(client RawClient) (*big.Int, error) {
		return client.BalanceAt(ctx, account, blockNumber)
	})
}

func (c *Clients) StorageAt(
	ctx context.Context,
	account common.Address,
	key common.Hash,
	blockNumber *big.Int,
) ([]byte, error) {
	return callFirstSuccessful(ctx, c, func(client RawClient) ([]byte, error) {
		return client.StorageAt(ctx, account, key, blockNumber)
	})
}

func (c *Clients) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	return callFirstSuccessful(ctx, c, func(client RawClient) ([]byte, error) {
		return client.CodeAt(ctx, account, blockNumber)
	})
}

func (c *Clients) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	return callFirstSuccessful(ctx, c, func(client RawClient) (uint64, error) {
		return client.NonceAt(ctx, account, blockNumber)
	})
}

func (c *Clients) BlockByHash(ctx context.Context, hash common.Hash) (*geth.Block, error) {
	return callFirstSuccessful(ctx, c, func(client RawClient) (*geth.Block, error) {
		return client.BlockByHash(ctx, hash)
	})
}

func (c *Clients) BlockByNumber(ctx context.Context, number *big.Int) (*geth.Block, error) {
	return callFirstSuccessful(ctx, c, func(client RawClient) (*geth.Block, error) {
		return client.BlockByNumber(ctx, number)
	})
}

func (c *Clients) HeaderByHash(ctx context.Context, hash common.Hash) (*geth.Header, error) {
	return callFirstSuccessful(ctx, c, func(client RawClient) (*geth.Header, error) {
		return client.HeaderByHash(ctx, hash)
	})
}

func (c *Clients) HeaderByNumber(ctx context.Context, number *big.Int) (*geth.Header, error) {
	return callFirstSuccessful(ctx, c, func(client RawClient) (*geth.Header, error) {
		return client.HeaderByNumber(ctx, number)
	})
}

func (c *Clients) TransactionCount(ctx context.Context, blockHash common.Hash) (uint, error) {
	return callFirstSuccessful(ctx, c, func(client RawClient) (uint, error) {
		return client.TransactionCount(ctx, blockHash)
	})
}

func (c *Clients) TransactionInBlock(
	ctx context.Context,
	blockHash common.Hash,
	index uint,
) (*geth.Transaction, error) {
	return callFirstSuccessful(ctx, c, func(client RawClient) (*geth.Transaction, error) {
		return client.TransactionInBlock(ctx, blockHash, index)
	})
}

func (c *Clients) RawRequest(ctx context.Context, body *RawJSONRpcRequest) (*RawJSONRpcResponse, error) {
	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	for _, client := range c.rawURLS {
		resp, err := c.callRPC(ctx, client, data)
		if err == nil {
			return resp, nil
		}

	}

	return nil, ErrFailedToCallAllUpstreams
}

func (c *Clients) callRPC(ctx context.Context, client string, body []byte) (*RawJSONRpcResponse, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		client,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusBadRequest:
		return nil, errors.New("bad request")
	default:
		return nil, fmt.Errorf("failed to call rpc")
	}

	var out RawJSONRpcResponse
	err = json.NewDecoder(resp.Body).Decode(&out)
	if err != nil {
		return nil, err
	}

	if out.Error != nil {
		log.GetLogger(ctx).Warn("invalid response from upstream", log.Any("error", out.Error))
		return nil, errors.New("invalid response from upstream")
	}

	return &out, nil
}
