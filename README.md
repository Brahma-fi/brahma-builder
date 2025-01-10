# Brahma Builder

Scalable, high-performance workflow scheduling engine built on Temporal for executing custom strategies with the Brahma Builder Kit API. Seamlessly manages workflow orchestration, fault tolerance, and Brahma ecosystem integration.

- **[Architecture reference](./docs/architecture.md)**
- **[Additional context](https://www.notion.so/brahmafi/Bringing-AI-Agents-On-Chain-Automate-Execute-and-Scale-with-Brahma-175a53ecb04c80d9b9d6cf16cd1dd98a#175a53ecb04c8078b143f61bd9681e2b)**

## Features

- Zero-config workflow scheduling
- Automatic retry handling
- Native Brahma Builder Kit integration
- Temporal engine under the hood

## Running locally

1. Setup vault

```
make setup-local-vault
make setup-local-plugin
```

2. Setup env

```
export VAULT_ADDR=127.0.0.1:8200
export ENV=local
```

3. Run scheduler & workers

```
go run cmd/main.go scheduler|base-worker|morpho-worker
```

## Example

Morpho Yield Optimizer is a strategy that is built using Brahma builder. It maximises user’s Morpho positions by taking decisions on which vaults to choose based on APY; liquidity and TVL, on every rebalance.

For reference implementation, see [Morpho Optimizer Strategy](https://github.com/Brahma-fi/brahma-builder/blob/main/internal/usecase/workflows/activities/morpho/activity.go).
