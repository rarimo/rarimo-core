package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	rarimotypes "github.com/rarimo/rarimo-core/x/rarimocore/types"
)

// AccountKeeper defines the account contract that must be fulfilled when
type AccountKeeper interface {
	GetModuleAddress(moduleName string) sdk.AccAddress
	GetModuleAddressAndPermissions(moduleName string) (addr sdk.AccAddress, permissions []string)
	GetModuleAccountAndPermissions(ctx sdk.Context, moduleName string) (types.ModuleAccountI, []string)
	GetModuleAccount(ctx sdk.Context, moduleName string) types.ModuleAccountI
	SetModuleAccount(ctx sdk.Context, macc types.ModuleAccountI)
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
	// Methods imported from bank should be defined here
}

// RarimoKeeper defines the expected interface of rarimo core module keeper
type RarimoKeeper interface {
	IterateVotes(ctx sdk.Context, operation string, f func(vote rarimotypes.Vote) (stop bool))

	GetIdentityDefaultTransfer(_ sdk.Context, msg *MsgCreateIdentityDefaultTransferOp) (*rarimotypes.IdentityDefaultTransfer, error)
	CreateIdentityDefaultTransferOperation(ctx sdk.Context, creator string, transfer *rarimotypes.IdentityDefaultTransfer) error

	GetIdentityGISTTransfer(_ sdk.Context, msg *MsgCreateIdentityGISTTransferOp) (*rarimotypes.IdentityGISTTransfer, error)
	CreateIdentityGISTTransferOperation(ctx sdk.Context, creator string, transfer *rarimotypes.IdentityGISTTransfer) error

	GetIdentityStateTransfer(_ sdk.Context, msg *MsgCreateIdentityStateTransferOp) (*rarimotypes.IdentityStateTransfer, error)
	CreateIdentityStateTransferOperation(ctx sdk.Context, creator string, transfer *rarimotypes.IdentityStateTransfer) error

	GetTransfer(ctx sdk.Context, msg *MsgCreateTransferOp) (*rarimotypes.Transfer, error)
	CreateTransferOperation(ctx sdk.Context, creator string, transfer *rarimotypes.Transfer, approved bool) error

	GetVote(ctx sdk.Context, index *rarimotypes.VoteIndex) (val rarimotypes.Vote, found bool)
	CreateVote(ctx sdk.Context, vote rarimotypes.Vote) (bool, error)

	GetOperation(ctx sdk.Context, index string) (val rarimotypes.Operation, found bool)
	ApproveOperation(ctx sdk.Context, op rarimotypes.Operation) error
	UnapproveOperation(ctx sdk.Context, op rarimotypes.Operation) error
}
