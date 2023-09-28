# `x/identity`

## Abstract

The goal of identity module is to provide an aggregated information about Iden3 issuer states.
It listens to deployed into Rarimo EVM StateV2 contract and collects the state transition events.

----

## Concepts

Collected events' information is recoded to the dynamic Merkle tree (based on treap data structure).
The proof of concept for dynamic merkle tree is here: "<https://github.com/olegfomenko/go-treap-merkle>".

To collect events `identity` module uses `evm` module hooks by defining the following method:
```go
func (k Keeper) PostTxProcessing(ctx sdk.Context, msg core.Message, receipt *ethtypes.Receipt) error
```

It receives all emitted transaction logs in EVM module and filters them to process only `StateTransited` events.

After all state updates in the module `EndBlock` method the `IDENTITY_AGGREGATED_TRANSFER` operation
will be created with current GIST and States root hash information. Also, the in all changed states `lastUpdateOperationIndex` will be updated.

### Architecture

The basic methods of dynamic merkle to work with is defined in `x/identity/keeper/treap.go`:
```go
func (t Treap) Split(ctx sdk.Context, root, key string) (string, string)

func (t Treap) Merge(ctx sdk.Context, r1, r2 string) string

func (t Treap) Remove(ctx sdk.Context, key string)

func (t Treap) Insert(ctx sdk.Context, key string, priority uint64)

and other...
```

That method are operating on the `Node` structure index.

Node index is based on corresponding `StateInfo` object hash (use `CalculateHash()`).

Also, every `Node` contains additional hash field that stores the `HASH(self,HASH(left,right))`.

Please refer to the corresponding `hash(a, b string) string` and `updateNode(ctx sdk.Context, node *types.Node)` to get more context about how we're constructing merkle tree.

In the `x/identity/keeper/keeper.go` file we are defining the main entrypoints to interact with:

- updates the dynamic Merle tree corresponding to the received event information:

  ```go
  func (k Keeper) UpdateIdentity(ctx sdk.Context, gist string, id string, hash string)
  ```

- returns the Merkle path for any tree node.
  ```go
  func (k Keeper) Path(ctx sdk.Context, id string) []string
  ```
  
----

## State

### Params

Definition:
  ```protobuf
  message Params {
    // Linear congruential generator params
    // https://en.wikipedia.org/wiki/Linear_congruential_generator
    uint64 lcgA = 1;
    uint64 lcgB = 2;
    uint64 lcgMod = 3;
    uint64 lcgValue = 4;
    // Address of identity state smart contract in rarimo chain
    string identityContractAddress = 5;
    string chainName = 6;
    string GISTHash = 7;
    uint64 GISTUpdatedTimestamp = 8;
    string treapRootKey = 9;
    repeated string statesWaitingForSign = 10;
  }
  ```

  <details>
    <summary>Example</summary>

    ```json
    {
      "params": {
        "lcgA": "1664525",
        "lcgB": "1013904223",
        "lcgMod": "4294967296",
        "lcgValue": "2900471886",
        "identityContractAddress": "0x753a8678c85d5fb70A97CFaE37c84CE2fD67EDE8",
        "chainName": "Rarimo",
        "GISTHash": "0x049f1325d5227edcefbca1dc4dc1b76dd981e54c874ec49ba964443086b49950",
        "GISTUpdatedTimestamp": "1691866982",
        "treapRootKey": "0x36141b81b879c28068b3df0bbe9fad19c202b3ef7a140046e018c4153a8ce4c1",
        "statesWaitingForSign": []
      }
    }
    ```
  </details>

### Node

Definition:
  ```protobuf
  message Node {
    // Node key (identity state hash)
    string key = 1;
    // Node priority (should be random)
    uint64 priority = 2;
    // Node left son key
    string left = 4;
    // Node right son key
    string right = 5;
    // Merkle hash. H = Hash(Hash(left_key|right_key)|self_key)
    string hash = 6;
    // Hash(left_key|right_key)
    string childrenHash = 7;
  }
  ```

  <details>
    <summary>Example</summary>

    ```json
    {
      "node": {
        "key": "0x36141b81b879c28068b3df0bbe9fad19c202b3ef7a140046e018c4153a8ce4c1",
        "priority": "4267815944",
        "left": "0x2d6a7c009097397071398f3b2a1855a5df9f6d9ce258846ba92de23aee0dfdf9",
        "right": "0x371e7f58b71fea562aa728619fed387134051e19a3efe0dac2c09557852c5a5c",
        "hash": "0x9cc3d207a5e341279f698cad512f517edb0e9d8df44181680f8d1d75b5573be2",
        "childrenHash": "0xeb2c9ef79b7415a7d38bd1497550084f6bf3ba2f871771b5030c8084b9ff51c8"
      }
    }
    ```
  </details>

### StateInfo

Definition:
  ```protobuf
message StateInfo {
  // State info index (issuer id)
  string index = 1;
  // State hash in hex with 0x
  string hash = 2;
  // Creation timestamps
  uint64 createdAtTimestamp = 3;
  // Creation block (will not be used in state hash)
  uint64 createdAtBlock = 4;
  // Index of last update/create operation (will not be used in state hash)
  string lastUpdateOperationIndex = 5;
}
  ```

  <details>
    <summary>Example</summary>

    ```json
    {
      "state": {
        "index": "0x106d23bb7bedce6caadddf7480ade7f2b8e93fa304fc51cc4030a66de90001",
        "hash": "0x22121ba37492dbb16203cd6dcdb446c4a5c56a4395b145b9403819bcf34141bf",
        "createdAtTimestamp": "1691866982",
        "createdAtBlock": "923813",
        "lastUpdateOperationIndex": "0x2fd7af49f584db04cc8048fd09be7fccf01bd7efe6c93127c0dbae55e643d625"
      }
    }
    ```
  </details>

----

## Transactions

### SetIdentityContractAddress

**SetIdentityContractAddress** - sets the Rarimo EVM StateV2 contract address.
Can be called only once when the current address is 0x.
  ```protobuf
  message MsgSetIdentityContractAddress {
    string creator = 1;
    string address = 2;
  }
  ```

----