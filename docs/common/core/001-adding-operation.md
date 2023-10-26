---
layout: default
title: Adding new operation
---

# Adding new operation

To provide TSS signature core uses operations and confirmation entities.
Operation entity represents some data to be signed by threshold signature producers.
Confirmation entity represents information about signature: indexes list, merkle root based on provided list and
signature.

Operation entity contains the following fields:

```protobuf
message Operation {
  string index = 1;
  opType operationType = 2;
  google.protobuf.Any details = 3;
  opStatus status = 4;
  string creator = 5;
  uint64 timestamp = 6;
}
```

- index is the unique string that should be deterministic created depending on operation data
- operation type defines the type of operation details
- details contain any necessary information about operation to provide signature for
- status defines the current status of operation (signed, approved, initialize, etc.)
- creator defines the creator of certain operation
- timestamp contains the unix block timestamp then operation was created (the timestamp should be received from the
  context)

----

**To add new operation developer should lead the following steps:**

1. Add operation data definition in the `proto/rarimocore`. Example:

   ___proto/ratimocore/op_fee_token_management.proto___

    ```protobuf
    enum FeeTokenManagementType {
      ADD_FEE_TOKEN = 0;
      REMOVE_FEE_TOKEN = 1;
      UPDATE_FEE_TOKEN = 2;
      WITHDRAW_FEE_TOKEN = 3;
    }

    message FeeTokenManagement {
      FeeTokenManagementType opType = 1;
      rarimo.rarimocore.tokenmanager.FeeToken token = 2 [(gogoproto.nullable) = false];
      string chain = 3;
      string receiver = 4;
    }
    ```

   Also, add new operation type in `proto/rarimocore/operation.proto`.

   ___proto/rarimocore/operation.proto___

    ```protobuf
    enum opType {
      TRANSFER = 0;
      CHANGE_PARTIES = 1;
      FEE_TOKEN_MANAGEMENT = 2;
    }
    ```

----

2. In `x/rarimocore/crypto/operation` define the operation content that should implement `merkle.Content` interface
   from `merkle "github.com/rarimo/go-merkle"`.
   Example:

   ___x/rarimocore/crypto/operation/op_fee_token_management.go___

    ```go
    package operation

    import (
      "bytes"

      eth "github.com/ethereum/go-ethereum/crypto"
      merkle "github.com/rarimo/go-merkle"
    )

    // FeeTokenManagementContent implements the Content interface provided by go-merkle and represents the content stored in the tree.
    type FeeTokenManagementContent struct {
      // Hash of the deposit tx info
      Origin        OriginData
      TargetNetwork string
      // Receiver address on target network
      Receiver []byte
      // Target bridge contract
      TargetContract []byte
      // Can contain any specific data for target chain to validate.
      Data ContentData
    }

    var _ merkle.Content = FeeTokenManagementContent{}

    func (f FeeTokenManagementContent) CalculateHash() []byte {
      return eth.Keccak256(f.Data, f.Origin[:], []byte(f.TargetNetwork), f.Receiver, f.TargetContract)
    }

    // Equals tests for equality of two Contents
    func (f FeeTokenManagementContent) Equals(other merkle.Content) bool {
      if oc, ok := other.(ChangePartiesContent); ok {
        return bytes.Equal(oc.CalculateHash(), f.CalculateHash())
      }

      return false
    }
    ```

----

3. In `x/rarimocore/crypto/pkg` define the following methods: `Get{op name}` and `Get{op name} content`.

   Example:
   ___x/rarimocore/crypto/operation/op_fee_token_management.go___

    ```go
    package operation

    func GetFeeTokenManagement(operation types.Operation) (*types.FeeTokenManagement, error) {
      if operation.OperationType == types.OpType_FEE_TOKEN_MANAGEMENT {
        op := new(types.FeeTokenManagement)
        return op, proto.Unmarshal(operation.Details.Value, op)
      }
      return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "invalid operation type")
    }

    func GetFeeTokenManagementContent(strOrigin string, params tokentypes.NetworkParams, op *types.FeeTokenManagement) (*operation.FeeTokenManagementContent, error) {
      return &operation.FeeTokenManagementContent{
        Origin:         origin.NewStringOriginBuilder().SetString(strOrigin).Build().GetOrigin(),
        TargetNetwork:  params.Name,
        Receiver:       hexutil.MustDecode(op.Receiver),
        TargetContract: hexutil.MustDecode(params.Contract),
        Data:           data.NewFeeTokenDataBuilder().SetOpType(op.OpType).SetAmount(op.Token.Amount).SetAmount(op.Token.Amount).Build().GetContent(),
      }, nil
    }

    ```

**Tips**: explore the `x/rarimocore/crypto/operation/data` and `x/rarimocore/crypto/operation/origin` packages to use
some useful utils from it or add the new if required.

----

4. In the `x/rarimocore/keeper` define function that creates the operation and define function call where it is
   required.

----

5. In the `x/rarimocore/keeper/msg_server_confirmation.go` extend the existing logic
   of `getContent(ctx sdk.Context, op types.Operation) (merkle.Content, error)` method.
   For example add:

   ```go
    case types.OpType_FEE_TOKEN_MANAGEMENT:
        manage, err := pkg.GetFeeTokenManagement(op)
        if err != nil {
            return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "failed to unmarshal details")
        }
        return k.getFeeTokenManagementContent(ctx, op.Index, manage)
   ```

   Example:

   ```go
   func (k msgServer) getContent(ctx sdk.Context, op types.Operation) (merkle.Content, error) {
     switch op.OperationType {
     case types.OpType_TRANSFER:
       transfer, err := pkg.GetTransfer(op)
       if err != nil {
         return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "failed to unmarshal details")
       }

       return k.getTransferOperationContent(ctx, transfer)
     case types.OpType_CHANGE_PARTIES:
       change, err := pkg.GetChangeParties(op)
       if err != nil {
         return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "failed to unmarshal details")
       }

       return pkg.GetChangePartiesContent(change)
     case types.OpType_FEE_TOKEN_MANAGEMENT:
       manage, err := pkg.GetFeeTokenManagement(op)
       if err != nil {
         return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "failed to unmarshal details")
       }
       return k.getFeeTokenManagementContent(ctx, op.Index, manage)
     default:
       return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid operation")
     }
   }
   ```


----

6. Also, you can provide additional logic in `ApplyOperation(ctx sdk.Context, op types.Operation) error` to execute some
   stuff after signing if required.

----

7. Extend `tss-svc`
   service ` GetContents(client *grpc.ClientConn, operations ...*rarimo.Operation) ([]merkle.Content, error)` method
   in `internal/core/controllers/util.go` to include new operation in the signing process.
