package types

import (
	cosmostypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/gogo/protobuf/proto"
)

func (n *Network) GetBridgeParams() *BridgeNetworkParams {
	for _, p := range n.Params {
		if p.Type == NetworkParamType_BRIDGE {
			params := new(BridgeNetworkParams)
			proto.Unmarshal(p.Details.Value, params)
			return params
		}
	}
	return nil
}

func (n *Network) GetIdentityParams() *IdentityNetworkParams {
	for _, p := range n.Params {
		if p.Type == NetworkParamType_IDENTITY {
			params := new(IdentityNetworkParams)
			proto.Unmarshal(p.Details.Value, params)
			return params
		}
	}
	return nil
}

func (n *Network) GetFeeParams() *FeeNetworkParams {
	for _, p := range n.Params {
		if p.Type == NetworkParamType_FEE {
			params := new(FeeNetworkParams)
			proto.Unmarshal(p.Details.Value, params)
			return params
		}
	}
	return nil
}

func (n *Network) SetBridgeParams(params *BridgeNetworkParams) {
	details, _ := cosmostypes.NewAnyWithValue(params)
	newParams := NetworkParams{
		Type:    NetworkParamType_BRIDGE,
		Details: details,
	}

	for i, p := range n.Params {
		if p.Type == NetworkParamType_BRIDGE {
			n.Params[i] = newParams
			return
		}
	}

	n.Params = append(n.Params, newParams)
}

func (n *Network) SetIdentityParams(params *IdentityNetworkParams) {
	details, _ := cosmostypes.NewAnyWithValue(params)
	newParams := NetworkParams{
		Type:    NetworkParamType_IDENTITY,
		Details: details,
	}

	for i, p := range n.Params {
		if p.Type == NetworkParamType_IDENTITY {
			n.Params[i] = newParams
			return
		}
	}

	n.Params = append(n.Params, newParams)
}

func (n *Network) SetFeeParams(params *FeeNetworkParams) {
	details, _ := cosmostypes.NewAnyWithValue(params)
	newParams := NetworkParams{
		Type:    NetworkParamType_FEE,
		Details: details,
	}

	for i, p := range n.Params {
		if p.Type == NetworkParamType_FEE {
			n.Params[i] = newParams
			return
		}
	}

	n.Params = append(n.Params, newParams)
}

func (p *FeeNetworkParams) GetFeeToken(contract string) *FeeToken {
	if p == nil {
		return nil
	}

	for _, t := range p.FeeTokens {
		if t.Contract == contract {
			return t
		}
	}

	return nil
}

func (p *FeeNetworkParams) SetFeeToken(token *FeeToken) {
	if p == nil || token == nil {
		return
	}

	for i, t := range p.FeeTokens {
		if t.Contract == token.Contract {
			p.FeeTokens[i] = token
			return
		}
	}

	p.FeeTokens = append(p.FeeTokens, token)
}

func (p *FeeNetworkParams) RemoveFeeToken(token *FeeToken) {
	if p == nil || token == nil {
		return
	}

	tokens := make([]*FeeToken, 0, len(p.FeeTokens))

	for _, t := range p.FeeTokens {
		if t.Contract != token.Contract {
			tokens = append(tokens, t)
		}
	}

	p.FeeTokens = tokens
}
