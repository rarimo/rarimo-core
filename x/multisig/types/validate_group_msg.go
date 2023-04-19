package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type groupMsg interface {
	GetCreator() string
	GetMembers() []string
	GetThreshold() uint64
}

func validateGroupMessage(msg groupMsg) error {
	if _, err := sdk.AccAddressFromBech32(msg.GetCreator()); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	members := msg.GetMembers()
	index := make(map[string]struct{}, len(members))

	for _, member := range members {
		if _, err := sdk.AccAddressFromBech32(member); err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid member address (%s)", err)
		}
		if _, exists := index[member]; exists {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "duplicate member address: %s", member)
		}
		index[member] = struct{}{}
	}

	threshold := msg.GetThreshold()
	if threshold == 0 || threshold > uint64(len(members)) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "threshold must be greater than 0 and less or equal to the number of members")
	}

	return nil
}
