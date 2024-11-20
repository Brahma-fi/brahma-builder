package morpho

import (
	"context"
	"math/big"

	"github.com/Brahma-fi/brahma-builder/internal/entity"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
)

type morphoClient interface {
	Vaults(
		ctx context.Context,
		assetAddress common.Address,
		chainID int64,
	) ([]entity.VaultInfo, error)
	User(ctx context.Context, address common.Address) ([]entity.UserInfo, error)
	Deposit(
		depositor common.Address,
		amt *big.Int,
	) ([]byte, error)
	RedeemMax(
		ctx context.Context,
		vaultAddr common.Address,
		depositor common.Address,
	) ([]byte, error)
	PreviewRedeem(
		ctx context.Context,
		vaultAddr common.Address,
		depositor common.Address,
	) (*big.Int, error)
	Shares(
		ctx context.Context,
		vaultAddr common.Address,
		depositor common.Address,
	) (*big.Int, error)
	Bundle(
		calls []entity.BundlerCall,
	) ([]byte, error)
	PreviewDeposit(
		ctx context.Context,
		vaultAddr common.Address,
		amt *big.Int,
	) (*big.Int, error)
}

type consoleExecutor interface {
	Execute(ctx context.Context, req *entity.SignAndExecuteRequest) (string, error)
}

type executionsLogRepo interface {
	LatestBySubID(
		ctx context.Context,
		subID uuid.UUID,
	) (*entity.Log, error)
}

type pricingOracle interface {
	ConvertUSDToToken(
		ctx context.Context,
		chainID int64,
		amtUSD float64,
		tokenAddress common.Address,
	) (*big.Int, error)
}
