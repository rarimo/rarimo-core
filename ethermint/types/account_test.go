package types_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/suite"
	cryptocodec "gitlab.com/rarimo/rarimo-core/ethermint/crypto/codec"
	"gitlab.com/rarimo/rarimo-core/ethermint/crypto/ethsecp256k1"
	ethermintcodec "gitlab.com/rarimo/rarimo-core/ethermint/encoding/codec"
	types2 "gitlab.com/rarimo/rarimo-core/ethermint/types"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func init() {
	amino := codec.NewLegacyAmino()
	cryptocodec.RegisterCrypto(amino)
}

type AccountTestSuite struct {
	suite.Suite

	account *types2.EthAccount
	cdc     codec.JSONCodec
}

func (suite *AccountTestSuite) SetupTest() {
	privKey, err := ethsecp256k1.GenerateKey()
	suite.Require().NoError(err)
	pubKey := privKey.PubKey()
	addr := sdk.AccAddress(pubKey.Address())
	baseAcc := authtypes.NewBaseAccount(addr, pubKey, 10, 50)
	suite.account = &types2.EthAccount{
		BaseAccount: baseAcc,
		CodeHash:    common.Hash{}.String(),
	}

	interfaceRegistry := codectypes.NewInterfaceRegistry()
	ethermintcodec.RegisterInterfaces(interfaceRegistry)
	suite.cdc = codec.NewProtoCodec(interfaceRegistry)
}

func TestAccountTestSuite(t *testing.T) {
	suite.Run(t, new(AccountTestSuite))
}

func (suite *AccountTestSuite) TestAccountType() {
	suite.account.CodeHash = common.BytesToHash(crypto.Keccak256(nil)).Hex()
	suite.Require().Equal(types2.AccountTypeEOA, suite.account.Type())
	suite.account.CodeHash = common.BytesToHash(crypto.Keccak256([]byte{1, 2, 3})).Hex()
	suite.Require().Equal(types2.AccountTypeContract, suite.account.Type())
}
