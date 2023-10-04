# Rarimo Identity

Rarimo identity transfers provides issuers with opportunity to publish cheap state updates into Rarimo chain directly,
that allows to skip transferring from Polygon to Rarimo in default flow. Also, all states from all issuers will be
aggregated into one Merkle hash that allows us to update all states by single transaction.

Generally, common flow looks like this:
- There is an operation type `IDENTITY_AGGREGATED_TRANSFER` that is used to transfer aggregated identity state overall issuers.

- Issuer publishes state transition into Rarimo EVM identity state contract.
  That state contract is fully compatible with original Iden3 state smartcontract.

- Identity module collects the state transition events and trigger the dynamic merkle tree update.

- In the `identiy` module `EndBlock` method the corresponding operation will be created.

- After some time, the threshold signature producers provides the ECDSA signature for aggregated state update information hash.

- There is a [modified state smart contracts](https://gitlab.com/rarimo/rarimoid/state-contracts) that accepts aggregated state updates with ECDSA signature instead of ZK proof of state update validity.

- Such smart contract should be deployed into every chain that we have to support.

- The information about state update with its witness (ECDSA signature) can be delivered into modified state smart contracts on the target chain by everyone.
  After such signed update execution, all state updates on Rarimo chain becomes delivered on target chain by single transaction.

----

## DApp flow:

On the DApp side the following flow should be executed to use our aggregated state:

1. Get state information from Rarimo: `/rarimo/rarimo-core/identity/state/{id}`. Also get path from `/rarimo/rarimo-core/identity/state/{id}/proof`.

2. If the last state update timestamp is equal (or less) to the timestamp on target chain - go to (5).

3. Using `lastUpdateOperationIndex` call: `/rarimo/rarimo-core/rarimocore/operation/{index}`. It provides the information about GISTRoot and StatesRoot hash. Use it in state transition.

4. Using `lastUpdateOperationIndex` call: `/rarimo/rarimo-core/rarimocore/operation/{index}/proof`. Using signature, path and information from previous request execute `signedStateTransition` on target chain state smart contract.

5. Generate ZKP using state information from (1) request. Execute ZKP verification on [modified Verificator contract](https://gitlab.com/rarimo/rarimoid/state-contracts) using generated proof information, information and path from (1).
