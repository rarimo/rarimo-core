---
layout: default
title: Oracles (Savers)
---

# Oracles (Savers)

- Repo: [Saver Lib](https://github.com/rarimo/saver-grpc-lib)
- Repo: [EVM Saver](https://github.com/rarimo/evm-saver-svc)
- Repo: [Solana Saver](https://github.com/rarimo/solana-saver-svc)
- Repo: [Near Saver](https://github.com/rarimo/near-saver-svc)

Managing events' creation requires some special logic for every chain that will observe blockchain state for new events,
parse in accordance with chain rules, etc. Rarimo core can not fetch events data by itself because interaction with
other systems can not be deterministic.
So there should be custom software for observing blockchain state and transferring events from
source chain to Rarimo core. That is why we’ve created Oracles.

Rarimo oracles are built to provide everyone an opportunity to deliver and validate cross-chain messages.
We believe that architecture constructed on worldwide open and transparent oracles brings the real decentralization
into any cross-chain protocol.

In Rarimo oracle’s goal is to observe blockchains state for new events, deliver such events into the Rarimo core and
verify events delivered from other oracles. In particular oracles should observe new token deposits,
create transfer operations, submit them into the core and vote for correctness of submitted operations.

The oracle services designed to be launched by anyone in two supported modes:
- saver (connects to the chain rpc and submits new transfer operations to the Rarimo core)
- voter (fetches new operations from core, verifies the content and votes for its correctness)

Currently, there are three implementations exists:
- evm-saver (used with any evm-compatible chains)
- solana-saver (used with any solana-compatible chains)
- near-saver (used with any near-compatible chains)

----

## Logic

### Native and FT
For native and fungible tokens all information should be pre-defined in collections of `tokenmanager` module.
So oracles should only fetch the corresponding data and submit it in `MsgCreateTransferOp` transaction.

### NFT
NFT tokens collection can contain huge amount of tokens under it, and also they can be minted if future.
So we can not define all tokens in core during initialization or token add operation. That is why collections flow was created.
Using `tokenmanager` `Collection` and `CollectionData` we can define collection global and chain information regardless
of the number of tokens in collection or their metadata.
Token information (`Item` and `OnChainItem`) will be set up during the first token transfer.
During the first transfer oracles should fill the metadata field in `MsgCreateTransferOp` and provide all required token information to create `Item` and `OnChainItem`.

Token addresses are equal to the token collection addresses that is already defined in `Collection` but for token id we should provide more complicated implementation.
Let's describe the flow more accurately:

1. If token was already transferred between source and destination chain no additional actions required. Oracle just 
  fetches `OnChainItem` from core and leaves metadata field empty.

2. If source `OnChainItem` exists but destination `OnChainItem` does no (it means that token never was transferred on destination chain)
  oracle should construct destination `OnChainItem` using the following rules:
  - if destination chain is Solana, tokenId should be derived using bridge address and seed from `Item` metadata (`Item` should exist cause `OnChainItem` exists)
-   otherwise tokenId will be equal to the home chain tokenId (the chain where token was minted firstly).

3. If source `OnChainItem` does not exists it means that we are transferring token at first time and `Item` with its metadata also does not exist.
  So oracle should create `Item` metadata which means fetching on-chain metadata and generation of solana seed. Then using paragraph 2 rules oracle constructs destination `OnChainItem`.
  Source `OnChainItem`, which is home `OnChainItem`, can be constructed using deposit event information.

----

## Library

In `github.com/rarimo/saver-grpc-lib` we defined the common utils for all oracles.
Every oracle should implement `verifiers.TransferOperator` interface for every supported token type
and put that implementations into the `voter.Subscriber` and `voter.Catchupper` using `verifiers.TransferVerifier` wrapper.

Also, we recommend to split `verifiers.TransferOperator` logic in two methods: verifier and message creator.
It allows to use message creator while fetching new events and submitting them into core.

----

## Broadcaster integration

All oracles recommended to integrate with broadcaster service (`broadcaster-svc`) to submit queued transactions. Cause any transaction should contain valid account sequence
(is incremental value) it becomes problematically to submit concurrent transactions.
So our oracles implementations use broadcaster service as an endpoint for sending messages using a queue.

----

## Flow

Saver mode flow:
1. Oracle subscribes to the new events on bridge contract.
2. Oracle receives new event.
3. Oracle creates and submits `MsgCreateTransferOp` using `broadcaster-svc`.

Voter mode flow:
1. Oracle subscribes to the new operations on Rarimo core for defined oracle's chain.
2. Oracle receives new operation.
3. Oracle verifies operation and if it is correct submits `MsgVote` transaction with `YES` vote. Otherwise, oracle submits `MsgVote` transaction with `NO` vote.
4. After operation votes reaches `Quorum` (defined in `Rarimocore` module) `rarimocore` calculates the voting power and vote results. To become `APPROVED` there should be at least `Threshold` (defined in `Gov` module) percentage positive votes.
  Otherwise, operation becomes `NOT_APPROVED`.
