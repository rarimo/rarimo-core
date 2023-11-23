---
layout: default
title: x/tokenamanager
---

# `x/tokenamanager`

## Abstract

The `tokenamanager` module is responsible for storing information about all supported tokens.

## Concepts

For FT and Native tokens such an information will be strongly predefined. NFT collection can contain huge amount of
tokens under it, and also they can be minted in future.
So we can not define all tokens in core during initialization or token add operation. That is why collections flow was
created.
Using `tokenmanager` `Collection` and `CollectionData` we can define collection global and chain information regardless
of the number of tokens in collection or their metadata.
Token information (`Item` and `OnChainItem`) will be set up during the first token transfer.
During the first transfer oracles should fill the metadata field in `MsgCreateTransferOp` and provide all required token
information to create `Item` and `OnChainItem`.

## State

### Params

**Params** - defines supported networks and their bridge contracts

Definition:

```protobuf
enum NetworkType {
  EVM = 0;
  Solana = 1;
  Near = 2;
  Other = 3;
}

enum NetworkParamType {
  BRIDGE = 0;
  FEE = 1;
  IDENTITY = 2;
}

message Network {
  // network name
  string name = 1;
  NetworkType type = 2;
  repeated NetworkParams params = 3 [(gogoproto.nullable) = false];
}

message NetworkParams {
  NetworkParamType type = 1;
  // Corresponding to type details
  google.protobuf.Any details = 2;
}

message BridgeNetworkParams {
  string contract = 1;
}

message FeeNetworkParams {
  string contract = 1;
  repeated FeeToken feeTokens = 2;
}

message IdentityNetworkParams {
  string contract = 1;
}

message FeeToken {
  // contract address hex
  string contract = 1;
  string amount = 2;
}

// Params defines the parameters for the module.
message Params {
  repeated Network networks = 1;
}
```

Example:

```json
{
  "networks": [
    {
      "name": "Solana",
      "type": "Solana",
      "params": [
        {
          "type": "BRIDGE",
          "details": {
            "@type": "/rarimo.rarimocore.tokenmanager.BridgeNetworkParams",
            "contract": "0x6e632137764e7ba6b053ba055a0e63b5151d5866f846fc5c23302517d8b778a8"
          }
        },
        {
          "type": "FEE",
          "details": {
            "@type": "/rarimo.rarimocore.tokenmanager.FeeNetworkParams",
            "contract": "0x00",
            "feeTokens": [
              {
                "contract": "",
                "amount": "1"
              }
            ]
          }
        }
      ]
    }
  ]
}
```

### Collection

**Collection** - the high level concept of token that is supported on Rarimo bridge.

Definition:

```protobuf
message CollectionMetadata {
  string name = 1;
  string symbol = 2;
  string metadataURI = 3;
}

message CollectionDataIndex {
  // Chain name
  string chain = 1;
  // Collection contract address
  string address = 2;
}

message Collection {
  string index = 1;
  CollectionMetadata meta = 2;
  repeated CollectionDataIndex data = 3;
}
```

Example:

```json
{
  "index": "usdt",
  "meta": {
    "name": "Tether USD",
    "symbol": "USDT",
    "metadataURI": ""
  },
  "data": [
    {
      "chain": "Solana",
      "address": "0xc4192d3818061c9d9333b6e658422b095cc2417dfe313f7286f96108d543253d"
    },
    {
      "chain": "Near",
      "address": "0x757364742e726172696d6f2e746573746e6574"
    },
    {
      "chain": "Goerli",
      "address": "0x2f362dF82622120f7181e432446E7665f51fFb03"
    }
  ]
}
```

### CollectionData

**CollectionData** - defines collection data on the certain chain. Indexed by chain|address and contains id of the
collection.

Definition:

```protobuf
message CollectionData {
  CollectionDataIndex index = 1;
  string collection = 2;
  Type tokenType = 3;
  bool wrapped = 4;
  uint32 decimals = 5;
}
```

Example:

```json
[
  {
    "index": {
      "chain": "Goerli",
      "address": "0x2f362dF82622120f7181e432446E7665f51fFb03"
    },
    "collection": "usdt",
    "tokenType": "ERC20",
    "wrapped": true,
    "decimals": 18
  },
  {
    "index": {
      "chain": "Near",
      "address": "0x757364742e726172696d6f2e746573746e6574"
    },
    "collection": "usdt",
    "tokenType": "NEAR_FT",
    "wrapped": true,
    "decimals": 18
  },
  {
    "index": {
      "chain": "Solana",
      "address": "0xc4192d3818061c9d9333b6e658422b095cc2417dfe313f7286f96108d543253d"
    },
    "collection": "usdt",
    "tokenType": "METAPLEX_FT",
    "wrapped": true,
    "decimals": 9
  }
]
```

### Item

**Item** - defines one simple token across all chains. Contains all metadata required to perform transfers.

Definition:

```protobuf
message ItemMetadata {
  string imageUri = 1;
  // Hash of the token image. Encoded into hex string. (optional)
  string imageHash = 2;
  // Seed is used to generate PDA address for Solana
  string seed = 3;
  string uri = 4;
}

message OnChainItemIndex {
  string chain = 1;
  string address = 2;
  string tokenID = 3;
}

message Item {
  string index = 1;
  string collection = 2;
  ItemMetadata meta = 3;
  repeated OnChainItemIndex onChain = 4;
}
```

Example:

```json
{
  "index": "usdt",
  "collection": "usdt",
  "meta": {
    "imageUri": "",
    "imageHash": "",
    "seed": "",
    "uri": ""
  },
  "onChain": [
    {
      "chain": "Solana",
      "address": "0xc4192d3818061c9d9333b6e658422b095cc2417dfe313f7286f96108d543253d",
      "tokenID": ""
    },
    {
      "chain": "Near",
      "address": "0x757364742e726172696d6f2e746573746e6574",
      "tokenID": ""
    },
    {
      "chain": "Goerli",
      "address": "0x2f362dF82622120f7181e432446E7665f51fFb03",
      "tokenID": ""
    }
  ]
}
```

### OnChainItem

**OnChainItem** - defines one simple token on the certain chain. Indexed by `chain|address|tokenId`.

Definition:

```protobuf
message OnChainItem {
  OnChainItemIndex index = 1;
  string item = 2;
}
```

Example:

```json
[
  {
    "index": {
      "chain": "Goerli",
      "address": "0x2f362dF82622120f7181e432446E7665f51fFb03",
      "tokenID": ""
    },
    "item": "usdt"
  },
  {
    "index": {
      "chain": "Near",
      "address": "0x757364742e726172696d6f2e746573746e6574",
      "tokenID": ""
    },
    "item": "usdt"
  },
  {
    "index": {
      "chain": "Solana",
      "address": "0xc4192d3818061c9d9333b6e658422b095cc2417dfe313f7286f96108d543253d",
      "tokenID": ""
    },
    "item": "usdt"
  }
]
```

----

## Keepers

The `tokenmanager` module only exposes one keeper, which can be used to access information about supported tokens to
transfer and supported network params.


----

## Governance

### SetNetworkProposal

**SetNetworkProposal** - creating or updating existing network params.

```protobuf
message SetNetworkProposal {
  string title = 1;
  string description = 2;
  NetworkParams networkParams = 3;
}
```

### UpdateTokenItemProposal

**UpdateTokenItemProposal** - updating token item params (metadata)

```protobuf
message UpdateTokenItemProposal {
  string title = 1;
  string description = 2;

  repeated Item item = 3;
}
```

### RemoveTokenItemProposal

**RemoveTokenItemProposal** - removing token item from core.

```protobuf
message RemoveTokenItemProposal {
  string title = 1;
  string description = 2;

  repeated string index = 3;
}
```

### CreateCollectionProposal

**CreateCollectionProposal** - adding new collection. CollectionData should be provided. Items and OnChainItems can be
optionally provided.

```protobuf
message CreateCollectionProposal {
  string title = 1;
  string description = 2;

  string index = 3;
  CollectionMetadata metadata = 4;
  // All supported networks described
  repeated CollectionData data = 5;
  repeated Item item = 6;
  repeated OnChainItem onChainItem = 7;
}
```

### UpdateCollectionDataProposal

**UpdateCollectionDataProposal** - updating collection data information.

```protobuf
message UpdateCollectionDataProposal {
  string title = 1;
  string description = 2;

  repeated CollectionData data = 3;
}
```

### AddCollectionDataProposal

**AddCollectionDataProposal** - adding collection data (which means support of new chain for collection)

```protobuf
message AddCollectionDataProposal {
  string title = 1;
  string description = 2;

  repeated CollectionData data = 3;
}
```

### RemoveCollectionDataProposal

**RemoveCollectionDataProposal** - removing collection data (which means disabling of certain chain for collection)

```protobuf
message RemoveCollectionDataProposal {
  string title = 1;
  string description = 2;

  repeated CollectionDataIndex index = 3;
}
```

### RemoveCollectionProposal

**RemoveCollectionProposal** - removing the hole collection

```protobuf
message RemoveCollectionProposal {
  string title = 1;
  string description = 2;

  string index = 3;
}
```

----