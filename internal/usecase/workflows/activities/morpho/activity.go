package morpho

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"slices"
	"strings"

	"github.com/Brahma-fi/brahma-builder/internal/entity"
	utils "github.com/Brahma-fi/brahma-builder/pkg/utils/abis/erc20"
	"github.com/Brahma-fi/go-safe/encoders"
	safetypes "github.com/Brahma-fi/go-safe/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/google/uuid"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type ReBalancingStrategy struct {
	client         morphoClient
	executor       consoleExecutor
	config         *Config
	logsRepo       executionsLogRepo
	bundlerAddress common.Address
	caller         bind.ContractCaller
	oracle         pricingOracle
}

func NewReBalancingStrategy(
	client morphoClient,
	executor consoleExecutor,
	caller bind.ContractCaller,
	logsRepo executionsLogRepo,
	config *Config,
	oracle pricingOracle,
) (*ReBalancingStrategy, error) {
	return &ReBalancingStrategy{
		client:         client,
		caller:         caller,
		executor:       executor,
		config:         config,
		logsRepo:       logsRepo,
		bundlerAddress: common.HexToAddress(config.BundlerAddress),
		oracle:         oracle,
	}, nil
}

func (m *ReBalancingStrategy) ExecutionHandler(
	ctx context.Context,
	execCtx entity.ExecCtx,
) error {
	logger := activity.GetLogger(ctx)
	params := &StrategyParams{}
	if err := json.Unmarshal(execCtx.Params.Subscription.Metadata, params); err != nil {
		return fmt.Errorf("failed to unmarshal strategy params: %w", err)
	}

	initialState, err := m.getInitialState(ctx, execCtx, params, execCtx.Params.ChainID)
	if err != nil {
		return fmt.Errorf("failed to get initial state: %w", err)
	}

	bestVault, _ := m.findBestVault(initialState.vaults, initialState.minUnderlyingLiquidity)

	if bestVault == initialState.currentVault {
		logger.Info("No re-balance signal")
		return nil
	}

	if len(m.config.WhitelistedVaults) != 0 && !slices.Contains(m.config.WhitelistedVaults, bestVault.Hex()) {
		return fmt.Errorf("vault not whitelisted %s", bestVault.Hex())
	}

	if !initialState.isAlreadyInVault && initialState.hasAvailableBalance {
		return m.handleDeposit(ctx, logger, execCtx, initialState.subaccount, bestVault, params, int64(execCtx.Params.Subscription.ChainId))
	}

	if initialState.isAlreadyInVault && bestVault != initialState.currentVault {
		return m.handleRebalance(ctx, logger, execCtx, initialState.subaccount, initialState.currentVault, bestVault, int64(execCtx.Params.Subscription.ChainId), params)
	}

	return nil
}

type State struct {
	vaults                 []entity.VaultInfo
	subaccount             common.Address
	currentVault           common.Address
	subAccBalance          *big.Int
	currentVaultBalance    *big.Int
	minUnderlyingLiquidity *big.Int
	hasAvailableBalance    bool
	isAlreadyInVault       bool
}

func (m *ReBalancingStrategy) getInitialState(
	ctx context.Context,
	execCtx entity.ExecCtx,
	params *StrategyParams,
	chainID int64,
) (*State, error) {
	vaults, err := m.client.Vaults(ctx, params.BaseToken, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to get vaults: %w", err)
	}

	subaccount := common.HexToAddress(execCtx.ExecuteWorkflowParams.Params.Subscription.SubAccountAddress)
	currentVault, currentBalance, err := m.ActiveVault(ctx, subaccount, vaults)
	if err != nil {
		return nil, fmt.Errorf("failed to get active vault: %w", err)
	}

	subAccBalance, err := m.getSubAccountBalance(ctx, params.BaseToken, subaccount)
	if err != nil {
		return nil, fmt.Errorf("failed to get subaccount balance: %w", err)
	}

	minUnderlyingLiquidity := new(big.Int).Set(subAccBalance)
	if currentBalance != nil && subAccBalance.Cmp(currentBalance) < 0 {
		minUnderlyingLiquidity.Set(currentBalance)
	}

	return &State{
		vaults:                 vaults,
		subaccount:             subaccount,
		currentVault:           currentVault,
		subAccBalance:          subAccBalance,
		currentVaultBalance:    currentBalance,
		minUnderlyingLiquidity: minUnderlyingLiquidity,
		hasAvailableBalance:    subAccBalance.Cmp(big.NewInt(0)) > 0,
		isAlreadyInVault:       currentVault != common.Address{},
	}, nil
}

func (m *ReBalancingStrategy) getSubAccountBalance(
	ctx context.Context,
	baseToken common.Address,
	subaccount common.Address,
) (*big.Int, error) {
	baseTokenCaller, err := utils.NewErc20Caller(baseToken, m.caller)
	if err != nil {
		return nil, fmt.Errorf("failed to create base token caller: %w", err)
	}

	balance, err := baseTokenCaller.BalanceOf(&bind.CallOpts{Context: ctx}, subaccount)
	if err != nil {
		return nil, fmt.Errorf("failed to get balance: %w", err)
	}

	return balance, nil
}

func (m *ReBalancingStrategy) handleDeposit(
	ctx context.Context,
	logger log.Logger,
	execCtx entity.ExecCtx,
	subaccount, bestVault common.Address,
	params *StrategyParams,
	chainID int64,
) error {
	logger.Info("Entering strategy", "address", bestVault.String())
	_, err := m.Deposit(ctx, logger, subaccount, bestVault, params, chainID)
	if err != nil {
		return fmt.Errorf("failed to deposit: %w", err)
	}

	return nil
}

func (m *ReBalancingStrategy) handleRebalance(
	ctx context.Context,
	logger log.Logger,
	execCtx entity.ExecCtx,
	subaccount, currentVault, bestVault common.Address,
	chainID int64,
	params *StrategyParams,
) error {
	logger.Info("Re-balance strategy", "from", currentVault.String(), "to", bestVault.String())
	subID, err := uuid.Parse(execCtx.Params.Subscription.Id)
	if err != nil {
		return fmt.Errorf("failed to parse subscription ID: %w", err)
	}

	_, err = m.RedeemAndDeposit(ctx, logger, subID, subaccount, currentVault, bestVault, chainID, params)
	if err != nil {
		return fmt.Errorf("failed to redeem and deposit: %w", err)
	}

	return nil
}

func (m *ReBalancingStrategy) Deposit(
	ctx context.Context,
	logger log.Logger,
	user, vault common.Address,
	params *StrategyParams,
	chainID int64,
) (*ExecutionLog, error) {
	_, depositAmount, baseFees, err := m.calculateDepositAmountsAndFees(ctx, user, params, chainID)
	if err != nil {
		return nil, err
	}

	transactions, err := m.prepareDepositTransactions(ctx, user, vault, depositAmount, baseFees, params)
	if err != nil {
		return nil, err
	}

	return m.executeDeposit(ctx, logger, user, vault, depositAmount, baseFees, transactions, chainID)
}

func (m *ReBalancingStrategy) calculateDepositAmountsAndFees(
	ctx context.Context,
	user common.Address,
	params *StrategyParams,
	chainID int64,
) (*big.Int, *big.Int, *big.Int, error) {
	balance, err := m.getSubAccountBalance(ctx, params.BaseToken, user)
	if err != nil {
		return nil, nil, nil, err
	}

	baseFeeAmt, err := m.calculateBaseFee(ctx, params.BaseToken, chainID)
	if err != nil {
		return nil, nil, nil, err
	}

	if err = m.validateBalance(balance, baseFeeAmt); err != nil {
		return nil, nil, nil, err
	}

	depositAmount := new(big.Int).Sub(balance, baseFeeAmt)
	return balance, depositAmount, baseFeeAmt, nil
}

func (m *ReBalancingStrategy) prepareDepositTransactions(
	ctx context.Context,
	user, vault common.Address,
	depositAmount, baseFees *big.Int,
	params *StrategyParams,
) ([]safetypes.Transaction, error) {
	transferFeeTxn, err := m.prepareTransferFeeTxn(baseFees, params.BaseToken)
	if err != nil {
		return nil, err
	}

	approveTxn, err := m.prepareApproveTxn(depositAmount, params.BaseToken)
	if err != nil {
		return nil, err
	}

	depositTxn, err := m.prepareDepositTxn(ctx, user, vault, depositAmount, params.BaseToken)
	if err != nil {
		return nil, err
	}

	return []safetypes.Transaction{transferFeeTxn, approveTxn, depositTxn}, nil
}

func (m *ReBalancingStrategy) executeDeposit(
	ctx context.Context,
	logger log.Logger,
	user, vault common.Address,
	depositAmount, baseFees *big.Int,
	transactions []safetypes.Transaction,
	chainID int64,
) (*ExecutionLog, error) {
	safeTx, err := encoders.GetEncodedSafeTx(
		common.Address{},
		common.HexToAddress(entity.SafeMultiSendCallOnly),
		&entity.SafeMultiSendABI,
		transactions,
		chainID,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to encode safe transaction: %w", err)
	}

	req := &entity.SignAndExecuteRequest{
		Subaccount: user.Hex(),
		ChainID:    uint64(chainID),
		Operation:  safeTx.Operation,
		To:         safeTx.To.String(),
		Value:      safeTx.Value.String(),
		Data:       safeTx.Data.String(),
	}

	taskID, err := m.executor.Execute(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute transaction: %w", err)
	}

	logger.Info("Executed strategy signal", "taskID", taskID)
	return &ExecutionLog{
		Message: fmt.Sprintf("Entered into strategy %s", vault.Hex()),
		Metadata: ExecutionMetadata{
			TaskID: taskID,
			Req:    req,
			TransitionState: TransitionState{
				Current: AutomationState{
					TargetVault:    vault,
					InputAmount:    depositAmount.String(),
					FeesAmount:     baseFees.String(),
					GeneratedYield: "0",
				},
				Prev: nil,
			},
		},
	}, nil
}

func (m *ReBalancingStrategy) RedeemAndDeposit(
	ctx context.Context,
	logger log.Logger,
	subID uuid.UUID,
	user, from, to common.Address,
	chainID int64,
	params *StrategyParams,
) (*ExecutionLog, error) {
	latest, err := m.logsRepo.LatestBySubID(ctx, subID)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest logs: %w", err)
	}

	metadata := &ExecutionMetadata{}
	if err = json.Unmarshal(latest.Metadata.([]uint8), metadata); err != nil {
		return nil, fmt.Errorf("failed to unmarshal metadata: %w", err)
	}

	balance, err := m.client.PreviewRedeem(ctx, from, user)
	if err != nil {
		return nil, fmt.Errorf("failed to preview redeem: %w", err)
	}

	transactions, depositAmount, baseFees, yield, err := m.prepareRedeemAndDepositTransactions(
		ctx,
		user,
		from,
		to,
		balance,
		metadata,
		params,
		chainID,
	)
	if err != nil {
		return nil, err
	}

	return m.executeRedeemAndDeposit(ctx, logger, user, from, to, depositAmount, baseFees, yield, transactions, metadata.TransitionState.Current, chainID)
}

func (m *ReBalancingStrategy) prepareRedeemAndDepositTransactions(
	ctx context.Context,
	user, from, to common.Address,
	balance *big.Int,
	metadata *ExecutionMetadata,
	params *StrategyParams,
	chainID int64,
) ([]safetypes.Transaction, *big.Int, *big.Int, *big.Int, error) {
	redeemTxn, err := m.prepareRedeemTxn(ctx, from, user)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	baseFees, depositAmount, yield, err := m.calculateRedeemAndDepositAmounts(ctx, balance, metadata, params, chainID)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	approveTxn, err := m.prepareApproveTxn(depositAmount, params.BaseToken)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	reBalanceTxn, err := m.prepareReBalanceTxn(ctx, user, to, depositAmount, params.BaseToken)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	transferFeeTxn, err := m.prepareTransferFeeTxn(baseFees, params.BaseToken)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return []safetypes.Transaction{
		redeemTxn,
		approveTxn,
		reBalanceTxn,
		transferFeeTxn,
	}, depositAmount, baseFees, yield, nil
}

func (m *ReBalancingStrategy) executeRedeemAndDeposit(
	ctx context.Context,
	logger log.Logger,
	user, from, to common.Address,
	depositAmount, baseFees, yield *big.Int,
	transactions []safetypes.Transaction,
	prevState AutomationState,
	chainID int64,
) (*ExecutionLog, error) {
	safeTx, err := encoders.GetEncodedSafeTx(
		common.Address{},
		common.HexToAddress(entity.SafeMultiSendCallOnly),
		&entity.SafeMultiSendABI,
		transactions,
		chainID,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to encode safe transaction: %w", err)
	}

	req := &entity.SignAndExecuteRequest{
		Subaccount: user.Hex(),
		ChainID:    uint64(chainID),
		Operation:  safeTx.Operation,
		To:         safeTx.To.String(),
		Value:      safeTx.Value.String(),
		Data:       safeTx.Data.String(),
	}

	taskID, err := m.executor.Execute(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute transaction: %w", err)
	}

	logger.Info("Executed strategy signal", "taskID", taskID)
	return &ExecutionLog{
		Message: fmt.Sprintf("Rebalanced shares from %s to %s", from.Hex(), to.Hex()),
		Metadata: ExecutionMetadata{
			TaskID: taskID,
			Req:    req,
			TransitionState: TransitionState{
				Current: AutomationState{
					TargetVault:    to,
					InputAmount:    depositAmount.String(),
					FeesAmount:     baseFees.String(),
					GeneratedYield: yield.String(),
				},
				Prev: &prevState,
			},
		},
	}, nil
}

func (m *ReBalancingStrategy) findBestVault(
	vaults []entity.VaultInfo,
	minLiquidity *big.Int,
) (common.Address, float64) {
	var bestVault common.Address
	var bestApy float64

	for _, vault := range vaults {
		if vault.State.NetApy > bestApy && vault.Liquidity.Underlying.Cmp(minLiquidity) > 0 {
			bestApy = vault.State.NetApy
			bestVault = common.HexToAddress(vault.Address)
		}
	}

	return bestVault, bestApy
}

func (m *ReBalancingStrategy) ActiveVault(
	ctx context.Context,
	user common.Address,
	vaults []entity.VaultInfo,
) (common.Address, *big.Int, error) {
	for _, vault := range vaults {
		shares, err := m.client.Shares(ctx, common.HexToAddress(vault.Address), user)
		if err != nil {
			return common.Address{}, nil, fmt.Errorf("failed to get shares: %w", err)
		}

		if shares.Int64() != 0 {
			balance, err := m.client.PreviewRedeem(ctx, common.HexToAddress(vault.Address), user)
			if err != nil {
				return common.Address{}, nil, fmt.Errorf("failed to preview redeem: %w", err)
			}
			return common.HexToAddress(vault.Address), balance, nil
		}
	}
	return common.Address{}, nil, nil
}

func (m *ReBalancingStrategy) prepareTransferFeeTxn(
	baseFees *big.Int,
	baseToken common.Address,
) (*entity.Transaction, error) {
	erc20ABI, err := abi.JSON(strings.NewReader(utils.Erc20MetaData.ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse ERC20 ABI: %w", err)
	}

	transferFeeCallData, err := erc20ABI.Pack("transfer", common.HexToAddress(m.config.FeeReceiver), baseFees)
	if err != nil {
		return nil, fmt.Errorf("failed to pack transfer fee call data: %w", err)
	}

	return &entity.Transaction{
		Target: baseToken,
		Val:    new(big.Int).SetInt64(0),
		Data:   common.Bytes2Hex(transferFeeCallData),
	}, nil
}

func (m *ReBalancingStrategy) prepareApproveTxn(
	amount *big.Int,
	baseToken common.Address,
) (*entity.Transaction, error) {
	erc20ABI, err := abi.JSON(strings.NewReader(utils.Erc20MetaData.ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse ERC20 ABI: %w", err)
	}

	approveCallData, err := erc20ABI.Pack("approve", m.bundlerAddress, amount)
	if err != nil {
		return nil, fmt.Errorf("failed to pack approve call data: %w", err)
	}

	return &entity.Transaction{
		Target: baseToken,
		Val:    new(big.Int).SetInt64(0),
		Data:   common.Bytes2Hex(approveCallData),
	}, nil
}

func (m *ReBalancingStrategy) prepareDepositTxn(
	ctx context.Context,
	user, vault common.Address,
	depositAmount *big.Int,
	baseToken common.Address,
) (*entity.Transaction, error) {
	logger := activity.GetLogger(ctx)
	minShares, err := m.client.PreviewDeposit(ctx, vault, depositAmount)
	if err != nil {
		return nil, fmt.Errorf("failed to preview deposit: %w", err)
	}

	minSharesFloat := new(big.Float).SetInt(minShares)
	out, _ := new(big.Float).Mul(minSharesFloat, new(big.Float).SetFloat64(slippage)).Int(nil)
	logger.Info("calculated min deposit shares",
		"vault", vault.Hex(),
		"input", depositAmount.String(),
		"share", out.String(),
	)
	bundlerMultiCallData, err := m.client.Bundle([]entity.BundlerCall{
		{
			Type:   entity.BundlerCallTransferFrom,
			Params: []any{baseToken, depositAmount},
		},
		{
			Type:   entity.BundlerCallDeposit,
			Params: []any{vault, depositAmount, out, user},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to bundle calls: %w", err)
	}
	logger.Info("deposit callData",
		"val", hexutil.Encode(bundlerMultiCallData),
	)
	return &entity.Transaction{
		Target: m.bundlerAddress,
		Val:    new(big.Int).SetInt64(0),
		Data:   common.Bytes2Hex(bundlerMultiCallData),
	}, nil
}

func (m *ReBalancingStrategy) prepareRedeemTxn(
	ctx context.Context,
	from, user common.Address,
) (*entity.Transaction, error) {
	redeemCallData, err := m.client.RedeemMax(ctx, from, user)
	if err != nil {
		return nil, fmt.Errorf("failed to get redeem max call data: %w", err)
	}

	return &entity.Transaction{
		Target: from,
		Val:    new(big.Int).SetInt64(0),
		Data:   common.Bytes2Hex(redeemCallData),
	}, nil
}

func (m *ReBalancingStrategy) calculateRedeemAndDepositAmounts(
	ctx context.Context,
	balance *big.Int,
	metadata *ExecutionMetadata,
	params *StrategyParams,
	chainID int64,
) (*big.Int, *big.Int, *big.Int, error) {
	baseFeeAmt, err := m.calculateBaseFee(ctx, params.BaseToken, chainID)
	if err != nil {
		return nil, nil, nil, err
	}
	if err = m.validateBalance(balance, baseFeeAmt); err != nil {
		return nil, nil, nil, err
	}
	yield, err := m.calculateYield(balance, metadata)
	if err != nil {
		return nil, nil, nil, err
	}

	baseFeeAmt = m.addYieldFees(baseFeeAmt, yield)
	depositAmount := new(big.Int).Sub(balance, baseFeeAmt)
	return baseFeeAmt, depositAmount, yield, nil
}

func (m *ReBalancingStrategy) prepareReBalanceTxn(
	ctx context.Context,
	user, to common.Address,
	depositAmount *big.Int,
	baseToken common.Address,
) (*entity.Transaction, error) {
	minSharesIn, err := m.client.PreviewDeposit(ctx, to, depositAmount)
	if err != nil {
		return nil, fmt.Errorf("failed to preview deposit: %w", err)
	}

	minSharesInFloat := new(big.Float).SetInt(minSharesIn)
	in, _ := new(big.Float).Mul(minSharesInFloat, new(big.Float).SetFloat64(slippage)).Int(nil)

	bundlerMultiCallData, err := m.client.Bundle([]entity.BundlerCall{
		{
			Type:   entity.BundlerCallTransferFrom,
			Params: []any{baseToken, depositAmount},
		},
		{
			Type:   entity.BundlerCallDeposit,
			Params: []any{to, depositAmount, in, user},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to bundle calls: %w", err)
	}

	return &entity.Transaction{
		Target: m.bundlerAddress,
		Val:    new(big.Int).SetInt64(0),
		Data:   common.Bytes2Hex(bundlerMultiCallData),
	}, nil
}

func (m *ReBalancingStrategy) calculateBaseFee(
	ctx context.Context,
	baseToken common.Address,
	chainID int64,
) (*big.Int, error) {
	exactFeeAmt, ok := m.config.FeeConfig[baseToken.Hex()]
	if ok {
		return m.parseExactFee(exactFeeAmt)
	}

	// TODO: for mainnet calculate using current gasPrice and take gas units in config instead
	return m.convertUSDFeeToToken(ctx, baseToken, chainID)
}

func (m *ReBalancingStrategy) parseExactFee(exactFeeAmt string) (*big.Int, error) {
	baseFeeAmt, ok := new(big.Int).SetString(exactFeeAmt, 10)
	if !ok {
		return nil, fmt.Errorf("failed to parse base fee %s", exactFeeAmt)
	}
	return baseFeeAmt, nil
}

func (m *ReBalancingStrategy) convertUSDFeeToToken(
	ctx context.Context,
	baseToken common.Address,
	chainID int64,
) (*big.Int, error) {
	return m.oracle.ConvertUSDToToken(ctx, chainID, m.config.BaseFeesInUSD, baseToken)
}

func (m *ReBalancingStrategy) validateBalance(balance, baseFeeAmt *big.Int) error {
	if balance.Cmp(baseFeeAmt) <= 0 {
		return fmt.Errorf("input does not cover base fees want=%s have=%s", baseFeeAmt.String(), balance.String())
	}
	return nil
}

func (m *ReBalancingStrategy) calculateYield(balance *big.Int, metadata *ExecutionMetadata) (*big.Int, error) {
	vaultDepositAmount, ok := new(big.Int).SetString(metadata.TransitionState.Current.InputAmount, 10)
	if !ok {
		return nil, fmt.Errorf("failed to parse vault deposit amount %s", metadata.TransitionState.Current.InputAmount)
	}
	return new(big.Int).Sub(balance, vaultDepositAmount), nil
}

func (m *ReBalancingStrategy) addYieldFees(baseFeeAmt, yield *big.Int) *big.Int {
	if yield.Cmp(big.NewInt(0)) > 0 {
		yieldFees, _ := new(big.Float).Mul(
			new(big.Float).SetInt(yield),
			new(big.Float).SetFloat64(m.config.YieldFees),
		).Int(nil)
		return new(big.Int).Add(baseFeeAmt, yieldFees)
	}
	return baseFeeAmt
}
