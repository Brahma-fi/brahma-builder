package vault

import (
	"fmt"

	vault "github.com/hashicorp/vault/api"
)

const (
	_encryptMountPath  = "%s/encrypt/%s"
	_decryptMountPath  = "%s/decrypt/%s"
	_defaultMountPath  = "transit"
	_encryptRequestKey = "plaintext"
	_decryptRequestKey = "ciphertext"
)

type Options struct {
	mountPath string
}

type Option func(*Options)

func WithMountPath(mountPath string) Option {
	return func(o *Options) {
		o.mountPath = mountPath
	}
}

func (v *Vault) Encrypt(plainText string, transitKey string, opts ...Option) (*vault.Secret, error) {
	return performVaultLogicalWrite(plainText, transitKey, _encryptMountPath, _encryptRequestKey, v, opts...)
}

func (v *Vault) Decrypt(ciphertext string, transitKey string, opts ...Option) (*vault.Secret, error) {
	return performVaultLogicalWrite(ciphertext, transitKey, _decryptMountPath, _decryptRequestKey, v, opts...)
}

func performVaultLogicalWrite(
	text, transitKey, formattedBasePath, requestParameterKey string,
	v *Vault,
	opts ...Option,
) (*vault.Secret, error) {
	if !v.pathRegex.MatchString(transitKey) {
		return nil, fmt.Errorf("invalid transit key name: %s", transitKey)
	}

	var options = Options{mountPath: _defaultMountPath}
	for _, opt := range opts {
		opt(&options)
	}

	var path string
	if !v.pathRegex.MatchString(options.mountPath) {
		return nil, fmt.Errorf("invalid mount path: %s", options.mountPath)
	}

	path = fmt.Sprintf(formattedBasePath, options.mountPath, transitKey)
	var secret, err = v.vaultClient.Logical().Write(
		path,
		map[string]interface{}{
			requestParameterKey: text,
		},
	)

	if err != nil {
		return nil, fmt.Errorf("unable to perform operation: %w", err)
	}

	return secret, nil
}
