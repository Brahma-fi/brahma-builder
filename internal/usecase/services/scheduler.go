package services

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/Brahma-fi/brahma-builder/internal/entity"
	"github.com/Brahma-fi/brahma-builder/internal/usecase/workflows"
	"github.com/Brahma-fi/brahma-builder/pkg/log"
	"github.com/Brahma-fi/brahma-builder/pkg/utils"
	"github.com/ethereum/go-ethereum/common"
	"go.temporal.io/api/enums/v1"
	"go.temporal.io/api/operatorservice/v1"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
)

type Scheduler struct {
	client        client.Client
	console       console
	executors     executors
	schedulesRepo schedulesRepo
}

func NewScheduler(
	ctx context.Context,
	client client.Client,
	console console,
	executors executors,
	schedulesRepo schedulesRepo,
) (*Scheduler, error) {
	logger := log.NewLogger("scheduler", "debug")
	// ensure search attributes are always ready
	resp, err := client.OperatorService().AddSearchAttributes(ctx, &operatorservice.AddSearchAttributesRequest{
		SearchAttributes: map[string]enums.IndexedValueType{
			entity.SearchAttrKeyChainID:           enums.INDEXED_VALUE_TYPE_INT,
			entity.SearchAttrKeySubAccountAddress: enums.INDEXED_VALUE_TYPE_KEYWORD,
			entity.SearchAttrKeyExecutorAddress:   enums.INDEXED_VALUE_TYPE_KEYWORD,
			entity.SearchAttrKeyExecutorID:        enums.INDEXED_VALUE_TYPE_KEYWORD,
		},
		Namespace: entity.DefaultNamespace,
	})
	if err != nil && !strings.Contains(err.Error(), "already exists") {
		return nil, err
	}

	logger.Info("AddSearchAttributes", log.Str("resp", resp.String()))

	return &Scheduler{client: client, console: console, executors: executors, schedulesRepo: schedulesRepo}, nil
}

func (s *Scheduler) Run(ctx context.Context, config entity.ExecuteWorkflowParams) (string, error) {
	o := workflows.Orchestrator{}

	config.Schedule.ID = config.Params.ID()

	searchAttributes := temporal.NewSearchAttributes(
		temporal.NewSearchAttributeKeyKeyword(entity.SearchAttrKeySubAccountAddress).ValueSet(config.Params.SubAccountAddress.Hex()),
		temporal.NewSearchAttributeKeyKeyword(entity.SearchAttrKeyExecutorAddress).ValueSet(config.Params.ExecutorAddress.Hex()),
		temporal.NewSearchAttributeKeyInt64(entity.SearchAttrKeyChainID).ValueSet(config.Params.ChainID),
		temporal.NewSearchAttributeKeyKeyword(entity.SearchAttrKeyExecutorID).ValueSet(config.Params.ExecutorID),
	)

	memo, err := utils.Struct2Map(config)
	if err != nil {
		return "", err
	}

	schedule, err := s.client.
		ScheduleClient().Create(ctx, client.ScheduleOptions{
		ID: config.Schedule.ID,
		Spec: client.ScheduleSpec{
			Intervals: []client.ScheduleIntervalSpec{
				{
					Every: config.Schedule.Every,
				},
			},
		},
		Action: &client.ScheduleWorkflowAction{
			TaskQueue:             entity.BaseTaskQueue,
			Workflow:              o.OrchestratorWorkflow,
			Args:                  []interface{}{config},
			TypedSearchAttributes: searchAttributes,
		},
		Overlap:               enums.SCHEDULE_OVERLAP_POLICY_SKIP,
		PauseOnFailure:        false,
		Note:                  string(config.Params.Subscription.Metadata),
		TypedSearchAttributes: searchAttributes,
		Memo:                  memo,
		TriggerImmediately:    true,
	})
	if err != nil {
		return "", err
	}

	return schedule.GetID(), nil
}

func (s *Scheduler) Sync(ctx context.Context) error {
	activeAccounts, executorMetadata, err := s.fetchActiveAccountsAndMetadata(ctx)
	if err != nil {
		return err
	}

	accountsByChain := groupAccountsByChain(activeAccounts)

	for chainID, accounts := range accountsByChain {
		if err := s.syncChain(ctx, chainID, accounts, executorMetadata); err != nil {
			return err
		}
	}

	return nil
}

func (s *Scheduler) fetchActiveAccountsAndMetadata(ctx context.Context) ([]entity.ClientSubscription, map[string]*entity.ExecutorMetadata, error) {
	registeredExecutors := s.executors.List()
	activeAccounts := make([]entity.ClientSubscription, 0)
	executorMetadata := make(map[string]*entity.ExecutorMetadata)

	for _, e := range registeredExecutors {
		metadata, err := s.console.ExecutorByAddressAndChainID(ctx, common.HexToAddress(e.Address), uint64(e.ChainID))
		if err != nil {
			return nil, nil, err
		}

		executorMetadata[metadata.Id] = metadata
		subscription, err := s.console.Subscriptions(ctx, metadata.Id)
		if err != nil {
			return nil, nil, err
		}

		activeAccounts = append(activeAccounts, subscription...)
	}

	return activeAccounts, executorMetadata, nil
}

func groupAccountsByChain(accounts []entity.ClientSubscription) map[int64][]entity.ClientSubscription {
	accountsByChain := make(map[int64][]entity.ClientSubscription)
	for _, account := range accounts {
		chainID := int64(account.ChainId)
		accountsByChain[chainID] = append(accountsByChain[chainID], account)
	}
	return accountsByChain
}

func (s *Scheduler) syncChain(
	ctx context.Context,
	chainID int64,
	accounts []entity.ClientSubscription,
	executorMetadata map[string]*entity.ExecutorMetadata,
) error {
	logger := log.GetLogger(ctx)
	logger.Info("syncing chain", log.Int("chainID", int(chainID)), log.Any("account", accounts))
	subAccounts := extractSubAccounts(accounts)

	existingSchedules, err := s.schedulesRepo.BySubAccountAddressesChainIDAndStatus(ctx, subAccounts, chainID)
	if err != nil {
		return err
	}

	existingScheduleMap := createExistingScheduleMap(ctx, existingSchedules)
	if err := s.createNewSchedules(ctx, accounts, executorMetadata, existingScheduleMap, chainID); err != nil {
		return err
	}

	return s.terminateCancelledSchedules(ctx, existingSchedules, accounts)
}

func extractSubAccounts(accounts []entity.ClientSubscription) []common.Address {
	subAccounts := make([]common.Address, len(accounts))
	for i, account := range accounts {
		subAccounts[i] = common.HexToAddress(account.SubAccountAddress)
	}
	return subAccounts
}

func createExistingScheduleMap(ctx context.Context, existingSchedules []entity.Schedule) map[common.Address]bool {
	logger := log.GetLogger(ctx)
	existingScheduleMap := make(map[common.Address]bool)
	for _, schedule := range existingSchedules {
		logger.Info(
			"existing schedule",
			log.Str("id", schedule.ScheduleID),
			log.Str("subacc", schedule.Config.Params.SubAccountAddress.Hex()),
		)
		existingScheduleMap[schedule.Config.Params.SubAccountAddress] = true
	}
	return existingScheduleMap
}

func (s *Scheduler) createNewSchedules(
	ctx context.Context,
	accounts []entity.ClientSubscription,
	executorMetadata map[string]*entity.ExecutorMetadata,
	existingScheduleMap map[common.Address]bool,
	chainID int64,
) error {
	logger := log.GetLogger(ctx)
	for _, account := range accounts {
		subAccountAddress := common.HexToAddress(account.SubAccountAddress)
		if !existingScheduleMap[subAccountAddress] && account.Status == 2 {
			config, err := s.createWorkflowConfig(account, executorMetadata[account.RegistryId], chainID)
			if err != nil {
				return err
			}
			logger.Info(
				"creating new schedule",
				log.Str("scheduleID", config.Schedule.ID),
				log.Str("subaccount", account.SubAccountAddress),
			)
			if _, err = s.Run(ctx, config); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *Scheduler) createWorkflowConfig(
	account entity.ClientSubscription,
	metadata *entity.ExecutorMetadata,
	chainID int64,
) (entity.ExecuteWorkflowParams, error) {
	// Get the executor configuration
	cfg, err := s.executors.Config(common.HexToAddress(metadata.Executor))
	if err != nil {
		return entity.ExecuteWorkflowParams{}, err
	}

	// Parse the default duration from the executor config
	duration, err := time.ParseDuration(cfg.Every)
	if err != nil {
		return entity.ExecuteWorkflowParams{}, err
	}

	// Check for custom interval options in the account metadata
	custom := &entity.CustomIntervalOptions{}
	if err = json.Unmarshal(account.Metadata, custom); err == nil && custom.Every != "" {
		// Try to parse custom.Every as a duration string
		customDuration, err := time.ParseDuration(custom.Every)
		if err == nil {
			// If parsing as duration succeeds, use the custom duration
			duration = customDuration
		} else {
			// If parsing as duration fails, try to parse as an integer (seconds)
			seconds, err := strconv.Atoi(custom.Every)
			if err == nil {
				duration = time.Duration(seconds) * time.Second
			}
			// If both parsing attempts fail, we keep the default duration
		}
	}

	// Create and return the ExecuteWorkflowParams
	return entity.ExecuteWorkflowParams{
		Params: entity.OrchestratorParams{
			ExecutorAddress:   common.HexToAddress(metadata.Executor),
			SubAccountAddress: common.HexToAddress(account.SubAccountAddress),
			ExecutorID:        account.RegistryId,
			ChainID:           chainID,
			Subscription:      account,
		},
		Schedule: &entity.ScheduledWorkflowConfig{
			Every: duration,
		},
	}, nil
}

func (s *Scheduler) terminateCancelledSchedules(
	ctx context.Context,
	existingSchedules []entity.Schedule,
	accounts []entity.ClientSubscription,
) error {
	logger := log.GetLogger(ctx)
	activeSubAccounts := make(map[common.Address]entity.ClientSubscription)
	for _, account := range accounts {
		activeSubAccounts[common.HexToAddress(account.SubAccountAddress)] = account
	}

	for _, schedule := range existingSchedules {
		if sub, ok := activeSubAccounts[schedule.Config.Params.SubAccountAddress]; !ok || sub.Status == 4 {
			logger.Info(
				"terminating schedule",
				log.Str("scheduleID", schedule.ScheduleID),
				log.Str("subaccount", schedule.Config.Params.SubAccountAddress.Hex()),
			)
			if err := s.client.ScheduleClient().GetHandle(ctx, schedule.ScheduleID).Delete(ctx); err != nil {
				return err
			}
		}
	}

	return nil
}
