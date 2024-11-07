package vault

import (
	"fmt"

	vault "github.com/hashicorp/vault/api"
)

func (v *Vault) CreateOrphanToken(tokenRequest *vault.TokenCreateRequest) (*vault.Secret, error) {
	var orphanToken, err = v.vaultClient.Auth().Token().CreateOrphan(tokenRequest)

	if err != nil {
		return nil, fmt.Errorf("unable to create orphan token: %w", err)
	}

	return orphanToken, nil
}
