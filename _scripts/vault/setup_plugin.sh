export VAULT_ADDR='http://0.0.0.0:8200/'
cd $1
sha256="$(shasum -a 256 $2 | cut -d' ' -f1)"
ls -lah
echo "sha256: $sha256"
vault plugin register -sha256="$sha256" secret $2
vault secrets enable -path=ethereum -description="Ethereum signer" -plugin-name=$2 plugin
vault write ethereum/key-managers serviceName="brahma-builder" privateKey=$3
vault list ethereum/key-managers
vault read ethereum/key-managers/brahma-builder