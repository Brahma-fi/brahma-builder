package vault

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

const (
	_signPath            = "ethereum/key-managers/%s/sign"
	_signTransactionPath = "ethereum/key-managers/%s/txn/sign"
	_pathListSigners     = "ethereum/key-managers/%s"
)

type KeyManager struct {
	vault       *Vault
	serviceName string
}

func NewKeyManager(vault *Vault, serviceName string) *KeyManager {
	return &KeyManager{vault: vault, serviceName: serviceName}
}

func (m *KeyManager) Sign(ctx context.Context, hash string, signer common.Address) ([]byte, error) {
	secret, err := m.vault.vaultClient.Logical().WriteWithContext(
		ctx,
		fmt.Sprintf(_signPath, m.serviceName),
		map[string]any{
			"hash":    hash,
			"address": signer.Hex(),
		},
	)
	if err != nil {
		return nil, err
	}

	sig, ok := secret.Data["signature"].(string)
	if !ok {
		return nil, errors.New("failed to parse signed signature")
	}

	signature := common.Hex2Bytes(sig)
	if signature[crypto.RecoveryIDOffset] == 0 || signature[crypto.RecoveryIDOffset] == 1 {
		signature[crypto.RecoveryIDOffset] += 27 // Transform yellow paper V from 27/28 to 0/1
	}

	return signature, nil
}

func (m *KeyManager) SignTransaction(
	ctx context.Context,
	signer common.Address,
	transaction *types.Transaction,
	chainID *big.Int,
) (*types.Transaction, error) {
	signature, err := m.SignTxn(ctx, signer, transaction, chainID)
	if err != nil {
		return nil, err
	}

	tx := &types.Transaction{}

	err = tx.DecodeRLP(rlp.NewStream(bytes.NewReader(signature), 0))
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (m *KeyManager) SignTxn(
	ctx context.Context,
	signer common.Address,
	transaction *types.Transaction,
	chainID *big.Int,
) ([]byte, error) {
	// if encoded nonce is 0x, replace it with 0x0
	nonce := hexutil.Encode(new(big.Int).SetUint64(transaction.Nonce()).Bytes())
	if nonce == "0x" {
		nonce = "0x0"
	}

	response, err := m.vault.vaultClient.Logical().WriteWithContext(
		ctx,
		fmt.Sprintf(_signTransactionPath, m.serviceName),
		map[string]any{
			"input":   hexutil.Encode(transaction.Data()),
			"address": signer.Hex(),
			"to":      transaction.To().Hex(),
			"gas":     transaction.Gas(),
			//convert nonce to hex
			"nonce":     nonce,
			"gasFeeCap": transaction.GasFeeCap().String(),
			"gasTipCap": transaction.GasTipCap().String(),
			"chainId":   chainID.String(),
		},
	)
	if err != nil {
		return nil, err
	}

	signedTx := response.Data["signedTx"].(string)
	signatureBytes, err := hexutil.Decode(signedTx)
	if err != nil {
		return nil, err
	}

	return signatureBytes, nil
}
