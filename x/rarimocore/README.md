---
layout: default
title: x/rarimocore
---

# `x/rarimocore`

## Abstract

The `rarimocore` cosmos module contains logic for all cross-chain operations, its votes, confirmations and signers
management.

You can explore the messages and RPC server definition of the `rarimocore` module in the `rarimo-core/proto/rarimocore`
package.

----

## Concepts

The main flow consists of three steps:

1. Any of Oracles creates the operation on the core
2. Oracles vote for the operation correctness, and after that operation defined as approved/not approved.
3. Signature producers sign the operation
4. Signature producers submit confirmation to the core, and after that operation defined as signed.

Also, the shorter flow can exist: when operation is created by governance or module Begin/EndBlock.
In such case operation can be already approved.

### Operation indexes

- Transfer operation: `HASH(tx, event, chain)`
- Change parties: `HASH(new set, new key, signature)`
- Fee contract management: `HASH(block height, chain, fee token contract, fee amount)`
- Contract upgrade: `HASH(block, content [depends on chain])`
- Identity default
  transfer: `HASH(source contract, id, state hash, state timestamps, state replaced by, GIST hash, GIST replaced by, GIST timestamps, replaced state hash, replaced GIST hash)`
- Identity state transfer: `HASH(source contract, id, state hash, state timestamps, replaced state hash)`
- Identity GIST transfer: `HASH(source contract, id, GIST hash, GIST timestamps, replaced GIST hash)`
- Identity aggregated transfer: `HASH(GIST, timestamp, states root, contract address, chain)`
- WorldCoin identity transfer: `HASH(source contract, old state, new state, timestamp)`
- CSCA root update: `HASH(hashPrefix, root, timestamp)`, where `hashPrefix="Rarimo CSCA root"`

To add new operation check the following [manual](../../docs/common/core/001-adding-operation.md).

----

## State

### Params

**Params** - defines all TSS information: parties, threshold, pub key, etc.
Note that public keys (parties and global) should be in uncompressed form without leading 0x04 (64 bytes format).

Definition:

```protobuf
enum PartyStatus {
  Active = 0;
  Frozen = 1;
  Slashed = 2;
  Inactive = 3;
}

message Party {
  // PublicKey in hex format
  string pubKey = 1;
  // Server address connect to
  string address = 2;
  // Rarimo core account
  string account = 3;
  PartyStatus status = 4;
  uint64 violationsCount = 5;
  uint64 freezeEndBlock = 6;
  string delegator = 7;
  // service information used in initial setup
  string committedGlobalPublicKey = 8;
}

// Params defines the parameters for the module.
message Params {
  string keyECDSA = 1;
  uint64 threshold = 2;
  repeated Party parties = 3;
  bool isUpdateRequired = 5;
  string lastSignature = 6;
  string stakeAmount = 7;
  string stakeDenom = 8;
  uint64 maxViolationsCount = 9;
  uint64 freezeBlocksPeriod = 10;
}
```

Example:

```json
{
  "keyECDSA": "0x928d7a512b18fcbfd51c1b050e2d498f962e2c244bb7495e253731cddcfd164ef32c213e3b4fd4185d1de33dd596061473392aa73b532906e553e543801c0f3a",
  "threshold": "2",
  "parties": [
    {
      "pubKey": "0xb7e8d31ac044cae1ae0ad6f440c25bdf313828b3dd048bc290a8f839c41dd62b50f50693ac4c36ffea12b43e4b59c694a69f2e25a3a4b3135a444df6b5dd6423",
      "address": "tss-1:9000",
      "account": "rarimo1l2vdscjfm289mdxnlnvfwscku4w2l3ljt97kdq",
      "status": "ACTIVE",
      "violationsCount": 0,
      "freezeEndBlock": 0,
      "delegator": "",
      "committedGlobalPublicKey": ""
    },
    {
      "pubKey": "0x4b5139cecb56617c32133885bc7b93da9900e3bf6e5e6d7ce012780f668f1379f5f2fe1a9b40a36421104b6873c53f353aaaa1a696c19099022c290cbc6c8b89",
      "address": "tss-2:9000",
      "account": "rarimo1cqtx4qdtdemula69hpv9ksvg769u0tujgqqyhc",
      "status": "ACTIVE",
      "violationsCount": 0,
      "freezeEndBlock": 0,
      "delegator": "",
      "committedGlobalPublicKey": ""
    },
    {
      "pubKey": "0xc674cc7a7ec0ceacc7723e7a807c09ffe44b84bbb45de1c586dbc3fcc3724c71bbd9dba4fcf977772128f9e8c4b2f6b345f3d8ad944e853dc352a6eb343d1775",
      "address": "tss-3:9000",
      "account": "rarimo13xrlm9cg8ugp327qr0a4k2r9fvkqrsvtm9pcsn",
      "status": "ACTIVE",
      "violationsCount": 0,
      "freezeEndBlock": 0,
      "delegator": "",
      "committedGlobalPublicKey": ""
    }
  ],
  "isUpdateRequired": false,
  "lastSignature": "0x00",
  "stakeAmount": "1000000",
  "stakeDenom": "stake",
  "maxViolationsCount": 10,
  "freezeBlocksPeriod": 100
}
```

### Operation

**Operation** - defines generic operation to be signed by core.

Definition:

```protobuf
enum OpType {
  TRANSFER = 0;
  CHANGE_PARTIES = 1;
  FEE_TOKEN_MANAGEMENT = 2;
  CONTRACT_UPGRADE = 3;
  IDENTITY_DEFAULT_TRANSFER = 4;
  IDENTITY_AGGREGATED_TRANSFER = 5;
}

enum OpStatus {
  INITIALIZED = 0;
  APPROVED = 1;
  NOT_APPROVED = 2;
  SIGNED = 3;
}

message Operation {
  // Should be in a hex format 0x...
  string index = 1;
  opType operationType = 2;
  // Corresponding to type details
  google.protobuf.Any details = 3;
  opStatus status = 4;
  string creator = 5;
  uint64 timestamp = 6;
}
```

Example:

```json
{
  "index": "0x683f666c05682cdbf8ad0d2bd988515deca0dc82296d4366debfbff738f7eacb",
  "operationType": "TRANSFER",
  "details": {
    "@type": "/rarimo.rarimocore.rarimocore.Transfer",
    "origin": "0x683f666c05682cdbf8ad0d2bd988515deca0dc82296d4366debfbff738f7eacb",
    "tx": "e2YTbRzKSPXHkesQhSaANDEGyHHqBfyytEoWCeBjJ9Ugzzrh37kvpc5MDMNHAbwL17eAVQnrMc48EpD2RYx2GeL",
    "eventId": "0",
    "sender": "FCpFKSEboCUGg1Qs8NFwH2suMAHYWvFUUiVWk8cKwNqf",
    "receiver": "0xd30a6d9589a4ad1845f4cfb6cdafa47e2d444fcc568cef04525f1d700f4e53aa",
    "amount": "1000",
    "bundleData": "",
    "bundleSalt": "",
    "from": {
      "chain": "Solana",
      "address": "",
      "tokenID": ""
    },
    "to": {
      "chain": "Solana",
      "address": "",
      "tokenID": ""
    },
    "meta": null
  },
  "status": "SIGNED",
  "creator": "rarimo1g9p4ejp9p877j9vdnuyqtgqm4lhm4f6j7uaztx",
  "timestamp": "113310"
}
```

### Confirmation

**Confirmation** - defines tss signature for some set of operations.

Definition:

```protobuf
message Confirmation {
  // hex-encoded
  string root = 1;
  // hex-encoded
  repeated string indexes = 2;
  // hex-encoded
  string signatureECDSA = 3;
  string creator = 4;
}
```

Example:

```json
{
  "root": "0x17831935b02ca1784e920245c2ea4f6d5dd6298b1df4e3c870e5f6f45d2b0797",
  "indexes": [
    "0x836638ffdc89588a635c24921f1b524d118e9e75c97d38b675620c4211493d7f"
  ],
  "signatureECDSA": "0x8bf642743d6ca2361d503c7accdfdd269ff043a070beb22904a5a9c4442c1aa20ae3252198d60dedf8316908408099022c16c1a79823ccd3c49cdab84e473b7a00",
  "creator": "rarimo1yxzygxdk3eujpudr3ynux4gkg7r6ss6csqf9ew"
}
```

### Vote

**Vote** - defines validator votes for operation validity.

Definition:

```protobuf
enum VoteType {
  YES = 0;
  NO = 1;
}

message VoteIndex {
  string operation = 1;
  string validator = 2;
}

message Vote {
  VoteIndex index = 1;
  VoteType vote = 2;
}
```

Example:

```json
{
  "index": {
    "operation": "0x683f666c05682cdbf8ad0d2bd988515deca0dc82296d4366debfbff738f7eacb",
    "validator": "rarimo1eqmyvqge5mu7cfv4gjy8y7smfxhq5r0ctljf0n"
  },
  "vote": "YES"
}
```

Example:

### Violation report

**ViolationReport** - stored reports from TSS about other parties violations. Used to calculate violations amount and
exclude from parties set upon reaching the configured aount.

Definition:

```protobuf
enum ViolationType {
  Offline = 0;
  Spam = 1;
  Other = 3;
}

message ViolationReportIndex {
  string sessionId = 1;
  string offender = 2;
  ViolationType violationType = 3;
  string sender = 4;
}

message ViolationReport {
  ViolationReportIndex index = 1;
  // Optional message
  string msg = 3;
}
```

----

## Keepers

The `rarimocore` module only exposes one keeper, which can be used to create operations, access TSS parties information
in params, access and create votes.

Some of useful Keeper methods:

### CreateTransferOperation

**CreateTransferOperation** - to create transfer operation by other modules.

Definition:
`CreateTransferOperation(ctx sdk.Context, creator string, transfer *types.Transfer, approved bool) error`

Flow:

- Get source network from `tokenmanager` module. Check that bridge transfers are acceptable for that network.
- Create operation entry and check that operation does not exist.
- Only not approved operation can be replaced. If operation already exists and has to be replaced then we have to clear
  all existing votes.
- If approve is required then will call `ApproveOperation`.

### ApproveOperation

**ApproveOperation** - approve operation (change status to `APPROVED`).

Definition:
`ApproveOperation(ctx sdk.Context, op types.Operation) error`

Flow:
The method is designed to execute some special logic for operation approvals.
Currently, only for transfer operation such logic exists:

- Get `CollectionData` for source token
- Get `OnChainItem` for source token
- If `OnChainItem` does not exist then create `Item` (also does not exist), create `OnChainItem`, set seed if provided.
- Get target `OnChainItem`. If it does not exist then create it.

Also, the opposite method exists: `UnapproveOperation(ctx sdk.Context, op types.Operation) error`

### CreateVote

**CreateVote** - save vote for operation. Used by `oraclemanager` module.

Definition:
`CreateVote(ctx sdk.Context, vote types.Vote) (bool, error)`

Flow:

- Check that operation exists
- Check that vote does not exist
- Check operation status (should be `INITIALIZED`)
- Save vote

### CreateIdentityDefaultTransferOperation/CreateIdentityStateTransferOperation/CreateIdentityGISTTransferOperation/CreateWorldCoinIdentityTransferOperation

**CreateIdentityDefaultTransferOperation** - used by `oraclemanager` module to create identity transfer operations

Definition:
`CreateIdentityDefaultTransferOperation(ctx sdk.Context, creator string, transfer *types.IdentityDefaultTransfer) error`

Flow:

- Get source network from `tokenmanager` module. Check that bridge transfers are acceptable for that network.
- Create operation entry and check that operation does not exist.
- Only not approved operation can be replaced. If operation already exists and need to be replaced then we have to clear
  all existing votes.

### CreateFeeTokenManagementOperation

**CreateFeeTokenManagementOperation** - used by `tokenmanager` module to create fee management operation after proposal
acceptance.

Definition:
`CreateFeeTokenManagementOperation(ctx sdk.Context, op *types.FeeTokenManagement) error`

Flow:

- Create operation entry and check that operation does not exist.
- Approve operation

Also, use

- `CreateWithdrawFeeOperation(ctx sdk.Context, token tokentypes.FeeToken, chain string, receiver string, nonce string) error`
- `CreateUpdateFeeTokenOperation(ctx sdk.Context, token tokentypes.FeeToken, chain string, nonce string) error`
- `CreateRemoveFeeTokenOperation(ctx sdk.Context, token tokentypes.FeeToken, chain string, nonce string) error`
- `CreateAddFeeTokenOperation(ctx sdk.Context, token tokentypes.FeeToken, chain string, nonce string) error`

### CreateCSCARootUpdateOperation

**CreateCSCARootUpdateOperation** - used by `cscalist` module to create CSCA root update operation after proposal pass

Definition: 
`CreateCSCARootUpdateOperation(ctx sdk.Context, creator string, update *rarimo.CSCARootUpdate) (string, error)`

Flow:
- Create proposal to replace/edit CSCA list (see [x/cscalist/README.md](../cscalist/README.md))
- When proposal is accepted, root key is updated
- On the end block operation with **Merkle root** is created
- Operation is auto-approved

----

## Transactions

### CreateChangePartiesOperation

**CreateChangePartiesOperation** - creates new operation with type `CHANGE_PARTIES` and `APPROVED` status.
Used by TSS producers to reshare keys when parties set has been changed.
Signature field contains the signature of new public key by old public key and will be used to change keys on the bridge
contracts.

```protobuf
message MsgCreateChangePartiesOp {
  string creator = 1;
  repeated Party newSet = 2;
  string newPublicKey = 3;
  string signature = 4;
}
```

### CreateConfirmation

**CreateConfirmation** - creates the confirmation for set of operations. After confirmation created user can use
provided
signature and merkle tree to execute event unlocking on destination chain. Confirmation can be created only for
operation with `APPROVED` status.
After hash and signature verification operation status will be changed to `SIGNED`.

- Verify signature for provided tree root corresponds to the stored in params public key
- Verify that creator is a party
- Check that confirmation does not exist
- Build tree using crypto package methods and go-merkle library
- Save new confirmation and update last signature in params
- Use ApplyOperation for every operation in the set

Apply operation is designed to execute special functional for any operation after is becomes signed.
Also, ApplyOperation method changes the operation status to `SIGNED`

```protobuf
message MsgCreateConfirmation {
  string creator = 1;
  // hex-encoded
  string root = 2;
  repeated string indexes = 3;
  // hex-encoded
  string signatureECDSA = 4;
}
```

### SetupInitial

**SetupInitial** - used only in setup phase (when ECDSA key in params is empty and update flag is true). After threshold
keys generation TSS parties submits
`MsgSetupInitial` to confirm successful keys generation and provide new party public key data.

```protobuf
  message MsgSetupInitial {
  string creator = 1;
  string partyPublicKey = 2;
  string newPublicKey = 3;
}
```

### ChangePartyAddress

**ChangePartyAddress** - changes party address. Only party account can execute. Change will be applied immediately.

```protobuf
message MsgChangePartyAddress {
  string creator = 1;
  string newAddress = 2;
}
```

### Stake

**Stake** - stake to become a TSS producer.

- Check that account has not staked yet.
- Transfer coins from account to module
- Add inactive party to the list and set `UpdateIsRequired` flag.

```protobuf
message MsgStake {
  string creator = 1;
  string account = 2;
  string address = 3;
  string trialPublicKey = 4;
}
```

### Unstake

**Unstake** - unstake and stop producing tss signature.

- Check that party exist and active.
- Check that message signer is party account or delegator
- Transfer coins from module account to party or delegator (depends on msg signer)
- Remove party from list and set `UpdateIsRequired` flag.

```protobuf
 message MsgUnstake {
  string creator = 1;
  string account = 2;
}
```

### CreateViolationReport

**CreateViolationReport** - RPC for TSS services to repost about other TSS violation.

- Check that sender and offender is active parties
- Check that violation has not been created yet
- Save violation report
- Iterate over existing reports, increment party violations if there are more than threshold violations for same
  session.
- If party violations reaches `maxViolationsCount` (in params) then change party status to `Frozen` and
  set `UpdateIsRequired` flag.

```protobuf
 message MsgCreateViolationReport {
  string creator = 1;
  string sessionId = 2;
  rarimo.rarimocore.rarimocore.ViolationType violationType = 3;
  string offender = 4;
  // Optional message
  string msg = 5;
}
```

----

## Governance

### UnfreezeSignerPartyProposal

**UnfreezeSignerPartyProposal** - if party reaches ```maxViolationsCount``` is becomes freezed and there will be some
time (```freezeBlocksPeriod```) when unfreeze proposal can be created. If after ```freezeBlocksPeriod``` party is still
freezed it will be slashed forever.

```protobuf
message UnfreezeSignerPartyProposal {
  string title = 1;
  string description = 2;
  string account = 3;
}
```

### ReshareKeysProposal

**ReshareKeysProposal** - will set `isUpdateRequired=true` and launch key resharing if any of parties lost his secret
key data.

```protobuf
message ReshareKeysProposal {
  string title = 1;
  string description = 2;
}
```

### SlashProposal

**SlashProposal** - proposal to slash a party immediately.

```protobuf
message SlashProposal {
  string title = 1;
  string description = 2;
  string party = 3;
}
```

### DropPartiesProposal

**DropPartiesProposal** - in some cases we will need to drop all parties and regenerate keys again.
Using that proposal you can unstake all parties (that is active/inactive or freezed) and clear parties set.
After the new parties should use ```Stake``` operation and ```SetupInitial``` operation to generate new parties set.
Also changing bridge contracts will be also required because we can not provide signature for changing public key.

```protobuf
message DropPartiesProposal {
  string title = 1;
  string description = 2;
}
```

----
