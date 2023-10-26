---
layout: default
title: x/bridge
---

# `x/bridge`

## Abstract

The main goal of this module is to give the user ability to transfer the RMO native token from the Rarimo chain
to another one (EVM, Near, Solana, etc.) and the opposite: transfer the RMO token from the external chain to the Rarimo chain.
To achieve this we have added some custom logic to the ```cosmos-sdk/x/bank``` module to be able to mint
and burn tokens on the selected account.

----
## Concepts

### Deposit flow
The user calls the ```DepositNative``` RPC method and provides all required arguments for transfer into the parameters.
Under the hood, the module handler will create the ```TransferOperation``` in the ```rarimocore``` module after
the successful burn of the specified amount of tokens from the user’s account balance.
The transfer operation will be approved from the start because the module doesn’t need to check whether the tokens
were deposited cause module burns it by itself.

### Withdraw flow
After the call of the related RPC method, the module handler must check if ```TransferOperation``` was signed in
the ```rarimocore``` module. After all, if checks were successful, module stores used origin in the module store and mints
required amount of tokens to the recipient account balance.

### x/bank updates
To provide the implementation of the bridging feature in our module, the logic of ```cosmos-sdk/x/bank``` needs to be extended.
The ```x/bridge``` module requires the ability to mint and burn system tokens on the selected user account for implementing
deposit and withdraw logic.

So we need to add the custom MintTokens bank keeper’s method which accepts arguments that determine who
is the target of minting, and the amount of tokens to mint. And also BurnTokens should be implemented.

----

## State

### Params

Definition:
```protobuf
message Params {
   option (gogoproto.goproto_stringer) = false;
   string withdrawDenom = 1;
}
```

Example:
```json
{
   "params": {
      "withdrawDenom": "stake"
   }
}
```

### Hash

**Hash** - stores the hash of withdrawal operation to prevent double-spending.

Definition:

```protobuf
message Hash {
   // hex-encoded
   string index = 1;
}
```

----

## Transactions

### DepositNative

**DepositNative** - burns user's tokens and creates transfer operation in ```rarimocore``` module (will be already approved).

```protobuf
message MsgDepositNative {
   string creator = 1;
   // Random 32 bytes slice encoded to the hex string
   string seed = 2;
   // Information about deposit
   string receiver = 3;
   cosmos.base.v1beta1.Coin amount = 4;
   string bundleData = 5;// hex-encoded
   string bundleSalt = 6;// hex-encoded
   // Information about target chain
   rarimo.rarimocore.tokenmanager.OnChainItemIndex to = 7;
}
```

### WithdrawNative

**WithdrawNative** - checks that operation in ```rarimocore``` is signed and valid and mints tokens to the receiver account.
   Operation hash (origin) will be stored in modules hash list.

```protobuf
message MsgWithdrawNative {
   string creator = 1;
   // Evidence information
   string origin = 2;
}
```

----