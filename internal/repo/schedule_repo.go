package repo

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Brahma-fi/brahma-builder/internal/entity"
	"github.com/ethereum/go-ethereum/common"
	"go.temporal.io/sdk/client"
	"google.golang.org/protobuf/encoding/protojson"
)

type pbMemo struct {
	Fields struct {
		Params struct {
			Metadata struct {
				Encoding string `json:"encoding"`
			} `json:"metadata"`
			Data string `json:"data"`
		} `json:"params"`
		Schedule struct {
			Metadata struct {
				Encoding string `json:"encoding"`
			} `json:"metadata"`
			Data string `json:"data"`
		} `json:"schedule"`
	} `json:"fields"`
}

const (
	_maxPageSize       = 100
	KeyExecutionStatus = "ExecutionStatus"
)

type ScheduleRepo struct {
	client client.Client
}

func NewSchedulesRepo(client client.Client) *ScheduleRepo {
	return &ScheduleRepo{
		client: client,
	}
}

func (s *ScheduleRepo) BySubAccountAddressChainIDAndStatus(
	ctx context.Context,
	subAccount common.Address,
	chainID int64,
) ([]entity.Schedule, error) {
	query := fmt.Sprintf("%s = '%s' AND %s = %d",
		entity.SearchAttrKeySubAccountAddress, subAccount.Hex(),
		entity.SearchAttrKeyChainID, chainID)

	return s.listSchedules(ctx, query)
}

func (s *ScheduleRepo) BySubAccountAddressesChainIDAndStatus(
	ctx context.Context,
	subAccounts []common.Address,
	chainID int64,
) ([]entity.Schedule, error) {
	if len(subAccounts) == 0 {
		return nil, nil
	}

	return s.listSchedulesWithAdvancedQuery(ctx, subAccounts, chainID)
}

func (s *ScheduleRepo) listSchedulesWithAdvancedQuery(
	ctx context.Context,
	subAccounts []common.Address,
	chainID int64,
) ([]entity.Schedule, error) {
	subAccStr := make([]string, len(subAccounts))
	for i, subacc := range subAccounts {
		subAccStr[i] = fmt.Sprintf("'%s'", subacc.Hex())
	}

	query := fmt.Sprintf("%s IN (%s) AND %s = %d",
		entity.SearchAttrKeySubAccountAddress, strings.Join(subAccStr, ","),
		entity.SearchAttrKeyChainID, chainID)

	return s.listSchedules(ctx, query)
}

func (s *ScheduleRepo) filterSchedules(
	schedules []entity.Schedule,
	subAccounts []common.Address,
	chainID int64,
) []entity.Schedule {
	filterMap := make(map[common.Address]struct{}, len(subAccounts))
	for _, subacc := range subAccounts {
		filterMap[subacc] = struct{}{}
	}

	filteredList := make([]entity.Schedule, 0, len(schedules))
	for _, schedule := range schedules {
		if _, ok := filterMap[schedule.Config.Params.SubAccountAddress]; ok && schedule.Config.Params.ChainID == chainID {
			filteredList = append(filteredList, schedule)
		}
	}

	return filteredList
}

func (s *ScheduleRepo) listSchedulesStandard(
	ctx context.Context,
	subAccounts []common.Address,
	chainID int64,
) ([]entity.Schedule, error) {
	allSchedules, err := s.listSchedules(ctx, "")
	if err != nil {
		return nil, fmt.Errorf("failed to list schedules: %w", err)
	}

	filterMap := make(map[common.Address]struct{}, len(subAccounts))
	for _, subacc := range subAccounts {
		filterMap[subacc] = struct{}{}
	}

	filteredList := make([]entity.Schedule, 0, len(allSchedules))
	for _, schedule := range allSchedules {
		if _, ok := filterMap[schedule.Config.Params.SubAccountAddress]; ok && schedule.Config.Params.ChainID == chainID {
			filteredList = append(filteredList, schedule)
		}
	}

	return filteredList, nil
}

func (s *ScheduleRepo) listSchedules(ctx context.Context, query string) ([]entity.Schedule, error) {
	resp, err := s.client.ScheduleClient().List(ctx, client.ScheduleListOptions{
		PageSize: _maxPageSize,
		Query:    query,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list schedules: %w", err)
	}

	var schedules []entity.Schedule
	for resp.HasNext() {
		entry, err := resp.Next()
		if err != nil {
			return nil, fmt.Errorf("failed to get next schedule: %w", err)
		}

		schedule, err := parseScheduleEntry(entry)
		if err != nil {
			return nil, err
		}

		schedules = append(schedules, schedule)
	}

	return schedules, nil
}

func parseScheduleEntry(entry *client.ScheduleListEntry) (entity.Schedule, error) {
	memo := entity.ExecuteWorkflowParams{}
	if entry.Memo != nil {
		var t pbMemo
		jsonBytes, err := protojson.Marshal(entry.Memo)
		if err != nil {
			return entity.Schedule{}, fmt.Errorf("failed to marshal memo to JSON: %w", err)
		}
		if err := json.Unmarshal(jsonBytes, &t); err != nil {
			return entity.Schedule{}, fmt.Errorf("failed to unmarshal JSON to T struct: %w", err)
		}

		if err := decodeMemoField(t.Fields.Params.Data, &memo.Params); err != nil {
			return entity.Schedule{}, fmt.Errorf("failed to decode params: %w", err)
		}

		if err := decodeMemoField(t.Fields.Schedule.Data, &memo.Schedule); err != nil {
			return entity.Schedule{}, fmt.Errorf("failed to decode schedule: %w", err)
		}
	}

	return entity.Schedule{
		Config:     memo,
		ScheduleID: entry.ID,
		CreatedAt:  entry.Spec.StartAt,
	}, nil
}

func decodeMemoField(data string, target interface{}) error {
	decodedData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return fmt.Errorf("failed to decode base64 data: %w", err)
	}
	if err := json.Unmarshal(decodedData, target); err != nil {
		return fmt.Errorf("failed to unmarshal data: %w", err)
	}
	return nil
}
