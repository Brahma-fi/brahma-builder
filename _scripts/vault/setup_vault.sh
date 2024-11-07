export VAULT_ADDR='http://0.0.0.0:8200/'
cd $1
vault secrets enable -path=apps kv-v2
vault auth enable userpass
vault write auth/userpass/users/local-user password=local-pwd policies="default"
vault policy write default policy.hcl
vault kv put apps/brahma-builder/config @config.json