package integrations

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/Brahma-fi/brahma-builder/internal/entity"
	bundler "github.com/Brahma-fi/brahma-builder/pkg/utils/abis/bundler"
	metamorpho "github.com/Brahma-fi/brahma-builder/pkg/utils/abis/metamorpho"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shurcooL/graphql"
)

type VaultOrderBy string
type OrderDirection string

const (
	VaultOrderByNetApy VaultOrderBy   = "NetApy"
	OrderDirectionDesc OrderDirection = "Desc"
)

type MorphoClient struct {
	client *graphql.Client
	caller bind.ContractCaller
}

func NewMorphoClient(baseURL string, caller bind.ContractCaller) *MorphoClient {
	return &MorphoClient{client: graphql.NewClient(baseURL, nil), caller: caller}
}

func (c *MorphoClient) Vaults(
	ctx context.Context,
	assetAddress common.Address,
	chainID int64,
) ([]entity.VaultInfo, error) {
	var query entity.VaultQuery
	variables := map[string]interface{}{
		"asset":          []graphql.String{graphql.String(assetAddress.Hex())},
		"chainID":        []graphql.Int{graphql.Int(chainID)},
		"orderBy":        VaultOrderByNetApy,
		"orderDirection": OrderDirectionDesc,
		"first":          graphql.Int(15),
	}

	err := c.client.Query(ctx, &query, variables)
	if err != nil {
		return nil, err
	}

	return query.ToVaultInfo(), nil
}

func (c *MorphoClient) User(ctx context.Context, address common.Address) ([]entity.UserInfo, error) {
	var query entity.UserQuery
	variables := map[string]interface{}{
		"address": []graphql.String{graphql.String(address.Hex())},
	}

	err := c.client.Query(ctx, &query, variables)
	if err != nil {
		return nil, err
	}

	if len(query.Users.Items) == 0 {
		return make([]entity.UserInfo, 0), nil
	}

	return query.ToUserInfo(), nil
}

func (c *MorphoClient) Shares(
	ctx context.Context,
	vaultAddr common.Address,
	depositor common.Address,
) (*big.Int, error) {
	vault, err := metamorpho.NewMorphoCaller(vaultAddr, c.caller)
	if err != nil {
		return nil, err
	}

	return vault.BalanceOf(&bind.CallOpts{Context: ctx}, depositor)
}

func (c *MorphoClient) PreviewRedeem(
	ctx context.Context,
	vaultAddr common.Address,
	depositor common.Address,
) (*big.Int, error) {
	shares, err := c.Shares(ctx, vaultAddr, depositor)
	if err != nil {
		return nil, err
	}

	vault, err := metamorpho.NewMorphoCaller(vaultAddr, c.caller)
	if err != nil {
		return nil, err
	}

	return vault.PreviewRedeem(&bind.CallOpts{Context: ctx}, shares)
}

func (c *MorphoClient) PreviewDeposit(
	ctx context.Context,
	vaultAddr common.Address,
	amt *big.Int,
) (*big.Int, error) {
	vault, err := metamorpho.NewMorphoCaller(vaultAddr, c.caller)
	if err != nil {
		return nil, err
	}

	return vault.PreviewDeposit(&bind.CallOpts{Context: ctx}, amt)
}

func (c *MorphoClient) RedeemMax(
	ctx context.Context,
	vaultAddr common.Address,
	depositor common.Address,
) ([]byte, error) {
	shares, err := c.Shares(ctx, vaultAddr, depositor)
	if err != nil {
		return nil, err
	}

	vaultAbi, err := abi.JSON(strings.NewReader(metamorpho.MorphoMetaData.ABI))
	if err != nil {
		return nil, err
	}

	return vaultAbi.Pack("redeem", shares, depositor, depositor)
}

func (c *MorphoClient) Deposit(
	depositor common.Address,
	amt *big.Int,
) ([]byte, error) {
	vaultAbi, err := abi.JSON(strings.NewReader(metamorpho.MorphoMetaData.ABI))
	if err != nil {
		return nil, err
	}

	return vaultAbi.Pack("deposit", amt, depositor)
}

func (c *MorphoClient) Bundle(
	calls []entity.BundlerCall,
) ([]byte, error) {
	bundlerAbi, err := abi.JSON(strings.NewReader(bundler.BundlerMetaData.ABI))
	if err != nil {
		return nil, err
	}

	multicall := make([][]byte, 0)
	for _, call := range calls {
		switch call.Type {
		case entity.BundlerCallTransferFrom:
			packed, err := bundlerAbi.Pack("erc20TransferFrom", call.Params...)
			if err != nil {
				return nil, err
			}

			multicall = append(multicall, packed)
		case entity.BundlerCallDeposit:
			packed, err := bundlerAbi.Pack("erc4626Deposit", call.Params...)
			if err != nil {
				return nil, err
			}
			multicall = append(multicall, packed)
		case entity.BundlerCallRedeem:
			packed, err := bundlerAbi.Pack("erc4626Redeem", call.Params...)
			if err != nil {
				return nil, err
			}

			multicall = append(multicall, packed)
		default:
			return nil, fmt.Errorf("unsupported call %d", call.Type)
		}
	}

	return bundlerAbi.Pack("multicall", multicall)
}
