package vault

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/Brahma-fi/brahma-builder/pkg/log"

	vault "github.com/hashicorp/vault/api"
	auth "github.com/hashicorp/vault/api/auth/kubernetes"
)

type Vault struct {
	secret          *vault.KVSecret
	vaultClient     *vault.Client
	authSecret      *vault.Secret
	lifetimeWatcher *vault.Renewer
	pathRegex       *regexp.Regexp
}

const (
	// accepts only letters, numbers and hyphens to avoid path traversal and url encoding issues
	_regexPattern                   = "^[a-zA-Z0-9-]+$"
	_unableToAuthenticateLogMessage = "unable to authenticate to vault: %w"
)

// NewVault initializes the Vault client
// VAULT_ADDR must be set in the environment
//   - local: VAULT_ADDR=http://localhost:8200
//   - environment: VAULT_ADDR=https://vault.vault.svc.cluster.local:8200
func NewVault(
	ctx context.Context,
	roleName, mountPath, secretPAth string,
	options ...auth.LoginOption,
) (*Vault, error) {
	var err error

	out := &Vault{}

	vaultConfig := vault.DefaultConfig() // modify for more granular configuration
	vaultConfig.Address = getVaultAddress()

	err = out.setTLSConfig(vaultConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to set TLS config, %w", err)
	}

	out.vaultClient, err = vault.NewClient(vaultConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to initialize vault client: %w", err)
	}

	out.authSecret, err = out.performAuthentication(ctx, out.vaultClient, roleName, options...)
	if err != nil {
		return nil, fmt.Errorf(_unableToAuthenticateLogMessage, err)
	}

	out.secret, err = out.vaultClient.KVv2(mountPath).Get(ctx, secretPAth)
	if err != nil {
		return nil, fmt.Errorf("unable to read secret: %w", err)
	}

	regex, err := regexp.Compile(_regexPattern)
	if err != nil {
		return nil, fmt.Errorf("unable to compile regex: %w", err)
	}
	out.pathRegex = regex

	return out, nil
}

func (v *Vault) GetConfigValue(vaultKeyName string) any {
	return v.secret.Data[vaultKeyName]
}

func (v *Vault) StopTokenRenew() {
	if v.lifetimeWatcher != nil {
		v.lifetimeWatcher.Stop()
	}
}

func (v *Vault) RunLifetimeWatcher(logger log.Logger) error {
	var err error

	v.lifetimeWatcher, err = v.vaultClient.NewLifetimeWatcher(&vault.LifetimeWatcherInput{Secret: v.authSecret})
	if err != nil {
		return fmt.Errorf("unable to set up lifetime watcher: %w", err)
	}

	go v.monitorLifetimeWatcher(logger)
	go v.lifetimeWatcher.Start()

	return nil
}

func WithServiceAccountTokenPath(path string) auth.LoginOption {
	return auth.WithServiceAccountTokenPath(path)
}

func (v *Vault) monitorLifetimeWatcher(logger log.Logger) {
	for {
		select {
		case err := <-v.lifetimeWatcher.DoneCh():
			logger.Info("shutting down vault token renewal")
			if err != nil {
				logger.Error("error renewing token", log.Err(err))
			}

		case <-v.lifetimeWatcher.RenewCh():
			logger.Info("successfully renewed vault token")
		}
	}
}

func (v *Vault) performAuthentication(
	ctx context.Context,
	client *vault.Client,
	roleName string,
	options ...auth.LoginOption,
) (*vault.Secret, error) {
	var err error
	var authSecret *vault.Secret

	if os.Getenv("ENV") == "local" {
		authSecret, err = v.localAuth(client)
		if err != nil {
			return nil, fmt.Errorf(_unableToAuthenticateLogMessage, err)
		}
	} else {
		authSecret, err = v.k8sAuth(ctx, client, roleName, options...)
		if err != nil {
			return nil, fmt.Errorf(_unableToAuthenticateLogMessage, err)
		}
	}

	return authSecret, nil
}

func (v *Vault) userpassLogin(client *vault.Client) (string, error) {
	// to pass the password
	var options = map[string]interface{}{
		"password": "local-pwd",
	}

	// PUT call to get a token
	secret, err := client.Logical().Write("auth/userpass/login/local-user", options)
	if err != nil {
		return "", err
	}

	token := secret.Auth.ClientToken

	return token, nil
}

func (v *Vault) localAuth(client *vault.Client) (*vault.Secret, error) {
	var token, err = v.userpassLogin(client)

	client.SetToken(token)
	// Renew the token immediately to get a secret to pass to lifetime watcher
	authSecret, err := client.Auth().Token().RenewTokenAsSelf(client.Token(), 0)
	if err != nil {
		return nil, fmt.Errorf(_unableToAuthenticateLogMessage, err)
	}

	return authSecret, nil
}

// performAuthentication performs the authentication to Vault
// ENV: local or anything else will be considered as environment
//   - local: uses the token from the TOKEN env var
//   - environment: uses the Kubernetes auth method
func (v *Vault) k8sAuth(
	ctx context.Context,
	client *vault.Client,
	roleName string,
	options ...auth.LoginOption,
) (*vault.Secret, error) {
	var k8sAuth, err = auth.NewKubernetesAuth(roleName, options...)
	if err != nil {
		return nil, fmt.Errorf("unable to initialize Kubernetes auth method, %w", err)
	}

	secret, err := client.Auth().Login(ctx, k8sAuth)
	if err != nil {
		return nil, fmt.Errorf("unable to log in with Kubernetes auth, %w", err)
	}

	return secret, nil
}

func (v *Vault) setTLSConfig(vaultConfig *vault.Config) error {
	// regular path: /etc/ssl/vault/certs/ca.crt
	var caFilePath = os.Getenv("VAULT_CA_FILE_PATH")

	if caFilePath != "" && strings.TrimSpace(caFilePath) != "" {
		// Load the TLS configuration
		tlsConfig := &tls.Config{MinVersion: tls.VersionTLS12}

		caCert, err := os.ReadFile(caFilePath)
		if err != nil {
			return fmt.Errorf("unable to read CA cert, %w", err)
		}

		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)
		tlsConfig.RootCAs = caCertPool

		// Set the TLS configuration in the Vault client
		vaultConfig.HttpClient.Transport = &http.Transport{TLSClientConfig: tlsConfig}
	}

	return nil
}

func getVaultAddress() string {
	return os.Getenv("VAULT_ADDR")
}
