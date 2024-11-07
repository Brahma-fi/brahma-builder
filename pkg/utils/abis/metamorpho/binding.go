// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package utils

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// MarketAllocation is an auto generated low-level Go binding around an user-defined struct.
type MarketAllocation struct {
	MarketParams MarketParams
	Assets       *big.Int
}

// MarketParams is an auto generated low-level Go binding around an user-defined struct.
type MarketParams struct {
	LoanToken       common.Address
	CollateralToken common.Address
	Oracle          common.Address
	Irm             common.Address
	Lltv            *big.Int
}

// MorphoMetaData contains all meta data concerning the Morpho contract.
var MorphoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"morpho\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialTimelock\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AboveMaxTimelock\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllCapsReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyPending\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BelowMinTimelock\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"DuplicateMarket\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ERC2612ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC2612InvalidSigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxDeposit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxMint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxRedeem\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxWithdraw\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"InconsistentAsset\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InconsistentReallocation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"InvalidMarketRemovalNonZeroCap\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"InvalidMarketRemovalNonZeroSupply\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"InvalidMarketRemovalTimelockNotElapsed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MarketNotCreated\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"MarketNotEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MathOverflowedMulDiv\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxFeeExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxQueueLengthExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPendingValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonZeroCap\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllocatorRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotCuratorNorGuardianRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotCuratorRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughLiquidity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotGuardianRole\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"PendingCap\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingRemoval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"SupplyCapExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimelockNotElapsed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"UnauthorizedMarket\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroFeeRecipient\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalAssets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeShares\",\"type\":\"uint256\"}],\"name\":\"AccrueInterest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"suppliedAssets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"suppliedShares\",\"type\":\"uint256\"}],\"name\":\"ReallocateSupply\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawnAssets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawnShares\",\"type\":\"uint256\"}],\"name\":\"ReallocateWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"RevokePendingCap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"RevokePendingGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"RevokePendingMarketRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"RevokePendingTimelock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"name\":\"SetCap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newCurator\",\"type\":\"address\"}],\"name\":\"SetCurator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"SetFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newFeeRecipient\",\"type\":\"address\"}],\"name\":\"SetFeeRecipient\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"SetGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"allocator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isAllocator\",\"type\":\"bool\"}],\"name\":\"SetIsAllocator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newSkimRecipient\",\"type\":\"address\"}],\"name\":\"SetSkimRecipient\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"Id[]\",\"name\":\"newSupplyQueue\",\"type\":\"bytes32[]\"}],\"name\":\"SetSupplyQueue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTimelock\",\"type\":\"uint256\"}],\"name\":\"SetTimelock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"Id[]\",\"name\":\"newWithdrawQueue\",\"type\":\"bytes32[]\"}],\"name\":\"SetWithdrawQueue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Skim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"name\":\"SubmitCap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGuardian\",\"type\":\"address\"}],\"name\":\"SubmitGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"SubmitMarketRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTimelock\",\"type\":\"uint256\"}],\"name\":\"SubmitTimelock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"updatedTotalAssets\",\"type\":\"uint256\"}],\"name\":\"UpdateLastTotalAssets\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DECIMALS_OFFSET\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MORPHO\",\"outputs\":[{\"internalType\":\"contractIMorpho\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"loanToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"irm\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lltv\",\"type\":\"uint256\"}],\"internalType\":\"structMarketParams\",\"name\":\"marketParams\",\"type\":\"tuple\"}],\"name\":\"acceptCap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptTimelock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"config\",\"outputs\":[{\"internalType\":\"uint184\",\"name\":\"cap\",\"type\":\"uint184\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"removableAt\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"curator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"guardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isAllocator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTotalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"pendingCap\",\"outputs\":[{\"internalType\":\"uint192\",\"name\":\"value\",\"type\":\"uint192\"},{\"internalType\":\"uint64\",\"name\":\"validAt\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"validAt\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingTimelock\",\"outputs\":[{\"internalType\":\"uint192\",\"name\":\"value\",\"type\":\"uint192\"},{\"internalType\":\"uint64\",\"name\":\"validAt\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"loanToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"irm\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lltv\",\"type\":\"uint256\"}],\"internalType\":\"structMarketParams\",\"name\":\"marketParams\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"internalType\":\"structMarketAllocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"name\":\"reallocate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"revokePendingCap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revokePendingGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"revokePendingMarketRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revokePendingTimelock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newCurator\",\"type\":\"address\"}],\"name\":\"setCurator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newFeeRecipient\",\"type\":\"address\"}],\"name\":\"setFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAllocator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"newIsAllocator\",\"type\":\"bool\"}],\"name\":\"setIsAllocator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newSkimRecipient\",\"type\":\"address\"}],\"name\":\"setSkimRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Id[]\",\"name\":\"newSupplyQueue\",\"type\":\"bytes32[]\"}],\"name\":\"setSupplyQueue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"skim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"skimRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"loanToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"irm\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lltv\",\"type\":\"uint256\"}],\"internalType\":\"structMarketParams\",\"name\":\"marketParams\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"newSupplyCap\",\"type\":\"uint256\"}],\"name\":\"submitCap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newGuardian\",\"type\":\"address\"}],\"name\":\"submitGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"loanToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"irm\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lltv\",\"type\":\"uint256\"}],\"internalType\":\"structMarketParams\",\"name\":\"marketParams\",\"type\":\"tuple\"}],\"name\":\"submitMarketRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTimelock\",\"type\":\"uint256\"}],\"name\":\"submitTimelock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"supplyQueue\",\"outputs\":[{\"internalType\":\"Id\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supplyQueueLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timelock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"indexes\",\"type\":\"uint256[]\"}],\"name\":\"updateWithdrawQueue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdrawQueue\",\"outputs\":[{\"internalType\":\"Id\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawQueueLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MorphoABI is the input ABI used to generate the binding from.
// Deprecated: Use MorphoMetaData.ABI instead.
var MorphoABI = MorphoMetaData.ABI

// Morpho is an auto generated Go binding around an Ethereum contract.
type Morpho struct {
	MorphoCaller     // Read-only binding to the contract
	MorphoTransactor // Write-only binding to the contract
	MorphoFilterer   // Log filterer for contract events
}

// MorphoCaller is an auto generated read-only Go binding around an Ethereum contract.
type MorphoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MorphoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MorphoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MorphoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MorphoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MorphoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MorphoSession struct {
	Contract     *Morpho           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MorphoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MorphoCallerSession struct {
	Contract *MorphoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MorphoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MorphoTransactorSession struct {
	Contract     *MorphoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MorphoRaw is an auto generated low-level Go binding around an Ethereum contract.
type MorphoRaw struct {
	Contract *Morpho // Generic contract binding to access the raw methods on
}

// MorphoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MorphoCallerRaw struct {
	Contract *MorphoCaller // Generic read-only contract binding to access the raw methods on
}

// MorphoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MorphoTransactorRaw struct {
	Contract *MorphoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMorpho creates a new instance of Morpho, bound to a specific deployed contract.
func NewMorpho(address common.Address, backend bind.ContractBackend) (*Morpho, error) {
	contract, err := bindMorpho(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Morpho{MorphoCaller: MorphoCaller{contract: contract}, MorphoTransactor: MorphoTransactor{contract: contract}, MorphoFilterer: MorphoFilterer{contract: contract}}, nil
}

// NewMorphoCaller creates a new read-only instance of Morpho, bound to a specific deployed contract.
func NewMorphoCaller(address common.Address, caller bind.ContractCaller) (*MorphoCaller, error) {
	contract, err := bindMorpho(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MorphoCaller{contract: contract}, nil
}

// NewMorphoTransactor creates a new write-only instance of Morpho, bound to a specific deployed contract.
func NewMorphoTransactor(address common.Address, transactor bind.ContractTransactor) (*MorphoTransactor, error) {
	contract, err := bindMorpho(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MorphoTransactor{contract: contract}, nil
}

// NewMorphoFilterer creates a new log filterer instance of Morpho, bound to a specific deployed contract.
func NewMorphoFilterer(address common.Address, filterer bind.ContractFilterer) (*MorphoFilterer, error) {
	contract, err := bindMorpho(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MorphoFilterer{contract: contract}, nil
}

// bindMorpho binds a generic wrapper to an already deployed contract.
func bindMorpho(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MorphoMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Morpho *MorphoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Morpho.Contract.MorphoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Morpho *MorphoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Morpho.Contract.MorphoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Morpho *MorphoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Morpho.Contract.MorphoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Morpho *MorphoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Morpho.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Morpho *MorphoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Morpho.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Morpho *MorphoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Morpho.Contract.contract.Transact(opts, method, params...)
}

// DECIMALSOFFSET is a free data retrieval call binding the contract method 0xaea70acc.
//
// Solidity: function DECIMALS_OFFSET() view returns(uint8)
func (_Morpho *MorphoCaller) DECIMALSOFFSET(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Morpho.contract.Call(opts, &out, "DECIMALS_OFFSET")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DECIMALSOFFSET is a free data retrieval call binding the contract method 0xaea70acc.
//
// Solidity: function DECIMALS_OFFSET() view returns(uint8)
func (_Morpho *MorphoSession) DECIMALSOFFSET() (uint8, error) {
	return _Morpho.Contract.DECIMALSOFFSET(&_Morpho.CallOpts)
}

// DECIMALSOFFSET is a free data retrieval call binding the contract method 0xaea70acc.
//
// Solidity: function DECIMALS_OFFSET() view returns(uint8)
func (_Morpho *MorphoCallerSession) DECIMALSOFFSET() (uint8, error) {
	return _Morpho.Contract.DECIMALSOFFSET(&_Morpho.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Morpho *MorphoCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Morpho.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Morpho *MorphoSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Morpho.Contract.DOMAINSEPARATOR(&_Morpho.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Morpho *MorphoCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Morpho.Contract.DOMAINSEPARATOR(&_Morpho.CallOpts)
}

// MORPHO is a free data retrieval call binding the contract method 0x3acb5624.
//
// Solidity: function MORPHO() view returns(address)
func (_Morpho *MorphoCaller) MORPHO(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Morpho.contract.Call(opts, &out, "MORPHO")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MORPHO is a free data retrieval call binding the contract method 0x3acb5624.
//
// Solidity: function MORPHO() view returns(address)
func (_Morpho *MorphoSession) MORPHO() (common.Address, error) {
	return _Morpho.Contract.MORPHO(&_Morpho.CallOpts)
}

// MORPHO is a free data retrieval call binding the contract method 0x3acb5624.
//
// Solidity: function MORPHO() view returns(address)
func (_Morpho *MorphoCallerSession) MORPHO() (common.Address, error) {
	return _Morpho.Contract.MORPHO(&_Morpho.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Morpho *MorphoCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Morpho.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Morpho *MorphoSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Morpho.Contract.Allowance(&_Morpho.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Morpho *MorphoCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Morpho.Contract.Allowance(&_Morpho.CallOpts, owner, spender)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_Morpho *MorphoCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Morpho.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_Morpho *MorphoSession) Asset() (common.Address, error) {
	return _Morpho.Contract.Asset(&_Morpho.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_Morpho *MorphoCallerSession) Asset() (common.Address, error) {
	return _Morpho.Contract.Asset(&_Morpho.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Morpho *MorphoCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Morpho.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Morpho *MorphoSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Morpho.Contract.BalanceOf(&_Morpho.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Morpho *MorphoCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Morpho.Contract.BalanceOf(&_Morpho.CallOpts, account)
}

// Config is a free data retrieval call binding the contract method 0xcc718f76.
//
// Solidity: function config(bytes32 ) view returns(uint184 cap, bool enabled, uint64 removableAt)
func (_Morpho *MorphoCaller) Config(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Cap         *big.Int
	Enabled     bool
	RemovableAt uint64
}, error) {
	var out []interface{}
	err := _Morpho.contract.Call(opts, &out, "config", arg0)

	outstruct := new(struct {
		Cap         *big.Int
		Enabled     bool
		RemovableAt uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Cap = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Enabled = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.RemovableAt = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// Config is a free data retrieval call binding the contract method 0xcc718f76.
//
// Solidity: function config(bytes32 ) view returns(uint184 cap, bool enabled, uint64 removableAt)
func (_Morpho *MorphoSession) Config(arg0 [32]byte) (struct {
	Cap         *big.Int
	Enabled     bool
	RemovableAt uint64
}, error) {
	return _Morpho.Contract.Config(&_Morpho.CallOpts, arg0)
}

// Config is a free data retrieval call binding the contract method 0xcc718f76.
//
// Solidity: function config(bytes32 ) view returns(uint184 cap, bool enabled, uint64 removableAt)
func (_Morpho *MorphoCallerSession) Config(arg0 [32]byte) (struct {
	Cap         *big.Int
	Enabled     bool
	RemovableAt uint64
}, error) {
	return _Morpho.Contract.Config(&_Morpho.CallOpts, arg0)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_Morpho *MorphoCaller) ConvertToAssets(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Morpho.contract.Call(opts, &out, "convertToAssets", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_Morpho *MorphoSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _Morpho.Contract.ConvertToAssets(&_Morpho.CallOpts, shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_Morpho *MorphoCallerSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _Morpho.Contract.ConvertToAssets(&_Morpho.CallOpts, shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_Morpho *MorphoCaller) ConvertToShares(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Morpho.contract.Call(opts, &out, "convertToShares", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_Morpho *MorphoSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _Morpho.Contract.ConvertToShares(&_Morpho.CallOpts, assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_Morpho *MorphoCallerSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _Morpho.Contract.ConvertToShares(&_Morpho.CallOpts, assets)
}

// Curator is a free data retrieval call binding the contract method 0xe66f53b7.
//
// Solidity: function curator() view returns(address)
func (_Morpho *MorphoCaller) Curator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Morpho.contract.Call(opts, &out, "curator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Curator is a free data retrieval call binding the contract method 0xe66f53b7.
//
// Solidity: function curator() view returns(address)
func (_Morpho *MorphoSession) Curator() (common.Address, error) {
	return _Morpho.Contract.Curator(&_Morpho.CallOpts)
}

// Curator is a free data retrieval call binding the contract method 0xe66f53b7.
//
// Solidity: function curator() view returns(address)
func (_Morpho *MorphoCallerSession) Curator() (common.Address, error) {
	return _Morpho.Contract.Curator(&_Morpho.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Morpho *MorphoCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Morpho.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Morpho *MorphoSession) Decimals() (uint8, error) {
	return _Morpho.Contract.Decimals(&_Morpho.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Morpho *MorphoCallerSession) Decimals() (uint8, error) {
	return _Morpho.Contract.Decimals(&_Morpho.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Morpho *MorphoCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _Morpho.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Morpho *MorphoSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Morpho.Contract.Eip712Domain(&_Morpho.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Morpho *MorphoCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Morpho.Contract.Eip712Domain(&_Morpho.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint96)
func (_Morpho *MorphoCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Morpho.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint96)
func (_Morpho *MorphoSession) Fee() (*big.Int, error) {
	return _Morpho.Contract.Fee(&_Morpho.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint96)
func (_Morpho *MorphoCallerSession) Fee() (*big.Int, error) {
	return _Morpho.Contract.Fee(&_Morpho.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Morpho *MorphoCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Morpho.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Morpho *MorphoSession) FeeRecipient() (common.Address, error) {
	return _Morpho.Contract.FeeRecipient(&_Morpho.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Morpho *MorphoCallerSession) FeeRecipient() (common.Address, error) {
	return _Morpho.Contract.FeeRecipient(&_Morpho.CallOpts)
}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_Morpho *MorphoCaller) Guardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Morpho.contract.Call(opts, &out, "guardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_Morpho *MorphoSession) Guardian() (common.Address, error) {
	return _Morpho.Contract.Guardian(&_Morpho.CallOpts)
}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_Morpho *MorphoCallerSession) Guardian() (common.Address, error) {
	return _Morpho.Contract.Guardian(&_Morpho.CallOpts)
}

// IsAllocator is a free data retrieval call binding the contract method 0x4dedf20e.
//
// Solidity: function isAllocator(address ) view returns(bool)
func (_Morpho *MorphoCaller) IsAllocator(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Morpho.contract.Call(opts, &out, "isAllocator", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllocator is a free data retrieval call binding the contract method 0x4dedf20e.
//
// Solidity: function isAllocator(address ) view returns(bool)
func (_Morpho *MorphoSession) IsAllocator(arg0 common.Address) (bool, error) {
	return _Morpho.Contract.IsAllocator(&_Morpho.CallOpts, arg0)
}

// IsAllocator is a free data retrieval call binding the contract method 0x4dedf20e.
//
// Solidity: function isAllocator(address ) view returns(bool)
func (_Morpho *MorphoCallerSession) IsAllocator(arg0 common.Address) (bool, error) {
	return _Morpho.Contract.IsAllocator(&_Morpho.CallOpts, arg0)
}

// LastTotalAssets is a free data retrieval call binding the contract method 0x568efc07.
//
// Solidity: function lastTotalAssets() view returns(uint256)
func (_Morpho *MorphoCaller) LastTotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Morpho.contract.Call(opts, &out, "lastTotalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTotalAssets is a free data retrieval call binding the contract method 0x568efc07.
//
// Solidity: function lastTotalAssets() view returns(uint256)
func (_Morpho *MorphoSession) LastTotalAssets() (*big.Int, error) {
	return _Morpho.Contract.LastTotalAssets(&_Morpho.CallOpts)
}

// LastTotalAssets is a free data retrieval call binding the contract method 0x568efc07.
//
// Solidity: function lastTotalAssets() view returns(uint256)
func (_Morpho *MorphoCallerSession) LastTotalAssets() (*big.Int, error) {
	return _Morpho.Contract.LastTotalAssets(&_Morpho.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_Morpho *MorphoCaller) MaxDeposit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Morpho.contract.Call(opts, &out, "maxDeposit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_Morpho *MorphoSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _Morpho.Contract.MaxDeposit(&_Morpho.CallOpts, arg0)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_Morpho *MorphoCallerSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _Morpho.Contract.MaxDeposit(&_Morpho.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_Morpho *MorphoCaller) MaxMint(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Morpho.contract.Call(opts, &out, "maxMint", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_Morpho *MorphoSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _Morpho.Contract.MaxMint(&_Morpho.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_Morpho *MorphoCallerSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _Morpho.Contract.MaxMint(&_Morpho.CallOpts, arg0)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_Morpho *MorphoCaller) MaxRedeem(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Morpho.contract.Call(opts, &out, "maxRedeem", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_Morpho *MorphoSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _Morpho.Contract.MaxRedeem(&_Morpho.CallOpts, owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_Morpho *MorphoCallerSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _Morpho.Contract.MaxRedeem(&_Morpho.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256 assets)
func (_Morpho *MorphoCaller) MaxWithdraw(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Morpho.contract.Call(opts, &out, "maxWithdraw", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256 assets)
func (_Morpho *MorphoSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _Morpho.Contract.MaxWithdraw(&_Morpho.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256 assets)
func (_Morpho *MorphoCallerSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _Morpho.Contract.MaxWithdraw(&_Morpho.CallOpts, owner)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Morpho *MorphoCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Morpho.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Morpho *MorphoSession) Name() (string, error) {
	return _Morpho.Contract.Name(&_Morpho.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Morpho *MorphoCallerSession) Name() (string, error) {
	return _Morpho.Contract.Name(&_Morpho.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Morpho *MorphoCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Morpho.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Morpho *MorphoSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Morpho.Contract.Nonces(&_Morpho.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Morpho *MorphoCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Morpho.Contract.Nonces(&_Morpho.CallOpts, owner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Morpho *MorphoCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Morpho.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Morpho *MorphoSession) Owner() (common.Address, error) {
	return _Morpho.Contract.Owner(&_Morpho.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Morpho *MorphoCallerSession) Owner() (common.Address, error) {
	return _Morpho.Contract.Owner(&_Morpho.CallOpts)
}

// PendingCap is a free data retrieval call binding the contract method 0xa31be5d6.
//
// Solidity: function pendingCap(bytes32 ) view returns(uint192 value, uint64 validAt)
func (_Morpho *MorphoCaller) PendingCap(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Value   *big.Int
	ValidAt uint64
}, error) {
	var out []interface{}
	err := _Morpho.contract.Call(opts, &out, "pendingCap", arg0)

	outstruct := new(struct {
		Value   *big.Int
		ValidAt uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Value = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ValidAt = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// PendingCap is a free data retrieval call binding the contract method 0xa31be5d6.
//
// Solidity: function pendingCap(bytes32 ) view returns(uint192 value, uint64 validAt)
func (_Morpho *MorphoSession) PendingCap(arg0 [32]byte) (struct {
	Value   *big.Int
	ValidAt uint64
}, error) {
	return _Morpho.Contract.PendingCap(&_Morpho.CallOpts, arg0)
}

// PendingCap is a free data retrieval call binding the contract method 0xa31be5d6.
//
// Solidity: function pendingCap(bytes32 ) view returns(uint192 value, uint64 validAt)
func (_Morpho *MorphoCallerSession) PendingCap(arg0 [32]byte) (struct {
	Value   *big.Int
	ValidAt uint64
}, error) {
	return _Morpho.Contract.PendingCap(&_Morpho.CallOpts, arg0)
}

// PendingGuardian is a free data retrieval call binding the contract method 0x762c31ba.
//
// Solidity: function pendingGuardian() view returns(address value, uint64 validAt)
func (_Morpho *MorphoCaller) PendingGuardian(opts *bind.CallOpts) (struct {
	Value   common.Address
	ValidAt uint64
}, error) {
	var out []interface{}
	err := _Morpho.contract.Call(opts, &out, "pendingGuardian")

	outstruct := new(struct {
		Value   common.Address
		ValidAt uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Value = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ValidAt = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// PendingGuardian is a free data retrieval call binding the contract method 0x762c31ba.
//
// Solidity: function pendingGuardian() view returns(address value, uint64 validAt)
func (_Morpho *MorphoSession) PendingGuardian() (struct {
	Value   common.Address
	ValidAt uint64
}, error) {
	return _Morpho.Contract.PendingGuardian(&_Morpho.CallOpts)
}

// PendingGuardian is a free data retrieval call binding the contract method 0x762c31ba.
//
// Solidity: function pendingGuardian() view returns(address value, uint64 validAt)
func (_Morpho *MorphoCallerSession) PendingGuardian() (struct {
	Value   common.Address
	ValidAt uint64
}, error) {
	return _Morpho.Contract.PendingGuardian(&_Morpho.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Morpho *MorphoCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Morpho.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Morpho *MorphoSession) PendingOwner() (common.Address, error) {
	return _Morpho.Contract.PendingOwner(&_Morpho.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Morpho *MorphoCallerSession) PendingOwner() (common.Address, error) {
	return _Morpho.Contract.PendingOwner(&_Morpho.CallOpts)
}

// PendingTimelock is a free data retrieval call binding the contract method 0x7cc4d9a1.
//
// Solidity: function pendingTimelock() view returns(uint192 value, uint64 validAt)
func (_Morpho *MorphoCaller) PendingTimelock(opts *bind.CallOpts) (struct {
	Value   *big.Int
	ValidAt uint64
}, error) {
	var out []interface{}
	err := _Morpho.contract.Call(opts, &out, "pendingTimelock")

	outstruct := new(struct {
		Value   *big.Int
		ValidAt uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Value = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ValidAt = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// PendingTimelock is a free data retrieval call binding the contract method 0x7cc4d9a1.
//
// Solidity: function pendingTimelock() view returns(uint192 value, uint64 validAt)
func (_Morpho *MorphoSession) PendingTimelock() (struct {
	Value   *big.Int
	ValidAt uint64
}, error) {
	return _Morpho.Contract.PendingTimelock(&_Morpho.CallOpts)
}

// PendingTimelock is a free data retrieval call binding the contract method 0x7cc4d9a1.
//
// Solidity: function pendingTimelock() view returns(uint192 value, uint64 validAt)
func (_Morpho *MorphoCallerSession) PendingTimelock() (struct {
	Value   *big.Int
	ValidAt uint64
}, error) {
	return _Morpho.Contract.PendingTimelock(&_Morpho.CallOpts)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_Morpho *MorphoCaller) PreviewDeposit(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Morpho.contract.Call(opts, &out, "previewDeposit", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_Morpho *MorphoSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _Morpho.Contract.PreviewDeposit(&_Morpho.CallOpts, assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_Morpho *MorphoCallerSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _Morpho.Contract.PreviewDeposit(&_Morpho.CallOpts, assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_Morpho *MorphoCaller) PreviewMint(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Morpho.contract.Call(opts, &out, "previewMint", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_Morpho *MorphoSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _Morpho.Contract.PreviewMint(&_Morpho.CallOpts, shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_Morpho *MorphoCallerSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _Morpho.Contract.PreviewMint(&_Morpho.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_Morpho *MorphoCaller) PreviewRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Morpho.contract.Call(opts, &out, "previewRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_Morpho *MorphoSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _Morpho.Contract.PreviewRedeem(&_Morpho.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_Morpho *MorphoCallerSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _Morpho.Contract.PreviewRedeem(&_Morpho.CallOpts, shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_Morpho *MorphoCaller) PreviewWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Morpho.contract.Call(opts, &out, "previewWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_Morpho *MorphoSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _Morpho.Contract.PreviewWithdraw(&_Morpho.CallOpts, assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_Morpho *MorphoCallerSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _Morpho.Contract.PreviewWithdraw(&_Morpho.CallOpts, assets)
}

// SkimRecipient is a free data retrieval call binding the contract method 0x388af5b5.
//
// Solidity: function skimRecipient() view returns(address)
func (_Morpho *MorphoCaller) SkimRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Morpho.contract.Call(opts, &out, "skimRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SkimRecipient is a free data retrieval call binding the contract method 0x388af5b5.
//
// Solidity: function skimRecipient() view returns(address)
func (_Morpho *MorphoSession) SkimRecipient() (common.Address, error) {
	return _Morpho.Contract.SkimRecipient(&_Morpho.CallOpts)
}

// SkimRecipient is a free data retrieval call binding the contract method 0x388af5b5.
//
// Solidity: function skimRecipient() view returns(address)
func (_Morpho *MorphoCallerSession) SkimRecipient() (common.Address, error) {
	return _Morpho.Contract.SkimRecipient(&_Morpho.CallOpts)
}

// SupplyQueue is a free data retrieval call binding the contract method 0xf7d18521.
//
// Solidity: function supplyQueue(uint256 ) view returns(bytes32)
func (_Morpho *MorphoCaller) SupplyQueue(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Morpho.contract.Call(opts, &out, "supplyQueue", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SupplyQueue is a free data retrieval call binding the contract method 0xf7d18521.
//
// Solidity: function supplyQueue(uint256 ) view returns(bytes32)
func (_Morpho *MorphoSession) SupplyQueue(arg0 *big.Int) ([32]byte, error) {
	return _Morpho.Contract.SupplyQueue(&_Morpho.CallOpts, arg0)
}

// SupplyQueue is a free data retrieval call binding the contract method 0xf7d18521.
//
// Solidity: function supplyQueue(uint256 ) view returns(bytes32)
func (_Morpho *MorphoCallerSession) SupplyQueue(arg0 *big.Int) ([32]byte, error) {
	return _Morpho.Contract.SupplyQueue(&_Morpho.CallOpts, arg0)
}

// SupplyQueueLength is a free data retrieval call binding the contract method 0xa17b3130.
//
// Solidity: function supplyQueueLength() view returns(uint256)
func (_Morpho *MorphoCaller) SupplyQueueLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Morpho.contract.Call(opts, &out, "supplyQueueLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyQueueLength is a free data retrieval call binding the contract method 0xa17b3130.
//
// Solidity: function supplyQueueLength() view returns(uint256)
func (_Morpho *MorphoSession) SupplyQueueLength() (*big.Int, error) {
	return _Morpho.Contract.SupplyQueueLength(&_Morpho.CallOpts)
}

// SupplyQueueLength is a free data retrieval call binding the contract method 0xa17b3130.
//
// Solidity: function supplyQueueLength() view returns(uint256)
func (_Morpho *MorphoCallerSession) SupplyQueueLength() (*big.Int, error) {
	return _Morpho.Contract.SupplyQueueLength(&_Morpho.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Morpho *MorphoCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Morpho.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Morpho *MorphoSession) Symbol() (string, error) {
	return _Morpho.Contract.Symbol(&_Morpho.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Morpho *MorphoCallerSession) Symbol() (string, error) {
	return _Morpho.Contract.Symbol(&_Morpho.CallOpts)
}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(uint256)
func (_Morpho *MorphoCaller) Timelock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Morpho.contract.Call(opts, &out, "timelock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(uint256)
func (_Morpho *MorphoSession) Timelock() (*big.Int, error) {
	return _Morpho.Contract.Timelock(&_Morpho.CallOpts)
}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(uint256)
func (_Morpho *MorphoCallerSession) Timelock() (*big.Int, error) {
	return _Morpho.Contract.Timelock(&_Morpho.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256 assets)
func (_Morpho *MorphoCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Morpho.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256 assets)
func (_Morpho *MorphoSession) TotalAssets() (*big.Int, error) {
	return _Morpho.Contract.TotalAssets(&_Morpho.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256 assets)
func (_Morpho *MorphoCallerSession) TotalAssets() (*big.Int, error) {
	return _Morpho.Contract.TotalAssets(&_Morpho.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Morpho *MorphoCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Morpho.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Morpho *MorphoSession) TotalSupply() (*big.Int, error) {
	return _Morpho.Contract.TotalSupply(&_Morpho.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Morpho *MorphoCallerSession) TotalSupply() (*big.Int, error) {
	return _Morpho.Contract.TotalSupply(&_Morpho.CallOpts)
}

// WithdrawQueue is a free data retrieval call binding the contract method 0x62518ddf.
//
// Solidity: function withdrawQueue(uint256 ) view returns(bytes32)
func (_Morpho *MorphoCaller) WithdrawQueue(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Morpho.contract.Call(opts, &out, "withdrawQueue", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WithdrawQueue is a free data retrieval call binding the contract method 0x62518ddf.
//
// Solidity: function withdrawQueue(uint256 ) view returns(bytes32)
func (_Morpho *MorphoSession) WithdrawQueue(arg0 *big.Int) ([32]byte, error) {
	return _Morpho.Contract.WithdrawQueue(&_Morpho.CallOpts, arg0)
}

// WithdrawQueue is a free data retrieval call binding the contract method 0x62518ddf.
//
// Solidity: function withdrawQueue(uint256 ) view returns(bytes32)
func (_Morpho *MorphoCallerSession) WithdrawQueue(arg0 *big.Int) ([32]byte, error) {
	return _Morpho.Contract.WithdrawQueue(&_Morpho.CallOpts, arg0)
}

// WithdrawQueueLength is a free data retrieval call binding the contract method 0x33f91ebb.
//
// Solidity: function withdrawQueueLength() view returns(uint256)
func (_Morpho *MorphoCaller) WithdrawQueueLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Morpho.contract.Call(opts, &out, "withdrawQueueLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawQueueLength is a free data retrieval call binding the contract method 0x33f91ebb.
//
// Solidity: function withdrawQueueLength() view returns(uint256)
func (_Morpho *MorphoSession) WithdrawQueueLength() (*big.Int, error) {
	return _Morpho.Contract.WithdrawQueueLength(&_Morpho.CallOpts)
}

// WithdrawQueueLength is a free data retrieval call binding the contract method 0x33f91ebb.
//
// Solidity: function withdrawQueueLength() view returns(uint256)
func (_Morpho *MorphoCallerSession) WithdrawQueueLength() (*big.Int, error) {
	return _Morpho.Contract.WithdrawQueueLength(&_Morpho.CallOpts)
}

// AcceptCap is a paid mutator transaction binding the contract method 0x6fda3868.
//
// Solidity: function acceptCap((address,address,address,address,uint256) marketParams) returns()
func (_Morpho *MorphoTransactor) AcceptCap(opts *bind.TransactOpts, marketParams MarketParams) (*types.Transaction, error) {
	return _Morpho.contract.Transact(opts, "acceptCap", marketParams)
}

// AcceptCap is a paid mutator transaction binding the contract method 0x6fda3868.
//
// Solidity: function acceptCap((address,address,address,address,uint256) marketParams) returns()
func (_Morpho *MorphoSession) AcceptCap(marketParams MarketParams) (*types.Transaction, error) {
	return _Morpho.Contract.AcceptCap(&_Morpho.TransactOpts, marketParams)
}

// AcceptCap is a paid mutator transaction binding the contract method 0x6fda3868.
//
// Solidity: function acceptCap((address,address,address,address,uint256) marketParams) returns()
func (_Morpho *MorphoTransactorSession) AcceptCap(marketParams MarketParams) (*types.Transaction, error) {
	return _Morpho.Contract.AcceptCap(&_Morpho.TransactOpts, marketParams)
}

// AcceptGuardian is a paid mutator transaction binding the contract method 0xa5f31d61.
//
// Solidity: function acceptGuardian() returns()
func (_Morpho *MorphoTransactor) AcceptGuardian(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Morpho.contract.Transact(opts, "acceptGuardian")
}

// AcceptGuardian is a paid mutator transaction binding the contract method 0xa5f31d61.
//
// Solidity: function acceptGuardian() returns()
func (_Morpho *MorphoSession) AcceptGuardian() (*types.Transaction, error) {
	return _Morpho.Contract.AcceptGuardian(&_Morpho.TransactOpts)
}

// AcceptGuardian is a paid mutator transaction binding the contract method 0xa5f31d61.
//
// Solidity: function acceptGuardian() returns()
func (_Morpho *MorphoTransactorSession) AcceptGuardian() (*types.Transaction, error) {
	return _Morpho.Contract.AcceptGuardian(&_Morpho.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Morpho *MorphoTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Morpho.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Morpho *MorphoSession) AcceptOwnership() (*types.Transaction, error) {
	return _Morpho.Contract.AcceptOwnership(&_Morpho.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Morpho *MorphoTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Morpho.Contract.AcceptOwnership(&_Morpho.TransactOpts)
}

// AcceptTimelock is a paid mutator transaction binding the contract method 0x8a2c7b39.
//
// Solidity: function acceptTimelock() returns()
func (_Morpho *MorphoTransactor) AcceptTimelock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Morpho.contract.Transact(opts, "acceptTimelock")
}

// AcceptTimelock is a paid mutator transaction binding the contract method 0x8a2c7b39.
//
// Solidity: function acceptTimelock() returns()
func (_Morpho *MorphoSession) AcceptTimelock() (*types.Transaction, error) {
	return _Morpho.Contract.AcceptTimelock(&_Morpho.TransactOpts)
}

// AcceptTimelock is a paid mutator transaction binding the contract method 0x8a2c7b39.
//
// Solidity: function acceptTimelock() returns()
func (_Morpho *MorphoTransactorSession) AcceptTimelock() (*types.Transaction, error) {
	return _Morpho.Contract.AcceptTimelock(&_Morpho.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Morpho *MorphoTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Morpho.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Morpho *MorphoSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Morpho.Contract.Approve(&_Morpho.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Morpho *MorphoTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Morpho.Contract.Approve(&_Morpho.TransactOpts, spender, value)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_Morpho *MorphoTransactor) Deposit(opts *bind.TransactOpts, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Morpho.contract.Transact(opts, "deposit", assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_Morpho *MorphoSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Morpho.Contract.Deposit(&_Morpho.TransactOpts, assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_Morpho *MorphoTransactorSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Morpho.Contract.Deposit(&_Morpho.TransactOpts, assets, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_Morpho *MorphoTransactor) Mint(opts *bind.TransactOpts, shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Morpho.contract.Transact(opts, "mint", shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_Morpho *MorphoSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Morpho.Contract.Mint(&_Morpho.TransactOpts, shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_Morpho *MorphoTransactorSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Morpho.Contract.Mint(&_Morpho.TransactOpts, shares, receiver)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Morpho *MorphoTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Morpho.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Morpho *MorphoSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Morpho.Contract.Multicall(&_Morpho.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Morpho *MorphoTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Morpho.Contract.Multicall(&_Morpho.TransactOpts, data)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Morpho *MorphoTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Morpho.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Morpho *MorphoSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Morpho.Contract.Permit(&_Morpho.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Morpho *MorphoTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Morpho.Contract.Permit(&_Morpho.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Reallocate is a paid mutator transaction binding the contract method 0x7299aa31.
//
// Solidity: function reallocate(((address,address,address,address,uint256),uint256)[] allocations) returns()
func (_Morpho *MorphoTransactor) Reallocate(opts *bind.TransactOpts, allocations []MarketAllocation) (*types.Transaction, error) {
	return _Morpho.contract.Transact(opts, "reallocate", allocations)
}

// Reallocate is a paid mutator transaction binding the contract method 0x7299aa31.
//
// Solidity: function reallocate(((address,address,address,address,uint256),uint256)[] allocations) returns()
func (_Morpho *MorphoSession) Reallocate(allocations []MarketAllocation) (*types.Transaction, error) {
	return _Morpho.Contract.Reallocate(&_Morpho.TransactOpts, allocations)
}

// Reallocate is a paid mutator transaction binding the contract method 0x7299aa31.
//
// Solidity: function reallocate(((address,address,address,address,uint256),uint256)[] allocations) returns()
func (_Morpho *MorphoTransactorSession) Reallocate(allocations []MarketAllocation) (*types.Transaction, error) {
	return _Morpho.Contract.Reallocate(&_Morpho.TransactOpts, allocations)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_Morpho *MorphoTransactor) Redeem(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Morpho.contract.Transact(opts, "redeem", shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_Morpho *MorphoSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Morpho.Contract.Redeem(&_Morpho.TransactOpts, shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_Morpho *MorphoTransactorSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Morpho.Contract.Redeem(&_Morpho.TransactOpts, shares, receiver, owner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Morpho *MorphoTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Morpho.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Morpho *MorphoSession) RenounceOwnership() (*types.Transaction, error) {
	return _Morpho.Contract.RenounceOwnership(&_Morpho.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Morpho *MorphoTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Morpho.Contract.RenounceOwnership(&_Morpho.TransactOpts)
}

// RevokePendingCap is a paid mutator transaction binding the contract method 0x102f7b6c.
//
// Solidity: function revokePendingCap(bytes32 id) returns()
func (_Morpho *MorphoTransactor) RevokePendingCap(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _Morpho.contract.Transact(opts, "revokePendingCap", id)
}

// RevokePendingCap is a paid mutator transaction binding the contract method 0x102f7b6c.
//
// Solidity: function revokePendingCap(bytes32 id) returns()
func (_Morpho *MorphoSession) RevokePendingCap(id [32]byte) (*types.Transaction, error) {
	return _Morpho.Contract.RevokePendingCap(&_Morpho.TransactOpts, id)
}

// RevokePendingCap is a paid mutator transaction binding the contract method 0x102f7b6c.
//
// Solidity: function revokePendingCap(bytes32 id) returns()
func (_Morpho *MorphoTransactorSession) RevokePendingCap(id [32]byte) (*types.Transaction, error) {
	return _Morpho.Contract.RevokePendingCap(&_Morpho.TransactOpts, id)
}

// RevokePendingGuardian is a paid mutator transaction binding the contract method 0x1ecca77c.
//
// Solidity: function revokePendingGuardian() returns()
func (_Morpho *MorphoTransactor) RevokePendingGuardian(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Morpho.contract.Transact(opts, "revokePendingGuardian")
}

// RevokePendingGuardian is a paid mutator transaction binding the contract method 0x1ecca77c.
//
// Solidity: function revokePendingGuardian() returns()
func (_Morpho *MorphoSession) RevokePendingGuardian() (*types.Transaction, error) {
	return _Morpho.Contract.RevokePendingGuardian(&_Morpho.TransactOpts)
}

// RevokePendingGuardian is a paid mutator transaction binding the contract method 0x1ecca77c.
//
// Solidity: function revokePendingGuardian() returns()
func (_Morpho *MorphoTransactorSession) RevokePendingGuardian() (*types.Transaction, error) {
	return _Morpho.Contract.RevokePendingGuardian(&_Morpho.TransactOpts)
}

// RevokePendingMarketRemoval is a paid mutator transaction binding the contract method 0x4b998de5.
//
// Solidity: function revokePendingMarketRemoval(bytes32 id) returns()
func (_Morpho *MorphoTransactor) RevokePendingMarketRemoval(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _Morpho.contract.Transact(opts, "revokePendingMarketRemoval", id)
}

// RevokePendingMarketRemoval is a paid mutator transaction binding the contract method 0x4b998de5.
//
// Solidity: function revokePendingMarketRemoval(bytes32 id) returns()
func (_Morpho *MorphoSession) RevokePendingMarketRemoval(id [32]byte) (*types.Transaction, error) {
	return _Morpho.Contract.RevokePendingMarketRemoval(&_Morpho.TransactOpts, id)
}

// RevokePendingMarketRemoval is a paid mutator transaction binding the contract method 0x4b998de5.
//
// Solidity: function revokePendingMarketRemoval(bytes32 id) returns()
func (_Morpho *MorphoTransactorSession) RevokePendingMarketRemoval(id [32]byte) (*types.Transaction, error) {
	return _Morpho.Contract.RevokePendingMarketRemoval(&_Morpho.TransactOpts, id)
}

// RevokePendingTimelock is a paid mutator transaction binding the contract method 0xc9649aa9.
//
// Solidity: function revokePendingTimelock() returns()
func (_Morpho *MorphoTransactor) RevokePendingTimelock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Morpho.contract.Transact(opts, "revokePendingTimelock")
}

// RevokePendingTimelock is a paid mutator transaction binding the contract method 0xc9649aa9.
//
// Solidity: function revokePendingTimelock() returns()
func (_Morpho *MorphoSession) RevokePendingTimelock() (*types.Transaction, error) {
	return _Morpho.Contract.RevokePendingTimelock(&_Morpho.TransactOpts)
}

// RevokePendingTimelock is a paid mutator transaction binding the contract method 0xc9649aa9.
//
// Solidity: function revokePendingTimelock() returns()
func (_Morpho *MorphoTransactorSession) RevokePendingTimelock() (*types.Transaction, error) {
	return _Morpho.Contract.RevokePendingTimelock(&_Morpho.TransactOpts)
}

// SetCurator is a paid mutator transaction binding the contract method 0xe90956cf.
//
// Solidity: function setCurator(address newCurator) returns()
func (_Morpho *MorphoTransactor) SetCurator(opts *bind.TransactOpts, newCurator common.Address) (*types.Transaction, error) {
	return _Morpho.contract.Transact(opts, "setCurator", newCurator)
}

// SetCurator is a paid mutator transaction binding the contract method 0xe90956cf.
//
// Solidity: function setCurator(address newCurator) returns()
func (_Morpho *MorphoSession) SetCurator(newCurator common.Address) (*types.Transaction, error) {
	return _Morpho.Contract.SetCurator(&_Morpho.TransactOpts, newCurator)
}

// SetCurator is a paid mutator transaction binding the contract method 0xe90956cf.
//
// Solidity: function setCurator(address newCurator) returns()
func (_Morpho *MorphoTransactorSession) SetCurator(newCurator common.Address) (*types.Transaction, error) {
	return _Morpho.Contract.SetCurator(&_Morpho.TransactOpts, newCurator)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 newFee) returns()
func (_Morpho *MorphoTransactor) SetFee(opts *bind.TransactOpts, newFee *big.Int) (*types.Transaction, error) {
	return _Morpho.contract.Transact(opts, "setFee", newFee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 newFee) returns()
func (_Morpho *MorphoSession) SetFee(newFee *big.Int) (*types.Transaction, error) {
	return _Morpho.Contract.SetFee(&_Morpho.TransactOpts, newFee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 newFee) returns()
func (_Morpho *MorphoTransactorSession) SetFee(newFee *big.Int) (*types.Transaction, error) {
	return _Morpho.Contract.SetFee(&_Morpho.TransactOpts, newFee)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address newFeeRecipient) returns()
func (_Morpho *MorphoTransactor) SetFeeRecipient(opts *bind.TransactOpts, newFeeRecipient common.Address) (*types.Transaction, error) {
	return _Morpho.contract.Transact(opts, "setFeeRecipient", newFeeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address newFeeRecipient) returns()
func (_Morpho *MorphoSession) SetFeeRecipient(newFeeRecipient common.Address) (*types.Transaction, error) {
	return _Morpho.Contract.SetFeeRecipient(&_Morpho.TransactOpts, newFeeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address newFeeRecipient) returns()
func (_Morpho *MorphoTransactorSession) SetFeeRecipient(newFeeRecipient common.Address) (*types.Transaction, error) {
	return _Morpho.Contract.SetFeeRecipient(&_Morpho.TransactOpts, newFeeRecipient)
}

// SetIsAllocator is a paid mutator transaction binding the contract method 0xb192a84a.
//
// Solidity: function setIsAllocator(address newAllocator, bool newIsAllocator) returns()
func (_Morpho *MorphoTransactor) SetIsAllocator(opts *bind.TransactOpts, newAllocator common.Address, newIsAllocator bool) (*types.Transaction, error) {
	return _Morpho.contract.Transact(opts, "setIsAllocator", newAllocator, newIsAllocator)
}

// SetIsAllocator is a paid mutator transaction binding the contract method 0xb192a84a.
//
// Solidity: function setIsAllocator(address newAllocator, bool newIsAllocator) returns()
func (_Morpho *MorphoSession) SetIsAllocator(newAllocator common.Address, newIsAllocator bool) (*types.Transaction, error) {
	return _Morpho.Contract.SetIsAllocator(&_Morpho.TransactOpts, newAllocator, newIsAllocator)
}

// SetIsAllocator is a paid mutator transaction binding the contract method 0xb192a84a.
//
// Solidity: function setIsAllocator(address newAllocator, bool newIsAllocator) returns()
func (_Morpho *MorphoTransactorSession) SetIsAllocator(newAllocator common.Address, newIsAllocator bool) (*types.Transaction, error) {
	return _Morpho.Contract.SetIsAllocator(&_Morpho.TransactOpts, newAllocator, newIsAllocator)
}

// SetSkimRecipient is a paid mutator transaction binding the contract method 0x2b30997b.
//
// Solidity: function setSkimRecipient(address newSkimRecipient) returns()
func (_Morpho *MorphoTransactor) SetSkimRecipient(opts *bind.TransactOpts, newSkimRecipient common.Address) (*types.Transaction, error) {
	return _Morpho.contract.Transact(opts, "setSkimRecipient", newSkimRecipient)
}

// SetSkimRecipient is a paid mutator transaction binding the contract method 0x2b30997b.
//
// Solidity: function setSkimRecipient(address newSkimRecipient) returns()
func (_Morpho *MorphoSession) SetSkimRecipient(newSkimRecipient common.Address) (*types.Transaction, error) {
	return _Morpho.Contract.SetSkimRecipient(&_Morpho.TransactOpts, newSkimRecipient)
}

// SetSkimRecipient is a paid mutator transaction binding the contract method 0x2b30997b.
//
// Solidity: function setSkimRecipient(address newSkimRecipient) returns()
func (_Morpho *MorphoTransactorSession) SetSkimRecipient(newSkimRecipient common.Address) (*types.Transaction, error) {
	return _Morpho.Contract.SetSkimRecipient(&_Morpho.TransactOpts, newSkimRecipient)
}

// SetSupplyQueue is a paid mutator transaction binding the contract method 0x2acc56f9.
//
// Solidity: function setSupplyQueue(bytes32[] newSupplyQueue) returns()
func (_Morpho *MorphoTransactor) SetSupplyQueue(opts *bind.TransactOpts, newSupplyQueue [][32]byte) (*types.Transaction, error) {
	return _Morpho.contract.Transact(opts, "setSupplyQueue", newSupplyQueue)
}

// SetSupplyQueue is a paid mutator transaction binding the contract method 0x2acc56f9.
//
// Solidity: function setSupplyQueue(bytes32[] newSupplyQueue) returns()
func (_Morpho *MorphoSession) SetSupplyQueue(newSupplyQueue [][32]byte) (*types.Transaction, error) {
	return _Morpho.Contract.SetSupplyQueue(&_Morpho.TransactOpts, newSupplyQueue)
}

// SetSupplyQueue is a paid mutator transaction binding the contract method 0x2acc56f9.
//
// Solidity: function setSupplyQueue(bytes32[] newSupplyQueue) returns()
func (_Morpho *MorphoTransactorSession) SetSupplyQueue(newSupplyQueue [][32]byte) (*types.Transaction, error) {
	return _Morpho.Contract.SetSupplyQueue(&_Morpho.TransactOpts, newSupplyQueue)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address token) returns()
func (_Morpho *MorphoTransactor) Skim(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Morpho.contract.Transact(opts, "skim", token)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address token) returns()
func (_Morpho *MorphoSession) Skim(token common.Address) (*types.Transaction, error) {
	return _Morpho.Contract.Skim(&_Morpho.TransactOpts, token)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address token) returns()
func (_Morpho *MorphoTransactorSession) Skim(token common.Address) (*types.Transaction, error) {
	return _Morpho.Contract.Skim(&_Morpho.TransactOpts, token)
}

// SubmitCap is a paid mutator transaction binding the contract method 0x3b24c2bf.
//
// Solidity: function submitCap((address,address,address,address,uint256) marketParams, uint256 newSupplyCap) returns()
func (_Morpho *MorphoTransactor) SubmitCap(opts *bind.TransactOpts, marketParams MarketParams, newSupplyCap *big.Int) (*types.Transaction, error) {
	return _Morpho.contract.Transact(opts, "submitCap", marketParams, newSupplyCap)
}

// SubmitCap is a paid mutator transaction binding the contract method 0x3b24c2bf.
//
// Solidity: function submitCap((address,address,address,address,uint256) marketParams, uint256 newSupplyCap) returns()
func (_Morpho *MorphoSession) SubmitCap(marketParams MarketParams, newSupplyCap *big.Int) (*types.Transaction, error) {
	return _Morpho.Contract.SubmitCap(&_Morpho.TransactOpts, marketParams, newSupplyCap)
}

// SubmitCap is a paid mutator transaction binding the contract method 0x3b24c2bf.
//
// Solidity: function submitCap((address,address,address,address,uint256) marketParams, uint256 newSupplyCap) returns()
func (_Morpho *MorphoTransactorSession) SubmitCap(marketParams MarketParams, newSupplyCap *big.Int) (*types.Transaction, error) {
	return _Morpho.Contract.SubmitCap(&_Morpho.TransactOpts, marketParams, newSupplyCap)
}

// SubmitGuardian is a paid mutator transaction binding the contract method 0x9d6b4a45.
//
// Solidity: function submitGuardian(address newGuardian) returns()
func (_Morpho *MorphoTransactor) SubmitGuardian(opts *bind.TransactOpts, newGuardian common.Address) (*types.Transaction, error) {
	return _Morpho.contract.Transact(opts, "submitGuardian", newGuardian)
}

// SubmitGuardian is a paid mutator transaction binding the contract method 0x9d6b4a45.
//
// Solidity: function submitGuardian(address newGuardian) returns()
func (_Morpho *MorphoSession) SubmitGuardian(newGuardian common.Address) (*types.Transaction, error) {
	return _Morpho.Contract.SubmitGuardian(&_Morpho.TransactOpts, newGuardian)
}

// SubmitGuardian is a paid mutator transaction binding the contract method 0x9d6b4a45.
//
// Solidity: function submitGuardian(address newGuardian) returns()
func (_Morpho *MorphoTransactorSession) SubmitGuardian(newGuardian common.Address) (*types.Transaction, error) {
	return _Morpho.Contract.SubmitGuardian(&_Morpho.TransactOpts, newGuardian)
}

// SubmitMarketRemoval is a paid mutator transaction binding the contract method 0x84755b5f.
//
// Solidity: function submitMarketRemoval((address,address,address,address,uint256) marketParams) returns()
func (_Morpho *MorphoTransactor) SubmitMarketRemoval(opts *bind.TransactOpts, marketParams MarketParams) (*types.Transaction, error) {
	return _Morpho.contract.Transact(opts, "submitMarketRemoval", marketParams)
}

// SubmitMarketRemoval is a paid mutator transaction binding the contract method 0x84755b5f.
//
// Solidity: function submitMarketRemoval((address,address,address,address,uint256) marketParams) returns()
func (_Morpho *MorphoSession) SubmitMarketRemoval(marketParams MarketParams) (*types.Transaction, error) {
	return _Morpho.Contract.SubmitMarketRemoval(&_Morpho.TransactOpts, marketParams)
}

// SubmitMarketRemoval is a paid mutator transaction binding the contract method 0x84755b5f.
//
// Solidity: function submitMarketRemoval((address,address,address,address,uint256) marketParams) returns()
func (_Morpho *MorphoTransactorSession) SubmitMarketRemoval(marketParams MarketParams) (*types.Transaction, error) {
	return _Morpho.Contract.SubmitMarketRemoval(&_Morpho.TransactOpts, marketParams)
}

// SubmitTimelock is a paid mutator transaction binding the contract method 0x7224a512.
//
// Solidity: function submitTimelock(uint256 newTimelock) returns()
func (_Morpho *MorphoTransactor) SubmitTimelock(opts *bind.TransactOpts, newTimelock *big.Int) (*types.Transaction, error) {
	return _Morpho.contract.Transact(opts, "submitTimelock", newTimelock)
}

// SubmitTimelock is a paid mutator transaction binding the contract method 0x7224a512.
//
// Solidity: function submitTimelock(uint256 newTimelock) returns()
func (_Morpho *MorphoSession) SubmitTimelock(newTimelock *big.Int) (*types.Transaction, error) {
	return _Morpho.Contract.SubmitTimelock(&_Morpho.TransactOpts, newTimelock)
}

// SubmitTimelock is a paid mutator transaction binding the contract method 0x7224a512.
//
// Solidity: function submitTimelock(uint256 newTimelock) returns()
func (_Morpho *MorphoTransactorSession) SubmitTimelock(newTimelock *big.Int) (*types.Transaction, error) {
	return _Morpho.Contract.SubmitTimelock(&_Morpho.TransactOpts, newTimelock)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Morpho *MorphoTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Morpho.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Morpho *MorphoSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Morpho.Contract.Transfer(&_Morpho.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Morpho *MorphoTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Morpho.Contract.Transfer(&_Morpho.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Morpho *MorphoTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Morpho.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Morpho *MorphoSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Morpho.Contract.TransferFrom(&_Morpho.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Morpho *MorphoTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Morpho.Contract.TransferFrom(&_Morpho.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Morpho *MorphoTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Morpho.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Morpho *MorphoSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Morpho.Contract.TransferOwnership(&_Morpho.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Morpho *MorphoTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Morpho.Contract.TransferOwnership(&_Morpho.TransactOpts, newOwner)
}

// UpdateWithdrawQueue is a paid mutator transaction binding the contract method 0x41b67833.
//
// Solidity: function updateWithdrawQueue(uint256[] indexes) returns()
func (_Morpho *MorphoTransactor) UpdateWithdrawQueue(opts *bind.TransactOpts, indexes []*big.Int) (*types.Transaction, error) {
	return _Morpho.contract.Transact(opts, "updateWithdrawQueue", indexes)
}

// UpdateWithdrawQueue is a paid mutator transaction binding the contract method 0x41b67833.
//
// Solidity: function updateWithdrawQueue(uint256[] indexes) returns()
func (_Morpho *MorphoSession) UpdateWithdrawQueue(indexes []*big.Int) (*types.Transaction, error) {
	return _Morpho.Contract.UpdateWithdrawQueue(&_Morpho.TransactOpts, indexes)
}

// UpdateWithdrawQueue is a paid mutator transaction binding the contract method 0x41b67833.
//
// Solidity: function updateWithdrawQueue(uint256[] indexes) returns()
func (_Morpho *MorphoTransactorSession) UpdateWithdrawQueue(indexes []*big.Int) (*types.Transaction, error) {
	return _Morpho.Contract.UpdateWithdrawQueue(&_Morpho.TransactOpts, indexes)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_Morpho *MorphoTransactor) Withdraw(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Morpho.contract.Transact(opts, "withdraw", assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_Morpho *MorphoSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Morpho.Contract.Withdraw(&_Morpho.TransactOpts, assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_Morpho *MorphoTransactorSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Morpho.Contract.Withdraw(&_Morpho.TransactOpts, assets, receiver, owner)
}

// MorphoAccrueInterestIterator is returned from FilterAccrueInterest and is used to iterate over the raw logs and unpacked data for AccrueInterest events raised by the Morpho contract.
type MorphoAccrueInterestIterator struct {
	Event *MorphoAccrueInterest // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoAccrueInterestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoAccrueInterest)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoAccrueInterest)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoAccrueInterestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoAccrueInterestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoAccrueInterest represents a AccrueInterest event raised by the Morpho contract.
type MorphoAccrueInterest struct {
	NewTotalAssets *big.Int
	FeeShares      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAccrueInterest is a free log retrieval operation binding the contract event 0xf66f28b40975dbb933913542c7e6a0f50a1d0f20aa74ea6e0efe65ab616323ec.
//
// Solidity: event AccrueInterest(uint256 newTotalAssets, uint256 feeShares)
func (_Morpho *MorphoFilterer) FilterAccrueInterest(opts *bind.FilterOpts) (*MorphoAccrueInterestIterator, error) {

	logs, sub, err := _Morpho.contract.FilterLogs(opts, "AccrueInterest")
	if err != nil {
		return nil, err
	}
	return &MorphoAccrueInterestIterator{contract: _Morpho.contract, event: "AccrueInterest", logs: logs, sub: sub}, nil
}

// WatchAccrueInterest is a free log subscription operation binding the contract event 0xf66f28b40975dbb933913542c7e6a0f50a1d0f20aa74ea6e0efe65ab616323ec.
//
// Solidity: event AccrueInterest(uint256 newTotalAssets, uint256 feeShares)
func (_Morpho *MorphoFilterer) WatchAccrueInterest(opts *bind.WatchOpts, sink chan<- *MorphoAccrueInterest) (event.Subscription, error) {

	logs, sub, err := _Morpho.contract.WatchLogs(opts, "AccrueInterest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoAccrueInterest)
				if err := _Morpho.contract.UnpackLog(event, "AccrueInterest", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAccrueInterest is a log parse operation binding the contract event 0xf66f28b40975dbb933913542c7e6a0f50a1d0f20aa74ea6e0efe65ab616323ec.
//
// Solidity: event AccrueInterest(uint256 newTotalAssets, uint256 feeShares)
func (_Morpho *MorphoFilterer) ParseAccrueInterest(log types.Log) (*MorphoAccrueInterest, error) {
	event := new(MorphoAccrueInterest)
	if err := _Morpho.contract.UnpackLog(event, "AccrueInterest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Morpho contract.
type MorphoApprovalIterator struct {
	Event *MorphoApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoApproval represents a Approval event raised by the Morpho contract.
type MorphoApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Morpho *MorphoFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MorphoApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Morpho.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MorphoApprovalIterator{contract: _Morpho.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Morpho *MorphoFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MorphoApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Morpho.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoApproval)
				if err := _Morpho.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Morpho *MorphoFilterer) ParseApproval(log types.Log) (*MorphoApproval, error) {
	event := new(MorphoApproval)
	if err := _Morpho.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Morpho contract.
type MorphoDepositIterator struct {
	Event *MorphoDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoDeposit represents a Deposit event raised by the Morpho contract.
type MorphoDeposit struct {
	Sender common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_Morpho *MorphoFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, owner []common.Address) (*MorphoDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Morpho.contract.FilterLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &MorphoDepositIterator{contract: _Morpho.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_Morpho *MorphoFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *MorphoDeposit, sender []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Morpho.contract.WatchLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoDeposit)
				if err := _Morpho.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_Morpho *MorphoFilterer) ParseDeposit(log types.Log) (*MorphoDeposit, error) {
	event := new(MorphoDeposit)
	if err := _Morpho.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the Morpho contract.
type MorphoEIP712DomainChangedIterator struct {
	Event *MorphoEIP712DomainChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoEIP712DomainChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoEIP712DomainChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoEIP712DomainChanged represents a EIP712DomainChanged event raised by the Morpho contract.
type MorphoEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Morpho *MorphoFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*MorphoEIP712DomainChangedIterator, error) {

	logs, sub, err := _Morpho.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &MorphoEIP712DomainChangedIterator{contract: _Morpho.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Morpho *MorphoFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *MorphoEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _Morpho.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoEIP712DomainChanged)
				if err := _Morpho.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Morpho *MorphoFilterer) ParseEIP712DomainChanged(log types.Log) (*MorphoEIP712DomainChanged, error) {
	event := new(MorphoEIP712DomainChanged)
	if err := _Morpho.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the Morpho contract.
type MorphoOwnershipTransferStartedIterator struct {
	Event *MorphoOwnershipTransferStarted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoOwnershipTransferStarted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoOwnershipTransferStarted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the Morpho contract.
type MorphoOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Morpho *MorphoFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MorphoOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Morpho.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MorphoOwnershipTransferStartedIterator{contract: _Morpho.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Morpho *MorphoFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *MorphoOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Morpho.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoOwnershipTransferStarted)
				if err := _Morpho.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Morpho *MorphoFilterer) ParseOwnershipTransferStarted(log types.Log) (*MorphoOwnershipTransferStarted, error) {
	event := new(MorphoOwnershipTransferStarted)
	if err := _Morpho.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Morpho contract.
type MorphoOwnershipTransferredIterator struct {
	Event *MorphoOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoOwnershipTransferred represents a OwnershipTransferred event raised by the Morpho contract.
type MorphoOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Morpho *MorphoFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MorphoOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Morpho.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MorphoOwnershipTransferredIterator{contract: _Morpho.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Morpho *MorphoFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MorphoOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Morpho.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoOwnershipTransferred)
				if err := _Morpho.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Morpho *MorphoFilterer) ParseOwnershipTransferred(log types.Log) (*MorphoOwnershipTransferred, error) {
	event := new(MorphoOwnershipTransferred)
	if err := _Morpho.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoReallocateSupplyIterator is returned from FilterReallocateSupply and is used to iterate over the raw logs and unpacked data for ReallocateSupply events raised by the Morpho contract.
type MorphoReallocateSupplyIterator struct {
	Event *MorphoReallocateSupply // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoReallocateSupplyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoReallocateSupply)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoReallocateSupply)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoReallocateSupplyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoReallocateSupplyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoReallocateSupply represents a ReallocateSupply event raised by the Morpho contract.
type MorphoReallocateSupply struct {
	Caller         common.Address
	Id             [32]byte
	SuppliedAssets *big.Int
	SuppliedShares *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterReallocateSupply is a free log retrieval operation binding the contract event 0x89bf199df65bf65155e3e0a8abc4ad4a1be606220c8295840dba2ab5656c1f6d.
//
// Solidity: event ReallocateSupply(address indexed caller, bytes32 indexed id, uint256 suppliedAssets, uint256 suppliedShares)
func (_Morpho *MorphoFilterer) FilterReallocateSupply(opts *bind.FilterOpts, caller []common.Address, id [][32]byte) (*MorphoReallocateSupplyIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Morpho.contract.FilterLogs(opts, "ReallocateSupply", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return &MorphoReallocateSupplyIterator{contract: _Morpho.contract, event: "ReallocateSupply", logs: logs, sub: sub}, nil
}

// WatchReallocateSupply is a free log subscription operation binding the contract event 0x89bf199df65bf65155e3e0a8abc4ad4a1be606220c8295840dba2ab5656c1f6d.
//
// Solidity: event ReallocateSupply(address indexed caller, bytes32 indexed id, uint256 suppliedAssets, uint256 suppliedShares)
func (_Morpho *MorphoFilterer) WatchReallocateSupply(opts *bind.WatchOpts, sink chan<- *MorphoReallocateSupply, caller []common.Address, id [][32]byte) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Morpho.contract.WatchLogs(opts, "ReallocateSupply", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoReallocateSupply)
				if err := _Morpho.contract.UnpackLog(event, "ReallocateSupply", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReallocateSupply is a log parse operation binding the contract event 0x89bf199df65bf65155e3e0a8abc4ad4a1be606220c8295840dba2ab5656c1f6d.
//
// Solidity: event ReallocateSupply(address indexed caller, bytes32 indexed id, uint256 suppliedAssets, uint256 suppliedShares)
func (_Morpho *MorphoFilterer) ParseReallocateSupply(log types.Log) (*MorphoReallocateSupply, error) {
	event := new(MorphoReallocateSupply)
	if err := _Morpho.contract.UnpackLog(event, "ReallocateSupply", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoReallocateWithdrawIterator is returned from FilterReallocateWithdraw and is used to iterate over the raw logs and unpacked data for ReallocateWithdraw events raised by the Morpho contract.
type MorphoReallocateWithdrawIterator struct {
	Event *MorphoReallocateWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoReallocateWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoReallocateWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoReallocateWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoReallocateWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoReallocateWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoReallocateWithdraw represents a ReallocateWithdraw event raised by the Morpho contract.
type MorphoReallocateWithdraw struct {
	Caller          common.Address
	Id              [32]byte
	WithdrawnAssets *big.Int
	WithdrawnShares *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterReallocateWithdraw is a free log retrieval operation binding the contract event 0xdd8bf5226dff861316e0fa7863fdb7dc7b87c614eb29a135f524eb79d5a1189a.
//
// Solidity: event ReallocateWithdraw(address indexed caller, bytes32 indexed id, uint256 withdrawnAssets, uint256 withdrawnShares)
func (_Morpho *MorphoFilterer) FilterReallocateWithdraw(opts *bind.FilterOpts, caller []common.Address, id [][32]byte) (*MorphoReallocateWithdrawIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Morpho.contract.FilterLogs(opts, "ReallocateWithdraw", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return &MorphoReallocateWithdrawIterator{contract: _Morpho.contract, event: "ReallocateWithdraw", logs: logs, sub: sub}, nil
}

// WatchReallocateWithdraw is a free log subscription operation binding the contract event 0xdd8bf5226dff861316e0fa7863fdb7dc7b87c614eb29a135f524eb79d5a1189a.
//
// Solidity: event ReallocateWithdraw(address indexed caller, bytes32 indexed id, uint256 withdrawnAssets, uint256 withdrawnShares)
func (_Morpho *MorphoFilterer) WatchReallocateWithdraw(opts *bind.WatchOpts, sink chan<- *MorphoReallocateWithdraw, caller []common.Address, id [][32]byte) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Morpho.contract.WatchLogs(opts, "ReallocateWithdraw", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoReallocateWithdraw)
				if err := _Morpho.contract.UnpackLog(event, "ReallocateWithdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReallocateWithdraw is a log parse operation binding the contract event 0xdd8bf5226dff861316e0fa7863fdb7dc7b87c614eb29a135f524eb79d5a1189a.
//
// Solidity: event ReallocateWithdraw(address indexed caller, bytes32 indexed id, uint256 withdrawnAssets, uint256 withdrawnShares)
func (_Morpho *MorphoFilterer) ParseReallocateWithdraw(log types.Log) (*MorphoReallocateWithdraw, error) {
	event := new(MorphoReallocateWithdraw)
	if err := _Morpho.contract.UnpackLog(event, "ReallocateWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoRevokePendingCapIterator is returned from FilterRevokePendingCap and is used to iterate over the raw logs and unpacked data for RevokePendingCap events raised by the Morpho contract.
type MorphoRevokePendingCapIterator struct {
	Event *MorphoRevokePendingCap // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoRevokePendingCapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoRevokePendingCap)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoRevokePendingCap)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoRevokePendingCapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoRevokePendingCapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoRevokePendingCap represents a RevokePendingCap event raised by the Morpho contract.
type MorphoRevokePendingCap struct {
	Caller common.Address
	Id     [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRevokePendingCap is a free log retrieval operation binding the contract event 0x1026ceca5ed3747eb5edec555732d4a6f901ce1a875ecf981064628cadde1120.
//
// Solidity: event RevokePendingCap(address indexed caller, bytes32 indexed id)
func (_Morpho *MorphoFilterer) FilterRevokePendingCap(opts *bind.FilterOpts, caller []common.Address, id [][32]byte) (*MorphoRevokePendingCapIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Morpho.contract.FilterLogs(opts, "RevokePendingCap", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return &MorphoRevokePendingCapIterator{contract: _Morpho.contract, event: "RevokePendingCap", logs: logs, sub: sub}, nil
}

// WatchRevokePendingCap is a free log subscription operation binding the contract event 0x1026ceca5ed3747eb5edec555732d4a6f901ce1a875ecf981064628cadde1120.
//
// Solidity: event RevokePendingCap(address indexed caller, bytes32 indexed id)
func (_Morpho *MorphoFilterer) WatchRevokePendingCap(opts *bind.WatchOpts, sink chan<- *MorphoRevokePendingCap, caller []common.Address, id [][32]byte) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Morpho.contract.WatchLogs(opts, "RevokePendingCap", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoRevokePendingCap)
				if err := _Morpho.contract.UnpackLog(event, "RevokePendingCap", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRevokePendingCap is a log parse operation binding the contract event 0x1026ceca5ed3747eb5edec555732d4a6f901ce1a875ecf981064628cadde1120.
//
// Solidity: event RevokePendingCap(address indexed caller, bytes32 indexed id)
func (_Morpho *MorphoFilterer) ParseRevokePendingCap(log types.Log) (*MorphoRevokePendingCap, error) {
	event := new(MorphoRevokePendingCap)
	if err := _Morpho.contract.UnpackLog(event, "RevokePendingCap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoRevokePendingGuardianIterator is returned from FilterRevokePendingGuardian and is used to iterate over the raw logs and unpacked data for RevokePendingGuardian events raised by the Morpho contract.
type MorphoRevokePendingGuardianIterator struct {
	Event *MorphoRevokePendingGuardian // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoRevokePendingGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoRevokePendingGuardian)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoRevokePendingGuardian)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoRevokePendingGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoRevokePendingGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoRevokePendingGuardian represents a RevokePendingGuardian event raised by the Morpho contract.
type MorphoRevokePendingGuardian struct {
	Caller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRevokePendingGuardian is a free log retrieval operation binding the contract event 0xc40a085ccfa20f5fd518ade5c3a77a7ecbdfbb4c75efcdca6146a8e3c841d663.
//
// Solidity: event RevokePendingGuardian(address indexed caller)
func (_Morpho *MorphoFilterer) FilterRevokePendingGuardian(opts *bind.FilterOpts, caller []common.Address) (*MorphoRevokePendingGuardianIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Morpho.contract.FilterLogs(opts, "RevokePendingGuardian", callerRule)
	if err != nil {
		return nil, err
	}
	return &MorphoRevokePendingGuardianIterator{contract: _Morpho.contract, event: "RevokePendingGuardian", logs: logs, sub: sub}, nil
}

// WatchRevokePendingGuardian is a free log subscription operation binding the contract event 0xc40a085ccfa20f5fd518ade5c3a77a7ecbdfbb4c75efcdca6146a8e3c841d663.
//
// Solidity: event RevokePendingGuardian(address indexed caller)
func (_Morpho *MorphoFilterer) WatchRevokePendingGuardian(opts *bind.WatchOpts, sink chan<- *MorphoRevokePendingGuardian, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Morpho.contract.WatchLogs(opts, "RevokePendingGuardian", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoRevokePendingGuardian)
				if err := _Morpho.contract.UnpackLog(event, "RevokePendingGuardian", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRevokePendingGuardian is a log parse operation binding the contract event 0xc40a085ccfa20f5fd518ade5c3a77a7ecbdfbb4c75efcdca6146a8e3c841d663.
//
// Solidity: event RevokePendingGuardian(address indexed caller)
func (_Morpho *MorphoFilterer) ParseRevokePendingGuardian(log types.Log) (*MorphoRevokePendingGuardian, error) {
	event := new(MorphoRevokePendingGuardian)
	if err := _Morpho.contract.UnpackLog(event, "RevokePendingGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoRevokePendingMarketRemovalIterator is returned from FilterRevokePendingMarketRemoval and is used to iterate over the raw logs and unpacked data for RevokePendingMarketRemoval events raised by the Morpho contract.
type MorphoRevokePendingMarketRemovalIterator struct {
	Event *MorphoRevokePendingMarketRemoval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoRevokePendingMarketRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoRevokePendingMarketRemoval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoRevokePendingMarketRemoval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoRevokePendingMarketRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoRevokePendingMarketRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoRevokePendingMarketRemoval represents a RevokePendingMarketRemoval event raised by the Morpho contract.
type MorphoRevokePendingMarketRemoval struct {
	Caller common.Address
	Id     [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRevokePendingMarketRemoval is a free log retrieval operation binding the contract event 0xcbeb8ecdaa5a3c133e62219b63bfc35bce3fda13065d2bed32e3b7dde60a59f4.
//
// Solidity: event RevokePendingMarketRemoval(address indexed caller, bytes32 indexed id)
func (_Morpho *MorphoFilterer) FilterRevokePendingMarketRemoval(opts *bind.FilterOpts, caller []common.Address, id [][32]byte) (*MorphoRevokePendingMarketRemovalIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Morpho.contract.FilterLogs(opts, "RevokePendingMarketRemoval", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return &MorphoRevokePendingMarketRemovalIterator{contract: _Morpho.contract, event: "RevokePendingMarketRemoval", logs: logs, sub: sub}, nil
}

// WatchRevokePendingMarketRemoval is a free log subscription operation binding the contract event 0xcbeb8ecdaa5a3c133e62219b63bfc35bce3fda13065d2bed32e3b7dde60a59f4.
//
// Solidity: event RevokePendingMarketRemoval(address indexed caller, bytes32 indexed id)
func (_Morpho *MorphoFilterer) WatchRevokePendingMarketRemoval(opts *bind.WatchOpts, sink chan<- *MorphoRevokePendingMarketRemoval, caller []common.Address, id [][32]byte) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Morpho.contract.WatchLogs(opts, "RevokePendingMarketRemoval", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoRevokePendingMarketRemoval)
				if err := _Morpho.contract.UnpackLog(event, "RevokePendingMarketRemoval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRevokePendingMarketRemoval is a log parse operation binding the contract event 0xcbeb8ecdaa5a3c133e62219b63bfc35bce3fda13065d2bed32e3b7dde60a59f4.
//
// Solidity: event RevokePendingMarketRemoval(address indexed caller, bytes32 indexed id)
func (_Morpho *MorphoFilterer) ParseRevokePendingMarketRemoval(log types.Log) (*MorphoRevokePendingMarketRemoval, error) {
	event := new(MorphoRevokePendingMarketRemoval)
	if err := _Morpho.contract.UnpackLog(event, "RevokePendingMarketRemoval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoRevokePendingTimelockIterator is returned from FilterRevokePendingTimelock and is used to iterate over the raw logs and unpacked data for RevokePendingTimelock events raised by the Morpho contract.
type MorphoRevokePendingTimelockIterator struct {
	Event *MorphoRevokePendingTimelock // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoRevokePendingTimelockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoRevokePendingTimelock)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoRevokePendingTimelock)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoRevokePendingTimelockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoRevokePendingTimelockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoRevokePendingTimelock represents a RevokePendingTimelock event raised by the Morpho contract.
type MorphoRevokePendingTimelock struct {
	Caller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRevokePendingTimelock is a free log retrieval operation binding the contract event 0x921828337692c347c634c5d2aacbc7b756014674bd236f3cc2058d8e284a951b.
//
// Solidity: event RevokePendingTimelock(address indexed caller)
func (_Morpho *MorphoFilterer) FilterRevokePendingTimelock(opts *bind.FilterOpts, caller []common.Address) (*MorphoRevokePendingTimelockIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Morpho.contract.FilterLogs(opts, "RevokePendingTimelock", callerRule)
	if err != nil {
		return nil, err
	}
	return &MorphoRevokePendingTimelockIterator{contract: _Morpho.contract, event: "RevokePendingTimelock", logs: logs, sub: sub}, nil
}

// WatchRevokePendingTimelock is a free log subscription operation binding the contract event 0x921828337692c347c634c5d2aacbc7b756014674bd236f3cc2058d8e284a951b.
//
// Solidity: event RevokePendingTimelock(address indexed caller)
func (_Morpho *MorphoFilterer) WatchRevokePendingTimelock(opts *bind.WatchOpts, sink chan<- *MorphoRevokePendingTimelock, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Morpho.contract.WatchLogs(opts, "RevokePendingTimelock", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoRevokePendingTimelock)
				if err := _Morpho.contract.UnpackLog(event, "RevokePendingTimelock", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRevokePendingTimelock is a log parse operation binding the contract event 0x921828337692c347c634c5d2aacbc7b756014674bd236f3cc2058d8e284a951b.
//
// Solidity: event RevokePendingTimelock(address indexed caller)
func (_Morpho *MorphoFilterer) ParseRevokePendingTimelock(log types.Log) (*MorphoRevokePendingTimelock, error) {
	event := new(MorphoRevokePendingTimelock)
	if err := _Morpho.contract.UnpackLog(event, "RevokePendingTimelock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoSetCapIterator is returned from FilterSetCap and is used to iterate over the raw logs and unpacked data for SetCap events raised by the Morpho contract.
type MorphoSetCapIterator struct {
	Event *MorphoSetCap // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoSetCapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoSetCap)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoSetCap)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoSetCapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoSetCapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoSetCap represents a SetCap event raised by the Morpho contract.
type MorphoSetCap struct {
	Caller common.Address
	Id     [32]byte
	Cap    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetCap is a free log retrieval operation binding the contract event 0xe86b6d3313d3098f4c5f689c935de8fde876a597c185def2cedab85efedac686.
//
// Solidity: event SetCap(address indexed caller, bytes32 indexed id, uint256 cap)
func (_Morpho *MorphoFilterer) FilterSetCap(opts *bind.FilterOpts, caller []common.Address, id [][32]byte) (*MorphoSetCapIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Morpho.contract.FilterLogs(opts, "SetCap", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return &MorphoSetCapIterator{contract: _Morpho.contract, event: "SetCap", logs: logs, sub: sub}, nil
}

// WatchSetCap is a free log subscription operation binding the contract event 0xe86b6d3313d3098f4c5f689c935de8fde876a597c185def2cedab85efedac686.
//
// Solidity: event SetCap(address indexed caller, bytes32 indexed id, uint256 cap)
func (_Morpho *MorphoFilterer) WatchSetCap(opts *bind.WatchOpts, sink chan<- *MorphoSetCap, caller []common.Address, id [][32]byte) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Morpho.contract.WatchLogs(opts, "SetCap", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoSetCap)
				if err := _Morpho.contract.UnpackLog(event, "SetCap", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetCap is a log parse operation binding the contract event 0xe86b6d3313d3098f4c5f689c935de8fde876a597c185def2cedab85efedac686.
//
// Solidity: event SetCap(address indexed caller, bytes32 indexed id, uint256 cap)
func (_Morpho *MorphoFilterer) ParseSetCap(log types.Log) (*MorphoSetCap, error) {
	event := new(MorphoSetCap)
	if err := _Morpho.contract.UnpackLog(event, "SetCap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoSetCuratorIterator is returned from FilterSetCurator and is used to iterate over the raw logs and unpacked data for SetCurator events raised by the Morpho contract.
type MorphoSetCuratorIterator struct {
	Event *MorphoSetCurator // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoSetCuratorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoSetCurator)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoSetCurator)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoSetCuratorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoSetCuratorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoSetCurator represents a SetCurator event raised by the Morpho contract.
type MorphoSetCurator struct {
	NewCurator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetCurator is a free log retrieval operation binding the contract event 0xbd0a63c12948fbc9194a5839019f99c9d71db924e5c70018265bc778b8f1a506.
//
// Solidity: event SetCurator(address indexed newCurator)
func (_Morpho *MorphoFilterer) FilterSetCurator(opts *bind.FilterOpts, newCurator []common.Address) (*MorphoSetCuratorIterator, error) {

	var newCuratorRule []interface{}
	for _, newCuratorItem := range newCurator {
		newCuratorRule = append(newCuratorRule, newCuratorItem)
	}

	logs, sub, err := _Morpho.contract.FilterLogs(opts, "SetCurator", newCuratorRule)
	if err != nil {
		return nil, err
	}
	return &MorphoSetCuratorIterator{contract: _Morpho.contract, event: "SetCurator", logs: logs, sub: sub}, nil
}

// WatchSetCurator is a free log subscription operation binding the contract event 0xbd0a63c12948fbc9194a5839019f99c9d71db924e5c70018265bc778b8f1a506.
//
// Solidity: event SetCurator(address indexed newCurator)
func (_Morpho *MorphoFilterer) WatchSetCurator(opts *bind.WatchOpts, sink chan<- *MorphoSetCurator, newCurator []common.Address) (event.Subscription, error) {

	var newCuratorRule []interface{}
	for _, newCuratorItem := range newCurator {
		newCuratorRule = append(newCuratorRule, newCuratorItem)
	}

	logs, sub, err := _Morpho.contract.WatchLogs(opts, "SetCurator", newCuratorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoSetCurator)
				if err := _Morpho.contract.UnpackLog(event, "SetCurator", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetCurator is a log parse operation binding the contract event 0xbd0a63c12948fbc9194a5839019f99c9d71db924e5c70018265bc778b8f1a506.
//
// Solidity: event SetCurator(address indexed newCurator)
func (_Morpho *MorphoFilterer) ParseSetCurator(log types.Log) (*MorphoSetCurator, error) {
	event := new(MorphoSetCurator)
	if err := _Morpho.contract.UnpackLog(event, "SetCurator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoSetFeeIterator is returned from FilterSetFee and is used to iterate over the raw logs and unpacked data for SetFee events raised by the Morpho contract.
type MorphoSetFeeIterator struct {
	Event *MorphoSetFee // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoSetFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoSetFee)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoSetFee)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoSetFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoSetFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoSetFee represents a SetFee event raised by the Morpho contract.
type MorphoSetFee struct {
	Caller common.Address
	NewFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetFee is a free log retrieval operation binding the contract event 0x01fe2943baee27f47add82886c2200f910c749c461c9b63c5fe83901a53bdb49.
//
// Solidity: event SetFee(address indexed caller, uint256 newFee)
func (_Morpho *MorphoFilterer) FilterSetFee(opts *bind.FilterOpts, caller []common.Address) (*MorphoSetFeeIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Morpho.contract.FilterLogs(opts, "SetFee", callerRule)
	if err != nil {
		return nil, err
	}
	return &MorphoSetFeeIterator{contract: _Morpho.contract, event: "SetFee", logs: logs, sub: sub}, nil
}

// WatchSetFee is a free log subscription operation binding the contract event 0x01fe2943baee27f47add82886c2200f910c749c461c9b63c5fe83901a53bdb49.
//
// Solidity: event SetFee(address indexed caller, uint256 newFee)
func (_Morpho *MorphoFilterer) WatchSetFee(opts *bind.WatchOpts, sink chan<- *MorphoSetFee, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Morpho.contract.WatchLogs(opts, "SetFee", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoSetFee)
				if err := _Morpho.contract.UnpackLog(event, "SetFee", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetFee is a log parse operation binding the contract event 0x01fe2943baee27f47add82886c2200f910c749c461c9b63c5fe83901a53bdb49.
//
// Solidity: event SetFee(address indexed caller, uint256 newFee)
func (_Morpho *MorphoFilterer) ParseSetFee(log types.Log) (*MorphoSetFee, error) {
	event := new(MorphoSetFee)
	if err := _Morpho.contract.UnpackLog(event, "SetFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoSetFeeRecipientIterator is returned from FilterSetFeeRecipient and is used to iterate over the raw logs and unpacked data for SetFeeRecipient events raised by the Morpho contract.
type MorphoSetFeeRecipientIterator struct {
	Event *MorphoSetFeeRecipient // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoSetFeeRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoSetFeeRecipient)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoSetFeeRecipient)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoSetFeeRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoSetFeeRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoSetFeeRecipient represents a SetFeeRecipient event raised by the Morpho contract.
type MorphoSetFeeRecipient struct {
	NewFeeRecipient common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetFeeRecipient is a free log retrieval operation binding the contract event 0x2e979f80fe4d43055c584cf4a8467c55875ea36728fc37176c05acd784eb7a73.
//
// Solidity: event SetFeeRecipient(address indexed newFeeRecipient)
func (_Morpho *MorphoFilterer) FilterSetFeeRecipient(opts *bind.FilterOpts, newFeeRecipient []common.Address) (*MorphoSetFeeRecipientIterator, error) {

	var newFeeRecipientRule []interface{}
	for _, newFeeRecipientItem := range newFeeRecipient {
		newFeeRecipientRule = append(newFeeRecipientRule, newFeeRecipientItem)
	}

	logs, sub, err := _Morpho.contract.FilterLogs(opts, "SetFeeRecipient", newFeeRecipientRule)
	if err != nil {
		return nil, err
	}
	return &MorphoSetFeeRecipientIterator{contract: _Morpho.contract, event: "SetFeeRecipient", logs: logs, sub: sub}, nil
}

// WatchSetFeeRecipient is a free log subscription operation binding the contract event 0x2e979f80fe4d43055c584cf4a8467c55875ea36728fc37176c05acd784eb7a73.
//
// Solidity: event SetFeeRecipient(address indexed newFeeRecipient)
func (_Morpho *MorphoFilterer) WatchSetFeeRecipient(opts *bind.WatchOpts, sink chan<- *MorphoSetFeeRecipient, newFeeRecipient []common.Address) (event.Subscription, error) {

	var newFeeRecipientRule []interface{}
	for _, newFeeRecipientItem := range newFeeRecipient {
		newFeeRecipientRule = append(newFeeRecipientRule, newFeeRecipientItem)
	}

	logs, sub, err := _Morpho.contract.WatchLogs(opts, "SetFeeRecipient", newFeeRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoSetFeeRecipient)
				if err := _Morpho.contract.UnpackLog(event, "SetFeeRecipient", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetFeeRecipient is a log parse operation binding the contract event 0x2e979f80fe4d43055c584cf4a8467c55875ea36728fc37176c05acd784eb7a73.
//
// Solidity: event SetFeeRecipient(address indexed newFeeRecipient)
func (_Morpho *MorphoFilterer) ParseSetFeeRecipient(log types.Log) (*MorphoSetFeeRecipient, error) {
	event := new(MorphoSetFeeRecipient)
	if err := _Morpho.contract.UnpackLog(event, "SetFeeRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoSetGuardianIterator is returned from FilterSetGuardian and is used to iterate over the raw logs and unpacked data for SetGuardian events raised by the Morpho contract.
type MorphoSetGuardianIterator struct {
	Event *MorphoSetGuardian // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoSetGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoSetGuardian)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoSetGuardian)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoSetGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoSetGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoSetGuardian represents a SetGuardian event raised by the Morpho contract.
type MorphoSetGuardian struct {
	Caller   common.Address
	Guardian common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetGuardian is a free log retrieval operation binding the contract event 0xcb11cc8aade2f5a556749d1b2380d108a16fac3431e6a5d5ce12ef9de0bd76e3.
//
// Solidity: event SetGuardian(address indexed caller, address indexed guardian)
func (_Morpho *MorphoFilterer) FilterSetGuardian(opts *bind.FilterOpts, caller []common.Address, guardian []common.Address) (*MorphoSetGuardianIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _Morpho.contract.FilterLogs(opts, "SetGuardian", callerRule, guardianRule)
	if err != nil {
		return nil, err
	}
	return &MorphoSetGuardianIterator{contract: _Morpho.contract, event: "SetGuardian", logs: logs, sub: sub}, nil
}

// WatchSetGuardian is a free log subscription operation binding the contract event 0xcb11cc8aade2f5a556749d1b2380d108a16fac3431e6a5d5ce12ef9de0bd76e3.
//
// Solidity: event SetGuardian(address indexed caller, address indexed guardian)
func (_Morpho *MorphoFilterer) WatchSetGuardian(opts *bind.WatchOpts, sink chan<- *MorphoSetGuardian, caller []common.Address, guardian []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _Morpho.contract.WatchLogs(opts, "SetGuardian", callerRule, guardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoSetGuardian)
				if err := _Morpho.contract.UnpackLog(event, "SetGuardian", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetGuardian is a log parse operation binding the contract event 0xcb11cc8aade2f5a556749d1b2380d108a16fac3431e6a5d5ce12ef9de0bd76e3.
//
// Solidity: event SetGuardian(address indexed caller, address indexed guardian)
func (_Morpho *MorphoFilterer) ParseSetGuardian(log types.Log) (*MorphoSetGuardian, error) {
	event := new(MorphoSetGuardian)
	if err := _Morpho.contract.UnpackLog(event, "SetGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoSetIsAllocatorIterator is returned from FilterSetIsAllocator and is used to iterate over the raw logs and unpacked data for SetIsAllocator events raised by the Morpho contract.
type MorphoSetIsAllocatorIterator struct {
	Event *MorphoSetIsAllocator // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoSetIsAllocatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoSetIsAllocator)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoSetIsAllocator)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoSetIsAllocatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoSetIsAllocatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoSetIsAllocator represents a SetIsAllocator event raised by the Morpho contract.
type MorphoSetIsAllocator struct {
	Allocator   common.Address
	IsAllocator bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetIsAllocator is a free log retrieval operation binding the contract event 0x74dc60cbc81a9472d04ad1d20e151d369c41104d655ed3f2f3091166a502cd8d.
//
// Solidity: event SetIsAllocator(address indexed allocator, bool isAllocator)
func (_Morpho *MorphoFilterer) FilterSetIsAllocator(opts *bind.FilterOpts, allocator []common.Address) (*MorphoSetIsAllocatorIterator, error) {

	var allocatorRule []interface{}
	for _, allocatorItem := range allocator {
		allocatorRule = append(allocatorRule, allocatorItem)
	}

	logs, sub, err := _Morpho.contract.FilterLogs(opts, "SetIsAllocator", allocatorRule)
	if err != nil {
		return nil, err
	}
	return &MorphoSetIsAllocatorIterator{contract: _Morpho.contract, event: "SetIsAllocator", logs: logs, sub: sub}, nil
}

// WatchSetIsAllocator is a free log subscription operation binding the contract event 0x74dc60cbc81a9472d04ad1d20e151d369c41104d655ed3f2f3091166a502cd8d.
//
// Solidity: event SetIsAllocator(address indexed allocator, bool isAllocator)
func (_Morpho *MorphoFilterer) WatchSetIsAllocator(opts *bind.WatchOpts, sink chan<- *MorphoSetIsAllocator, allocator []common.Address) (event.Subscription, error) {

	var allocatorRule []interface{}
	for _, allocatorItem := range allocator {
		allocatorRule = append(allocatorRule, allocatorItem)
	}

	logs, sub, err := _Morpho.contract.WatchLogs(opts, "SetIsAllocator", allocatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoSetIsAllocator)
				if err := _Morpho.contract.UnpackLog(event, "SetIsAllocator", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetIsAllocator is a log parse operation binding the contract event 0x74dc60cbc81a9472d04ad1d20e151d369c41104d655ed3f2f3091166a502cd8d.
//
// Solidity: event SetIsAllocator(address indexed allocator, bool isAllocator)
func (_Morpho *MorphoFilterer) ParseSetIsAllocator(log types.Log) (*MorphoSetIsAllocator, error) {
	event := new(MorphoSetIsAllocator)
	if err := _Morpho.contract.UnpackLog(event, "SetIsAllocator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoSetSkimRecipientIterator is returned from FilterSetSkimRecipient and is used to iterate over the raw logs and unpacked data for SetSkimRecipient events raised by the Morpho contract.
type MorphoSetSkimRecipientIterator struct {
	Event *MorphoSetSkimRecipient // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoSetSkimRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoSetSkimRecipient)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoSetSkimRecipient)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoSetSkimRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoSetSkimRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoSetSkimRecipient represents a SetSkimRecipient event raised by the Morpho contract.
type MorphoSetSkimRecipient struct {
	NewSkimRecipient common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetSkimRecipient is a free log retrieval operation binding the contract event 0x2e7908865670e21b9779422cadf5f1cba271a62bb95c71eaaf615c0a1c48ebee.
//
// Solidity: event SetSkimRecipient(address indexed newSkimRecipient)
func (_Morpho *MorphoFilterer) FilterSetSkimRecipient(opts *bind.FilterOpts, newSkimRecipient []common.Address) (*MorphoSetSkimRecipientIterator, error) {

	var newSkimRecipientRule []interface{}
	for _, newSkimRecipientItem := range newSkimRecipient {
		newSkimRecipientRule = append(newSkimRecipientRule, newSkimRecipientItem)
	}

	logs, sub, err := _Morpho.contract.FilterLogs(opts, "SetSkimRecipient", newSkimRecipientRule)
	if err != nil {
		return nil, err
	}
	return &MorphoSetSkimRecipientIterator{contract: _Morpho.contract, event: "SetSkimRecipient", logs: logs, sub: sub}, nil
}

// WatchSetSkimRecipient is a free log subscription operation binding the contract event 0x2e7908865670e21b9779422cadf5f1cba271a62bb95c71eaaf615c0a1c48ebee.
//
// Solidity: event SetSkimRecipient(address indexed newSkimRecipient)
func (_Morpho *MorphoFilterer) WatchSetSkimRecipient(opts *bind.WatchOpts, sink chan<- *MorphoSetSkimRecipient, newSkimRecipient []common.Address) (event.Subscription, error) {

	var newSkimRecipientRule []interface{}
	for _, newSkimRecipientItem := range newSkimRecipient {
		newSkimRecipientRule = append(newSkimRecipientRule, newSkimRecipientItem)
	}

	logs, sub, err := _Morpho.contract.WatchLogs(opts, "SetSkimRecipient", newSkimRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoSetSkimRecipient)
				if err := _Morpho.contract.UnpackLog(event, "SetSkimRecipient", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetSkimRecipient is a log parse operation binding the contract event 0x2e7908865670e21b9779422cadf5f1cba271a62bb95c71eaaf615c0a1c48ebee.
//
// Solidity: event SetSkimRecipient(address indexed newSkimRecipient)
func (_Morpho *MorphoFilterer) ParseSetSkimRecipient(log types.Log) (*MorphoSetSkimRecipient, error) {
	event := new(MorphoSetSkimRecipient)
	if err := _Morpho.contract.UnpackLog(event, "SetSkimRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoSetSupplyQueueIterator is returned from FilterSetSupplyQueue and is used to iterate over the raw logs and unpacked data for SetSupplyQueue events raised by the Morpho contract.
type MorphoSetSupplyQueueIterator struct {
	Event *MorphoSetSupplyQueue // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoSetSupplyQueueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoSetSupplyQueue)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoSetSupplyQueue)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoSetSupplyQueueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoSetSupplyQueueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoSetSupplyQueue represents a SetSupplyQueue event raised by the Morpho contract.
type MorphoSetSupplyQueue struct {
	Caller         common.Address
	NewSupplyQueue [][32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSetSupplyQueue is a free log retrieval operation binding the contract event 0x6ce31538fc7fba95714ddc8a275a09252b4b1fb8f33d2550aa58a5f62ad934de.
//
// Solidity: event SetSupplyQueue(address indexed caller, bytes32[] newSupplyQueue)
func (_Morpho *MorphoFilterer) FilterSetSupplyQueue(opts *bind.FilterOpts, caller []common.Address) (*MorphoSetSupplyQueueIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Morpho.contract.FilterLogs(opts, "SetSupplyQueue", callerRule)
	if err != nil {
		return nil, err
	}
	return &MorphoSetSupplyQueueIterator{contract: _Morpho.contract, event: "SetSupplyQueue", logs: logs, sub: sub}, nil
}

// WatchSetSupplyQueue is a free log subscription operation binding the contract event 0x6ce31538fc7fba95714ddc8a275a09252b4b1fb8f33d2550aa58a5f62ad934de.
//
// Solidity: event SetSupplyQueue(address indexed caller, bytes32[] newSupplyQueue)
func (_Morpho *MorphoFilterer) WatchSetSupplyQueue(opts *bind.WatchOpts, sink chan<- *MorphoSetSupplyQueue, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Morpho.contract.WatchLogs(opts, "SetSupplyQueue", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoSetSupplyQueue)
				if err := _Morpho.contract.UnpackLog(event, "SetSupplyQueue", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetSupplyQueue is a log parse operation binding the contract event 0x6ce31538fc7fba95714ddc8a275a09252b4b1fb8f33d2550aa58a5f62ad934de.
//
// Solidity: event SetSupplyQueue(address indexed caller, bytes32[] newSupplyQueue)
func (_Morpho *MorphoFilterer) ParseSetSupplyQueue(log types.Log) (*MorphoSetSupplyQueue, error) {
	event := new(MorphoSetSupplyQueue)
	if err := _Morpho.contract.UnpackLog(event, "SetSupplyQueue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoSetTimelockIterator is returned from FilterSetTimelock and is used to iterate over the raw logs and unpacked data for SetTimelock events raised by the Morpho contract.
type MorphoSetTimelockIterator struct {
	Event *MorphoSetTimelock // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoSetTimelockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoSetTimelock)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoSetTimelock)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoSetTimelockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoSetTimelockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoSetTimelock represents a SetTimelock event raised by the Morpho contract.
type MorphoSetTimelock struct {
	Caller      common.Address
	NewTimelock *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetTimelock is a free log retrieval operation binding the contract event 0xd28e9b90ee9b37c5936ff84392d71f29ff18117d7e76bcee60615262a90a3f75.
//
// Solidity: event SetTimelock(address indexed caller, uint256 newTimelock)
func (_Morpho *MorphoFilterer) FilterSetTimelock(opts *bind.FilterOpts, caller []common.Address) (*MorphoSetTimelockIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Morpho.contract.FilterLogs(opts, "SetTimelock", callerRule)
	if err != nil {
		return nil, err
	}
	return &MorphoSetTimelockIterator{contract: _Morpho.contract, event: "SetTimelock", logs: logs, sub: sub}, nil
}

// WatchSetTimelock is a free log subscription operation binding the contract event 0xd28e9b90ee9b37c5936ff84392d71f29ff18117d7e76bcee60615262a90a3f75.
//
// Solidity: event SetTimelock(address indexed caller, uint256 newTimelock)
func (_Morpho *MorphoFilterer) WatchSetTimelock(opts *bind.WatchOpts, sink chan<- *MorphoSetTimelock, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Morpho.contract.WatchLogs(opts, "SetTimelock", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoSetTimelock)
				if err := _Morpho.contract.UnpackLog(event, "SetTimelock", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetTimelock is a log parse operation binding the contract event 0xd28e9b90ee9b37c5936ff84392d71f29ff18117d7e76bcee60615262a90a3f75.
//
// Solidity: event SetTimelock(address indexed caller, uint256 newTimelock)
func (_Morpho *MorphoFilterer) ParseSetTimelock(log types.Log) (*MorphoSetTimelock, error) {
	event := new(MorphoSetTimelock)
	if err := _Morpho.contract.UnpackLog(event, "SetTimelock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoSetWithdrawQueueIterator is returned from FilterSetWithdrawQueue and is used to iterate over the raw logs and unpacked data for SetWithdrawQueue events raised by the Morpho contract.
type MorphoSetWithdrawQueueIterator struct {
	Event *MorphoSetWithdrawQueue // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoSetWithdrawQueueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoSetWithdrawQueue)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoSetWithdrawQueue)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoSetWithdrawQueueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoSetWithdrawQueueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoSetWithdrawQueue represents a SetWithdrawQueue event raised by the Morpho contract.
type MorphoSetWithdrawQueue struct {
	Caller           common.Address
	NewWithdrawQueue [][32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetWithdrawQueue is a free log retrieval operation binding the contract event 0xe0c2db6b54586be6d7d49943139fccf0dd315ba63e55364a76c73cd8fdba724d.
//
// Solidity: event SetWithdrawQueue(address indexed caller, bytes32[] newWithdrawQueue)
func (_Morpho *MorphoFilterer) FilterSetWithdrawQueue(opts *bind.FilterOpts, caller []common.Address) (*MorphoSetWithdrawQueueIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Morpho.contract.FilterLogs(opts, "SetWithdrawQueue", callerRule)
	if err != nil {
		return nil, err
	}
	return &MorphoSetWithdrawQueueIterator{contract: _Morpho.contract, event: "SetWithdrawQueue", logs: logs, sub: sub}, nil
}

// WatchSetWithdrawQueue is a free log subscription operation binding the contract event 0xe0c2db6b54586be6d7d49943139fccf0dd315ba63e55364a76c73cd8fdba724d.
//
// Solidity: event SetWithdrawQueue(address indexed caller, bytes32[] newWithdrawQueue)
func (_Morpho *MorphoFilterer) WatchSetWithdrawQueue(opts *bind.WatchOpts, sink chan<- *MorphoSetWithdrawQueue, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Morpho.contract.WatchLogs(opts, "SetWithdrawQueue", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoSetWithdrawQueue)
				if err := _Morpho.contract.UnpackLog(event, "SetWithdrawQueue", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetWithdrawQueue is a log parse operation binding the contract event 0xe0c2db6b54586be6d7d49943139fccf0dd315ba63e55364a76c73cd8fdba724d.
//
// Solidity: event SetWithdrawQueue(address indexed caller, bytes32[] newWithdrawQueue)
func (_Morpho *MorphoFilterer) ParseSetWithdrawQueue(log types.Log) (*MorphoSetWithdrawQueue, error) {
	event := new(MorphoSetWithdrawQueue)
	if err := _Morpho.contract.UnpackLog(event, "SetWithdrawQueue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoSkimIterator is returned from FilterSkim and is used to iterate over the raw logs and unpacked data for Skim events raised by the Morpho contract.
type MorphoSkimIterator struct {
	Event *MorphoSkim // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoSkimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoSkim)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoSkim)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoSkimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoSkimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoSkim represents a Skim event raised by the Morpho contract.
type MorphoSkim struct {
	Caller common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSkim is a free log retrieval operation binding the contract event 0x2ae72b44f59d038340fca5739135a1d51fc5ab720bb02d983e4c5ff4119ca7b8.
//
// Solidity: event Skim(address indexed caller, address indexed token, uint256 amount)
func (_Morpho *MorphoFilterer) FilterSkim(opts *bind.FilterOpts, caller []common.Address, token []common.Address) (*MorphoSkimIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Morpho.contract.FilterLogs(opts, "Skim", callerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &MorphoSkimIterator{contract: _Morpho.contract, event: "Skim", logs: logs, sub: sub}, nil
}

// WatchSkim is a free log subscription operation binding the contract event 0x2ae72b44f59d038340fca5739135a1d51fc5ab720bb02d983e4c5ff4119ca7b8.
//
// Solidity: event Skim(address indexed caller, address indexed token, uint256 amount)
func (_Morpho *MorphoFilterer) WatchSkim(opts *bind.WatchOpts, sink chan<- *MorphoSkim, caller []common.Address, token []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Morpho.contract.WatchLogs(opts, "Skim", callerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoSkim)
				if err := _Morpho.contract.UnpackLog(event, "Skim", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSkim is a log parse operation binding the contract event 0x2ae72b44f59d038340fca5739135a1d51fc5ab720bb02d983e4c5ff4119ca7b8.
//
// Solidity: event Skim(address indexed caller, address indexed token, uint256 amount)
func (_Morpho *MorphoFilterer) ParseSkim(log types.Log) (*MorphoSkim, error) {
	event := new(MorphoSkim)
	if err := _Morpho.contract.UnpackLog(event, "Skim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoSubmitCapIterator is returned from FilterSubmitCap and is used to iterate over the raw logs and unpacked data for SubmitCap events raised by the Morpho contract.
type MorphoSubmitCapIterator struct {
	Event *MorphoSubmitCap // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoSubmitCapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoSubmitCap)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoSubmitCap)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoSubmitCapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoSubmitCapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoSubmitCap represents a SubmitCap event raised by the Morpho contract.
type MorphoSubmitCap struct {
	Caller common.Address
	Id     [32]byte
	Cap    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSubmitCap is a free log retrieval operation binding the contract event 0xe851bb5856808a50efd748be463b8f35bcfb5ec74c5bfde776fe0a4d2a26db27.
//
// Solidity: event SubmitCap(address indexed caller, bytes32 indexed id, uint256 cap)
func (_Morpho *MorphoFilterer) FilterSubmitCap(opts *bind.FilterOpts, caller []common.Address, id [][32]byte) (*MorphoSubmitCapIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Morpho.contract.FilterLogs(opts, "SubmitCap", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return &MorphoSubmitCapIterator{contract: _Morpho.contract, event: "SubmitCap", logs: logs, sub: sub}, nil
}

// WatchSubmitCap is a free log subscription operation binding the contract event 0xe851bb5856808a50efd748be463b8f35bcfb5ec74c5bfde776fe0a4d2a26db27.
//
// Solidity: event SubmitCap(address indexed caller, bytes32 indexed id, uint256 cap)
func (_Morpho *MorphoFilterer) WatchSubmitCap(opts *bind.WatchOpts, sink chan<- *MorphoSubmitCap, caller []common.Address, id [][32]byte) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Morpho.contract.WatchLogs(opts, "SubmitCap", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoSubmitCap)
				if err := _Morpho.contract.UnpackLog(event, "SubmitCap", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSubmitCap is a log parse operation binding the contract event 0xe851bb5856808a50efd748be463b8f35bcfb5ec74c5bfde776fe0a4d2a26db27.
//
// Solidity: event SubmitCap(address indexed caller, bytes32 indexed id, uint256 cap)
func (_Morpho *MorphoFilterer) ParseSubmitCap(log types.Log) (*MorphoSubmitCap, error) {
	event := new(MorphoSubmitCap)
	if err := _Morpho.contract.UnpackLog(event, "SubmitCap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoSubmitGuardianIterator is returned from FilterSubmitGuardian and is used to iterate over the raw logs and unpacked data for SubmitGuardian events raised by the Morpho contract.
type MorphoSubmitGuardianIterator struct {
	Event *MorphoSubmitGuardian // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoSubmitGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoSubmitGuardian)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoSubmitGuardian)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoSubmitGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoSubmitGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoSubmitGuardian represents a SubmitGuardian event raised by the Morpho contract.
type MorphoSubmitGuardian struct {
	NewGuardian common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSubmitGuardian is a free log retrieval operation binding the contract event 0x7633313af54753bce8a149927263b1a55eba857ba4ef1d13c6aee25d384d3c4b.
//
// Solidity: event SubmitGuardian(address indexed newGuardian)
func (_Morpho *MorphoFilterer) FilterSubmitGuardian(opts *bind.FilterOpts, newGuardian []common.Address) (*MorphoSubmitGuardianIterator, error) {

	var newGuardianRule []interface{}
	for _, newGuardianItem := range newGuardian {
		newGuardianRule = append(newGuardianRule, newGuardianItem)
	}

	logs, sub, err := _Morpho.contract.FilterLogs(opts, "SubmitGuardian", newGuardianRule)
	if err != nil {
		return nil, err
	}
	return &MorphoSubmitGuardianIterator{contract: _Morpho.contract, event: "SubmitGuardian", logs: logs, sub: sub}, nil
}

// WatchSubmitGuardian is a free log subscription operation binding the contract event 0x7633313af54753bce8a149927263b1a55eba857ba4ef1d13c6aee25d384d3c4b.
//
// Solidity: event SubmitGuardian(address indexed newGuardian)
func (_Morpho *MorphoFilterer) WatchSubmitGuardian(opts *bind.WatchOpts, sink chan<- *MorphoSubmitGuardian, newGuardian []common.Address) (event.Subscription, error) {

	var newGuardianRule []interface{}
	for _, newGuardianItem := range newGuardian {
		newGuardianRule = append(newGuardianRule, newGuardianItem)
	}

	logs, sub, err := _Morpho.contract.WatchLogs(opts, "SubmitGuardian", newGuardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoSubmitGuardian)
				if err := _Morpho.contract.UnpackLog(event, "SubmitGuardian", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSubmitGuardian is a log parse operation binding the contract event 0x7633313af54753bce8a149927263b1a55eba857ba4ef1d13c6aee25d384d3c4b.
//
// Solidity: event SubmitGuardian(address indexed newGuardian)
func (_Morpho *MorphoFilterer) ParseSubmitGuardian(log types.Log) (*MorphoSubmitGuardian, error) {
	event := new(MorphoSubmitGuardian)
	if err := _Morpho.contract.UnpackLog(event, "SubmitGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoSubmitMarketRemovalIterator is returned from FilterSubmitMarketRemoval and is used to iterate over the raw logs and unpacked data for SubmitMarketRemoval events raised by the Morpho contract.
type MorphoSubmitMarketRemovalIterator struct {
	Event *MorphoSubmitMarketRemoval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoSubmitMarketRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoSubmitMarketRemoval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoSubmitMarketRemoval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoSubmitMarketRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoSubmitMarketRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoSubmitMarketRemoval represents a SubmitMarketRemoval event raised by the Morpho contract.
type MorphoSubmitMarketRemoval struct {
	Caller common.Address
	Id     [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSubmitMarketRemoval is a free log retrieval operation binding the contract event 0x3240fc70754c5a2b4dab10bf7081a00024bfc8491581ee3d355360ec0dd91f16.
//
// Solidity: event SubmitMarketRemoval(address indexed caller, bytes32 indexed id)
func (_Morpho *MorphoFilterer) FilterSubmitMarketRemoval(opts *bind.FilterOpts, caller []common.Address, id [][32]byte) (*MorphoSubmitMarketRemovalIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Morpho.contract.FilterLogs(opts, "SubmitMarketRemoval", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return &MorphoSubmitMarketRemovalIterator{contract: _Morpho.contract, event: "SubmitMarketRemoval", logs: logs, sub: sub}, nil
}

// WatchSubmitMarketRemoval is a free log subscription operation binding the contract event 0x3240fc70754c5a2b4dab10bf7081a00024bfc8491581ee3d355360ec0dd91f16.
//
// Solidity: event SubmitMarketRemoval(address indexed caller, bytes32 indexed id)
func (_Morpho *MorphoFilterer) WatchSubmitMarketRemoval(opts *bind.WatchOpts, sink chan<- *MorphoSubmitMarketRemoval, caller []common.Address, id [][32]byte) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Morpho.contract.WatchLogs(opts, "SubmitMarketRemoval", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoSubmitMarketRemoval)
				if err := _Morpho.contract.UnpackLog(event, "SubmitMarketRemoval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSubmitMarketRemoval is a log parse operation binding the contract event 0x3240fc70754c5a2b4dab10bf7081a00024bfc8491581ee3d355360ec0dd91f16.
//
// Solidity: event SubmitMarketRemoval(address indexed caller, bytes32 indexed id)
func (_Morpho *MorphoFilterer) ParseSubmitMarketRemoval(log types.Log) (*MorphoSubmitMarketRemoval, error) {
	event := new(MorphoSubmitMarketRemoval)
	if err := _Morpho.contract.UnpackLog(event, "SubmitMarketRemoval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoSubmitTimelockIterator is returned from FilterSubmitTimelock and is used to iterate over the raw logs and unpacked data for SubmitTimelock events raised by the Morpho contract.
type MorphoSubmitTimelockIterator struct {
	Event *MorphoSubmitTimelock // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoSubmitTimelockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoSubmitTimelock)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoSubmitTimelock)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoSubmitTimelockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoSubmitTimelockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoSubmitTimelock represents a SubmitTimelock event raised by the Morpho contract.
type MorphoSubmitTimelock struct {
	NewTimelock *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSubmitTimelock is a free log retrieval operation binding the contract event 0xb3aa0ade2442acf51d06713c2d1a5a3ec0373cce969d42b53f4689f97bccf380.
//
// Solidity: event SubmitTimelock(uint256 newTimelock)
func (_Morpho *MorphoFilterer) FilterSubmitTimelock(opts *bind.FilterOpts) (*MorphoSubmitTimelockIterator, error) {

	logs, sub, err := _Morpho.contract.FilterLogs(opts, "SubmitTimelock")
	if err != nil {
		return nil, err
	}
	return &MorphoSubmitTimelockIterator{contract: _Morpho.contract, event: "SubmitTimelock", logs: logs, sub: sub}, nil
}

// WatchSubmitTimelock is a free log subscription operation binding the contract event 0xb3aa0ade2442acf51d06713c2d1a5a3ec0373cce969d42b53f4689f97bccf380.
//
// Solidity: event SubmitTimelock(uint256 newTimelock)
func (_Morpho *MorphoFilterer) WatchSubmitTimelock(opts *bind.WatchOpts, sink chan<- *MorphoSubmitTimelock) (event.Subscription, error) {

	logs, sub, err := _Morpho.contract.WatchLogs(opts, "SubmitTimelock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoSubmitTimelock)
				if err := _Morpho.contract.UnpackLog(event, "SubmitTimelock", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSubmitTimelock is a log parse operation binding the contract event 0xb3aa0ade2442acf51d06713c2d1a5a3ec0373cce969d42b53f4689f97bccf380.
//
// Solidity: event SubmitTimelock(uint256 newTimelock)
func (_Morpho *MorphoFilterer) ParseSubmitTimelock(log types.Log) (*MorphoSubmitTimelock, error) {
	event := new(MorphoSubmitTimelock)
	if err := _Morpho.contract.UnpackLog(event, "SubmitTimelock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Morpho contract.
type MorphoTransferIterator struct {
	Event *MorphoTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoTransfer represents a Transfer event raised by the Morpho contract.
type MorphoTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Morpho *MorphoFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MorphoTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Morpho.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MorphoTransferIterator{contract: _Morpho.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Morpho *MorphoFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MorphoTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Morpho.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoTransfer)
				if err := _Morpho.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Morpho *MorphoFilterer) ParseTransfer(log types.Log) (*MorphoTransfer, error) {
	event := new(MorphoTransfer)
	if err := _Morpho.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoUpdateLastTotalAssetsIterator is returned from FilterUpdateLastTotalAssets and is used to iterate over the raw logs and unpacked data for UpdateLastTotalAssets events raised by the Morpho contract.
type MorphoUpdateLastTotalAssetsIterator struct {
	Event *MorphoUpdateLastTotalAssets // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoUpdateLastTotalAssetsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoUpdateLastTotalAssets)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoUpdateLastTotalAssets)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoUpdateLastTotalAssetsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoUpdateLastTotalAssetsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoUpdateLastTotalAssets represents a UpdateLastTotalAssets event raised by the Morpho contract.
type MorphoUpdateLastTotalAssets struct {
	UpdatedTotalAssets *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterUpdateLastTotalAssets is a free log retrieval operation binding the contract event 0x15c027cc4fd826d986cad358803439f7326d3aa4ed969ff90dbee4bc150f68e9.
//
// Solidity: event UpdateLastTotalAssets(uint256 updatedTotalAssets)
func (_Morpho *MorphoFilterer) FilterUpdateLastTotalAssets(opts *bind.FilterOpts) (*MorphoUpdateLastTotalAssetsIterator, error) {

	logs, sub, err := _Morpho.contract.FilterLogs(opts, "UpdateLastTotalAssets")
	if err != nil {
		return nil, err
	}
	return &MorphoUpdateLastTotalAssetsIterator{contract: _Morpho.contract, event: "UpdateLastTotalAssets", logs: logs, sub: sub}, nil
}

// WatchUpdateLastTotalAssets is a free log subscription operation binding the contract event 0x15c027cc4fd826d986cad358803439f7326d3aa4ed969ff90dbee4bc150f68e9.
//
// Solidity: event UpdateLastTotalAssets(uint256 updatedTotalAssets)
func (_Morpho *MorphoFilterer) WatchUpdateLastTotalAssets(opts *bind.WatchOpts, sink chan<- *MorphoUpdateLastTotalAssets) (event.Subscription, error) {

	logs, sub, err := _Morpho.contract.WatchLogs(opts, "UpdateLastTotalAssets")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoUpdateLastTotalAssets)
				if err := _Morpho.contract.UnpackLog(event, "UpdateLastTotalAssets", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateLastTotalAssets is a log parse operation binding the contract event 0x15c027cc4fd826d986cad358803439f7326d3aa4ed969ff90dbee4bc150f68e9.
//
// Solidity: event UpdateLastTotalAssets(uint256 updatedTotalAssets)
func (_Morpho *MorphoFilterer) ParseUpdateLastTotalAssets(log types.Log) (*MorphoUpdateLastTotalAssets, error) {
	event := new(MorphoUpdateLastTotalAssets)
	if err := _Morpho.contract.UnpackLog(event, "UpdateLastTotalAssets", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Morpho contract.
type MorphoWithdrawIterator struct {
	Event *MorphoWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoWithdraw represents a Withdraw event raised by the Morpho contract.
type MorphoWithdraw struct {
	Sender   common.Address
	Receiver common.Address
	Owner    common.Address
	Assets   *big.Int
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_Morpho *MorphoFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address, owner []common.Address) (*MorphoWithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Morpho.contract.FilterLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &MorphoWithdrawIterator{contract: _Morpho.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_Morpho *MorphoFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *MorphoWithdraw, sender []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Morpho.contract.WatchLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoWithdraw)
				if err := _Morpho.contract.UnpackLog(event, "Withdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdraw is a log parse operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_Morpho *MorphoFilterer) ParseWithdraw(log types.Log) (*MorphoWithdraw, error) {
	event := new(MorphoWithdraw)
	if err := _Morpho.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
