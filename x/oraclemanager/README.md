---
layout: default
title: x/oraclemanager
---

# `x/oraclemanager`

## Abstract

The `oraclemanager` cosmos module contains logic for managing distibuted public oracles that supports our bridge by
delivering
information about transfers.

----

## Concepts

In the `oraclemanager` the oracles will be organized into the groups for every supported chain.
Of course, we imply that the same account can be the oracle in several groups.

Also, oracle will have the following related data:

- Status: Slashed, Freezed, Jailed, Active and Inactive
- Stake amount
- Missed count - the count of missed vote
- Violations count - the count of possible_malicious
- Freeze end block - the block when an oracle becomes slashed.
- Voted operations
- Created operations

The following list full describes system rules and architecture:

1. There will be configured the following parameters:
    - min_oracle_stake - minimum amount of RMO tokens to become an oracle.
    - check_op_delta - the amount of blocks after operation voting finish to perform slashing of malicious oracles (
      described below in 6).
    - max_violation_count - the amount of violation that an oracle can reach before it will be freezed (described below
      in 6).
    - max_missed_count - the amount of violation that oracle can reach before it will be jailed (described below in 6).
    - slashed_freeze_blocks - the amount of blocks until the oracle stake will be burned (described below in 7).
    - min_oracles_count - minimal count of oracles to verify operations on the certain chain (global for all chains).

2. Oracles will be separated to the lists of accounts that support certain chains (chains defined in params
   of `tokenmanager` module).
   So every chain will have their own list of oracles that observes its state, submits and votes for operations.

3. To become an oracle for a certain chain account (that will be used by oracle service) should stake
   at least min_oracle_stake tokens through the `oraclemanager` by submitting the `Stake` transaction.
   To use the same oracle account for several chains, the account owner should stake tokens for every chain separately.

4. It is possible to unstake tokens and finish taking part in voting and creating new operations if the oracle owner
   wants.
   Through the `oraclemanager` oracle owner can submit `Unstake` transaction and return the staked coins.
   After that oracle will not be able to create operations and vote for the created one.

5. After staking, if the chain has at least min_oracles_count oracles, every oracle can create and vote for new
   operations,
   using the proxy method in `oraclemanager` (`CreateTransferOperation` and `Vote` that will trigger logic in
   the `rarimocore` module).
   Oracle voting power will be calculated depending on the oracle account staked tokens amount.

6. The `oraclemanager` will control votes and new operations to perform slashing of malicious oracles.
   For every operation that has Approved or NotApproved status after check_op_delta blocks `oraclemanager` EndBlock
   method will iterate over all votes and:
    - Increase missed counter for oracles that haven't submitted their Vote.
    - Increase violations counter for oracles that have submitted a `Vote` with answer `NO` for `Approved` operation.
    - Increase violations counter for oracles that have submitted a `Vote` with answer `YES` for `NotApproved`
      operation.
    - Increase violations counter for oracles that have created a `NotApproved` operation.

Also, EndBlock method will iterate over all oracles and:

- if oracle reaches max_violation_count its status will be set as `Freezed`. Freezed oracle accounts will not be able to
  create operations and vote for the created one.
- if oracle reaches max_missed_count its status will be set as `Jailed`. Jailed oracle accounts will not be able to
  create operations and vote for the created one.

7. After freezing the oracle owner can create a proposal to unfreeze his account
   (`CreateOracleUnfreezeProposal` transaction on `oraclemanager` module) and if it becomes accepted the oracle account
   will become `Active` with zero violations count.
   If after slashed_freeze_blocks from freezing the proposal was not created or
   accepted the oracle stake will be burned and that oracle account never will not be able to return its stake or
   become an active oracle.

8. For the `Jailed` oracles, the oracle owner should execute the `Unjail` transaction.

----

## State

### Params

Definition:

```protobuf
message Params {
  string minOracleStake = 1;
  uint64 checkOperationDelta = 2;
  uint64 maxViolationsCount = 3;
  uint64 maxMissedCount = 4;
  uint64 slashedFreezeBlocks = 5;
  uint64 minOraclesCount = 6;
  string stakeDenom = 7;
  string voteQuorum = 8;
  string voteThreshold = 9;
}
```

Example:

```json
{
  "minOracleStake": "1000000",
  "checkOperationDelta": "10",
  "maxViolationsCount": "10",
  "maxMissedCount": "10",
  "slashedFreezeBlocks": "240",
  "minOraclesCount": "1",
  "stakeDenom": "stake",
  "voteQuorum": "0.900000000000000000",
  "voteThreshold": "0.667000000000000000"
}
```

### Oracle

Definition:

```protobuf
enum OracleStatus {
  Inactive = 0;
  Active = 1;
  Jailed = 2;
  Freezed = 3;
  Slashed = 4;
}

message OracleIndex {
  string chain = 1;
  string account = 2;
}

message Oracle {
  OracleIndex index = 1;
  OracleStatus status = 2;
  string stake = 3;
  uint64 missedCount = 4;
  uint64 violationsCount = 5;
  uint64 freezeEndBlock = 6;
  uint64 votesCount = 7;
  uint64 createOperationsCount = 8;
}
```

Example:

```json
{
  "index": {
    "chain": "Solana",
    "account": "rarimo1g9p4ejp9p877j9vdnuyqtgqm4lhm4f6j7uaztx"
  },
  "status": "Active",
  "stake": "1000000",
  "missedCount": "0",
  "violationsCount": "0",
  "freezeEndBlock": "0",
  "votesCount": "22",
  "createOperationsCount": "22"
}
```

----

## Transactions

### CreateTransferOperation

**CreateTransferOperation** - creates Operation with type `TRANSFER` and `INITIALIZED` status.
Metadata should be provided in case of first NFT transfer. Tx, EventId, Sender can be specified in native for source
chain format.
Other data should be formatted into hex with `0x` prefix.

```protobuf
message MsgCreateTransferOp {
  string creator = 1;
  // Information to identify transfer
  string tx = 2;
  string eventId = 3;
  string sender = 4;
  // Information about deposit
  string receiver = 5;
  string amount = 6;
  string bundleData = 7;// hex-encoded
  string bundleSalt = 8;// hex-encoded
  // Information about current and target chains
  rarimo.rarimocore.tokenmanager.OnChainItemIndex from = 9;
  rarimo.rarimocore.tokenmanager.OnChainItemIndex to = 10;
  rarimo.rarimocore.tokenmanager.ItemMetadata meta = 11; // Optional (if item currently does not exists)
}
```

### CreateIdentityGISTTransferOperation

**CreateIdentityGISTTransferOperation** - creates Operation with type `IDENTITY_GIST_TRANSFER` and `INITIALIZED`
status.

```protobuf
message MsgCreateIdentityGISTTransferOp {
   string creator = 1;
   // Hex 0x
   string contract = 2;
   string chain = 3;
   // Hex 0x
   string GISTHash = 4;
   // Hex 0x
   string GISTReplacedBy = 5;
   // Dec
   string GISTCreatedAtTimestamp = 6;
   string GISTCreatedAtBlock = 7;
   // HEx 0x
   string replacedGISTtHash = 8;
}
```

### CreateIdentityStateTransferOperation

**CreateIdentityStateTransferOperation** - creates Operation with type `IDENTITY_STATE_TRANSFER` and `INITIALIZED`
status.

```protobuf
message MsgCreateIdentityStateTransferOp {
   string creator = 1;
   // Hex 0x
   string contract = 2;
   string chain = 3;
   // Hex 0x
   string id = 5;
   // Hex 0x
   string stateHash = 6;
   // Dec
   string stateCreatedAtTimestamp = 7;
   string stateCreatedAtBlock = 8;
   // HEx 0x
   string replacedStateHash = 17;
}
```

### CreateWorldCoinIdentifierTransferOperation

**CreateWorldCoinIdentifierTransferOperation** - creates Operation with type `WORLDCOIN_IDENTITY_TRANSFER` and
`INITIALIZED` status.

```protobuf
message MsgCreateWorldCoinIdentityTransferOp {
   string creator = 1;
   string contract = 2;
   string chain = 3;
   // Hex 0x uint256
   string prevState = 4;
   // Hex 0x uint256
   string state = 5;
   // Dec uint256
   string timestamp = 6;
   uint64 blockNumber = 7;
}
```

### Vote

**Vote** - vote for operation. Vote power will be equal to the voter staked balance. After total voting power reaches
required quorum operation status changes to `APPROVED` or `NOT_APPROVED`.

```protobuf
message MsgVote{
  rarimo.rarimocore.oraclemanager.OracleIndex index = 1;
  string operation = 2;
  rarimo.rarimocore.rarimocore.VoteType vote = 3;
}
```

### Stake

**Stake** - stake tokens to become active oracle. Also it is possible to re-activate unstaked oracle and stake more
tokens use that message.

```protobuf
message MsgStake {
  rarimo.rarimocore.oraclemanager.OracleIndex index = 1;
  string amount = 2;
}
```

### Unstake

**Unstake** - unstake all tokens and stop to be an active oracle

```protobuf
message MsgUnstake {
  rarimo.rarimocore.oraclemanager.OracleIndex index = 1;
}
```

### Unjail

**Unjail** - unjail `Jailed` oracle

```protobuf
message MsgUnjail{
  rarimo.rarimocore.oraclemanager.OracleIndex index = 1;
}
```

----

## Governance

### OracleUnfreezeProposal

**OracleUnfreezeProposal** - unfreeze slashed oracle proposal.

```protobuf
message OracleUnfreezeProposal {
  string title = 1;
  string description = 2;
  rarimo.rarimocore.oraclemanager.OracleIndex index = 3;
}
```

### ChangeParamsProposal

**ChangeParamsProposal** - changing of module params

```protobuf
message ChangeParamsProposal {
  string title = 1;
  string description = 2;
  Params params = 3 [(gogoproto.nullable) = false];
}
```

----
