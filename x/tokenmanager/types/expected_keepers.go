package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}

type RarimocoreKeeper interface {
	CreateAddFeeTokenOperation(ctx sdk.Context, token FeeToken, chain string, nonce string) error
	CreateRemoveFeeTokenOperation(ctx sdk.Context, token FeeToken, chain string, nonce string) error
	CreateUpdateFeeTokenOperation(ctx sdk.Context, token FeeToken, chain string, nonce string) error
	CreateWithdrawFeeOperation(ctx sdk.Context, token FeeToken, chain string, receiver string, nonce string) error
	CreateContractUpgradeOperation(ctx sdk.Context, upgradeDetails *ContractUpgradeDetails) error
}
