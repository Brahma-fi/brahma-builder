package executor

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

type GenerateExecutableTypedDataParams struct {
	PluginAddress common.Address
	ChainID       uint64
	To            common.Address
	Value         *big.Int
	Data          []byte
	Operation     uint8
	Account       common.Address
	Nonce         *big.Int
	Executor      common.Address
}

func GenerateExecutableDigest(params GenerateExecutableTypedDataParams) (common.Hash, error) {
	executionParamsTypedData := &apitypes.TypedData{
		Types: apitypes.Types{
			"ExecutionParams": {
				{Name: "operation", Type: "uint8"},
				{Name: "to", Type: "address"},
				{Name: "account", Type: "address"},
				{Name: "executor", Type: "address"},
				{Name: "gasToken", Type: "address"},
				{Name: "refundReceiver", Type: "address"},
				{Name: "value", Type: "uint256"},
				{Name: "nonce", Type: "uint256"},
				{Name: "safeTxGas", Type: "uint256"},
				{Name: "baseGas", Type: "uint256"},
				{Name: "gasPrice", Type: "uint256"},
				{Name: "data", Type: "bytes"},
			},
			"EIP712Domain": {
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},

				{Name: "verifyingContract", Type: "address"},
			},
		},
		PrimaryType: "ExecutionParams",
		Domain: apitypes.TypedDataDomain{
			Name:    "ExecutorPlugin",
			Version: "1.0",
			ChainId: math.NewHexOrDecimal256(int64(params.ChainID)),

			VerifyingContract: params.PluginAddress.String(),
		},
		Message: map[string]interface{}{
			"operation":      fmt.Sprintf("%d", params.Operation),
			"to":             params.To.Hex(),
			"account":        params.Account.Hex(),
			"executor":       params.Executor.Hex(),
			"value":          params.Value,
			"nonce":          params.Nonce,
			"data":           params.Data,
			"gasToken":       common.HexToAddress("").Hex(),
			"refundReceiver": common.HexToAddress("").Hex(),
			"safeTxGas":      new(big.Int).SetInt64(0),
			"baseGas":        new(big.Int).SetInt64(0),
			"gasPrice":       new(big.Int).SetInt64(0),
		},
	}

	domainSeparator, err := executionParamsTypedData.HashStruct("EIP712Domain", executionParamsTypedData.Domain.Map())
	if err != nil {
		return common.Hash{}, fmt.Errorf("eip712domain hash struct: %w", err)
	}

	typedDataHash, err := executionParamsTypedData.HashStruct(
		executionParamsTypedData.PrimaryType,
		executionParamsTypedData.Message,
	)
	if err != nil {
		return common.Hash{}, fmt.Errorf("primary type hash struct: %w", err)
	}

	return crypto.Keccak256Hash([]byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))), nil
}
