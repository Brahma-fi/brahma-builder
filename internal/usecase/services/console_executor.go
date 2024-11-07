package services

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/Brahma-fi/brahma-builder/internal/entity"
	"github.com/Brahma-fi/brahma-builder/pkg/rpc"
	utils "github.com/Brahma-fi/brahma-builder/pkg/utils/abis/executorplugin"
	"github.com/Brahma-fi/brahma-builder/pkg/utils/executor"
	"github.com/Brahma-fi/go-safe/wallet"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type ConsoleExecutor struct {
	executorAddress common.Address
	signerAddress   common.Address
	chainID         int64
	signer          keyManager
	client          console
	metadata        *entity.ExecutorMetadata
	executorPlugin  *utils.ExecutorpluginCaller
	pluginAddress   common.Address
}

func NewConsoleExecutor(
	ctx context.Context,
	executorAddress common.Address,
	chainID int64,
	rpc *rpc.RPC,
	signer keyManager,
	client console,
	signerAddress common.Address,
	executoPluginAddress common.Address,
) (*ConsoleExecutor, error) {
	metadata, err := client.ExecutorByAddressAndChainID(ctx, executorAddress, uint64(chainID))
	if err != nil {
		return nil, err
	}

	executorCaller, err := rpc.RetryableClient(int64(chainID))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch client: %w", err)
	}

	executorPlugin, err := utils.NewExecutorpluginCaller(executoPluginAddress, executorCaller)
	if err != nil {
		return nil, err
	}

	return &ConsoleExecutor{
		chainID:         chainID,
		metadata:        metadata,
		client:          client,
		signer:          signer,
		executorAddress: executorAddress,
		executorPlugin:  executorPlugin,
		signerAddress:   signerAddress,
		pluginAddress:   executoPluginAddress,
	}, nil
}

func (e *ConsoleExecutor) Execute(ctx context.Context, req *entity.SignAndExecuteRequest) (string, error) {
	val, _ := new(big.Int).SetString(req.Value, 10)
	callData, err := hexutil.Decode(req.Data)
	if err != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}

	nonce, err := e.executorPlugin.ExecutorNonce(&bind.CallOpts{
		Context: ctx,
	}, common.HexToAddress(req.Subaccount), e.executorAddress)
	if err != nil {
		return "", err
	}

	executableDigest, err := executor.GenerateExecutableDigest(executor.GenerateExecutableTypedDataParams{
		ChainID:       req.ChainID,
		PluginAddress: e.pluginAddress,
		To:            common.HexToAddress(req.To),
		Value:         val,
		Data:          callData,
		Operation:     req.Operation,
		Account:       common.HexToAddress(req.Subaccount),
		Nonce:         nonce,
		Executor:      e.executorAddress,
	})
	if err != nil {
		return "", err
	}

	safeMessageDigest, err := wallet.GetSafeMessageDigest(executableDigest, int64(req.ChainID), e.executorAddress)
	if err != nil {
		return "", err
	}
	signer := e.executorAddress
	if e.signerAddress != common.HexToAddress("") {
		signer = e.signerAddress
	}

	sig, err := e.signer.Sign(ctx, safeMessageDigest.Hex(), signer)
	if err != nil {
		return "", err
	}

	resp, err := e.client.Execute(ctx, &entity.ExecuteTaskReq{
		ChainID: int64(req.ChainID),
		Task: entity.Task{
			Subaccount:        req.Subaccount,
			Executor:          e.executorAddress.Hex(),
			ExecutorSignature: hexutil.Encode(sig),
			Executable: entity.Executable{
				CallType: req.Operation,
				To:       req.To,
				Value:    req.Value,
				Data:     req.Data,
			},
		},
		Webhook: "",
	})
	if err != nil {
		return "", err
	}

	switch {
	case resp.Error != "":
		return "", errors.New(resp.Error)
	case resp.Data.Error != "":
		return "", errors.New(resp.Data.Error)
	case resp.Data.Errors != "":
		return "", errors.New(resp.Data.Errors)
	case resp.Data.Data.TaskId != "":
		return resp.Data.Data.TaskId, nil
	default:
		return "", errors.New("failed to execute task")
	}
}

func (e *ConsoleExecutor) Subscriptions(ctx context.Context) ([]entity.ClientSubscription, error) {
	return e.client.ActiveSubscriptions(ctx, e.metadata.Id)
}
