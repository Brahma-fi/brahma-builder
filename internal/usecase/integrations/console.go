package integrations

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Brahma-fi/brahma-builder/internal/entity"
	"github.com/Brahma-fi/brahma-builder/pkg/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-resty/resty/v2"
)

const (
	CancelledStatus = 4
)

type ConsoleClient struct {
	client *resty.Client
}

func NewConsoleClient(base string) *ConsoleClient {
	return &ConsoleClient{
		client: resty.New().SetBaseURL(base),
	}
}

func (c *ConsoleClient) ExecutorByAddressAndChainID(
	ctx context.Context,
	address common.Address,
	chainID uint64,
) (*entity.ExecutorMetadata, error) {
	result := &entity.GetExecutorMetadataResp{}
	resp, err := c.client.R().
		SetContext(ctx).
		SetResult(result).
		Get(fmt.Sprintf("/v1/automations/executor/%s/%d", address.Hex(), chainID))
	switch {
	case err != nil:
		return nil, fmt.Errorf("failed to fetch executor by address: %w", err)
	case resp.StatusCode() == http.StatusNotFound:
		return nil, fmt.Errorf("executor not found: %s", address.Hex())
	case resp.StatusCode() != http.StatusOK:
		return nil, fmt.Errorf("failed to fetch executor: %d", resp.StatusCode())
	}

	return &result.Data, nil
}

func (c *ConsoleClient) ActiveSubscriptions(
	ctx context.Context,
	registryID string,
) ([]entity.ClientSubscription, error) {
	subscriptions, err := c.Subscriptions(ctx, registryID)
	if err != nil {
		return nil, err
	}

	var activeSubscriptions []entity.ClientSubscription
	for _, subscription := range subscriptions {
		if subscription.Status != CancelledStatus {
			activeSubscriptions = append(activeSubscriptions, subscription)
		}
	}

	return activeSubscriptions, nil
}

func (c *ConsoleClient) Subscriptions(ctx context.Context, registryID string) ([]entity.ClientSubscription, error) {
	result := &entity.GetClientSubscriptionsResp{}
	_, err := c.client.R().
		SetContext(ctx).
		SetResult(result).
		Get(fmt.Sprintf("/v1/automations/executor/%s/subscriptions", registryID))

	if err != nil {
		return nil, fmt.Errorf("failed to get executor subscriptions: %w", err)
	}

	return result.Data, nil
}

func (c *ConsoleClient) Execute(ctx context.Context, req *entity.ExecuteTaskReq) (*entity.ExecuteTaskResp, error) {
	logger := log.NewLogger("console-executor", "debug")
	result := &entity.ExecuteTaskResp{}
	raw, _ := json.Marshal(req)
	logger.Debug("executor request", log.Str("req", string(raw)))
	resp, err := c.client.R().
		SetContext(ctx).
		SetBody(req).
		Post(fmt.Sprintf("/v1/automations/tasks/execute/%d", req.ChainID))
	logger.Debug("executor result", log.Str("resp", string(resp.Body())))
	ctxDone := false
	select {
	case <-ctx.Done():
		ctxDone = true
	default:
	}
	logger.Debug("execution context status", log.Str("ctx", fmt.Sprintf("%t", ctxDone)))
	if err != nil {
		return nil, fmt.Errorf("failed to get executor subscriptions: %w", err)
	}
	if err = json.Unmarshal(resp.Body(), result); err != nil {
		return nil, err
	}

	return result, nil
}
