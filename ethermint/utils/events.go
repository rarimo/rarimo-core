package utils

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// unpackLog copy-pasted from logic in generated s-c bindings.
func UnpackLog(contractAbi abi.ABI, out interface{}, event string, log *ethtypes.Log) error {
	if log.Topics[0] != contractAbi.Events[event].ID {
		return fmt.Errorf("event signature mismatch")
	}

	if len(log.Data) > 0 {
		if err := contractAbi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return err
		}
	}
	var indexed abi.Arguments
	for _, arg := range contractAbi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	return abi.ParseTopics(out, indexed, log.Topics[1:])
}
