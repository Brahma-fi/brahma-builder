package keymanager

import (
	"context"
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type KeyManager struct {
	address    common.Address
	privateKey *ecdsa.PrivateKey
}

func (m *KeyManager) Sign(_ context.Context, hash string, _ common.Address) ([]byte, error) {
	signature, err := crypto.Sign(common.HexToHash(hash).Bytes(), m.privateKey)
	if err != nil {
		return nil, err
	}

	if signature[crypto.RecoveryIDOffset] == 0 || signature[crypto.RecoveryIDOffset] == 1 {
		signature[crypto.RecoveryIDOffset] += 27 // Transform yellow paper V from 27/28 to 0/1
	}

	return signature, nil
}

func NewKeyManager(privKey string) (*KeyManager, error) {
	ecdsaKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return nil, err
	}

	return &KeyManager{
		address:    crypto.PubkeyToAddress(ecdsaKey.PublicKey),
		privateKey: ecdsaKey,
	}, nil
}
