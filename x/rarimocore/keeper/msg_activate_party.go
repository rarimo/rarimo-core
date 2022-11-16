package keeper

import (
	"context"
	"crypto/elliptic"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
)

func (k msgServer) ActivateParty(goCtx context.Context, msg *types.MsgActivateParty) (*types.MsgActivatePartyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	defer k.disableFee(ctx.GasMeter().GasConsumed(), ctx.GasMeter())

	params := k.GetParams(ctx)
	for i, p := range params.Parties {
		if p.Account == msg.Creator {
			if p.Active {
				return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "only inactive parties can activate")
			}

			p.PubKey = msg.PubKey
			p.Active = true
			params.Parties[i] = p
			k.SetParams(ctx, params)

			if err := k.checkAllPartiesActive(ctx); err == nil {
				k.updateECDSAPubKey(ctx)
			}

			return &types.MsgActivatePartyResponse{}, nil
		}
	}

	return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "sender should be a party")
}

func (k msgServer) updateECDSAPubKey(ctx sdk.Context) {
	curve := crypto.S256()
	params := k.GetParams(ctx)

	x, y := big.NewInt(0), big.NewInt(0)
	for _, p := range params.Parties {
		kx, ky := elliptic.Unmarshal(curve, hexutil.MustDecode(p.PubKey))
		x, y = curve.Add(x, y, kx, ky)
	}

	params.KeyECDSA = hexutil.Encode(elliptic.Marshal(curve, x, y))
	k.SetParams(ctx, params)
}
