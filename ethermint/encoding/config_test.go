package encoding_test

import (
	"math/big"
	"testing"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
	encoding "github.com/rarimo/rarimo-core/ethermint/encoding"
	tests "github.com/rarimo/rarimo-core/ethermint/tests"
	"github.com/stretchr/testify/require"

	"github.com/rarimo/rarimo-core/app"
	evmtypes "github.com/rarimo/rarimo-core/x/evm/types"
)

func TestTxEncoding(t *testing.T) {
	addr, key := tests.NewAddrKey()
	signer := tests.NewSigner(key)

	msg := evmtypes.NewTxContract(big.NewInt(1), 1, big.NewInt(10), 100000, nil, big.NewInt(1), big.NewInt(1), []byte{}, nil)
	msg.From = addr.Hex()

	ethSigner := ethtypes.LatestSignerForChainID(big.NewInt(1))
	err := msg.Sign(ethSigner, signer)
	require.NoError(t, err)

	cfg := encoding.MakeConfig(app.ModuleBasics)

	_, err = cfg.TxConfig.TxEncoder()(msg)
	require.Error(t, err, "encoding failed")

	// FIXME: transaction hashing is hardcoded on Terndermint:
	// See https://github.com/tendermint/tendermint/issues/6539 for reference
	// txHash := msg.AsTransaction().Hash()
	// tmTx := tmtypes.Tx(bz)

	// require.Equal(t, txHash.Bytes(), tmTx.Hash())
}
