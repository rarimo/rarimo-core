package types

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rarimo/rarimo-core/x/multisig/tx"
)

var _ types.UnpackInterfacesMessage = &Proposal{}

func (p *Proposal) GetMsgs() ([]sdk.Msg, error) {
	return tx.GetMsgs(p.Messages, "proposal")
}

func (p *Proposal) SetMsgs(msgs []sdk.Msg) error {
	anys, err := tx.SetMsgs(msgs)
	if err != nil {
		return err
	}
	p.Messages = anys
	return nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (p *Proposal) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	return tx.UnpackInterfaces(unpacker, p.Messages)
}
