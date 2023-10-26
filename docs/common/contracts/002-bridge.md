---
layout: default
title: Bridge overview
---

# Bridge overview

Rarimo decentralized bridge contract is responsible for managing deposits and withdrawals for Native, fungible and non-fungible tokens. All withdrawal operations are protected by ECDSA secp256k1 threshold (t-n) signature.
This signature is produced by core multi-sig services depending on the validated core state.

The core service stores information to be signed in `operation` entries. `Operation` data structure consist of the following fields:
- index
- operation type
- details
- is signed
- creator address
- timestamp


For token transfer there is an operation type `Type_TRANSFER` and the details field contain the following proto-encoded information:
- origin hash (hash of tx, id and current network)
- tx
- event id (or operation id)
- from chain
- to chain
- receiver
- amount
- bundle info (data and salt)
- token information

For signing operations, the signature producer turns the operation content into special Merkle leaf content, creates a Merkle tree, signs its hash and submits the confirmation operation to the core. After, operation will be marked as signed and users can pool signature and operation information to submit into the bridge contract on the target chain. Smart-contract on the target chain should accept all information, perform validation, withdraw the corresponding token, and store information about the completed withdrawals to prevent double-spending.

----

## Signed content format

Common transfer hash can be derived from the following sequence:
`tokenAddress | tokenName | tokenId | uri | amount | imageUri | imageHash | tokenSymbol | tokenDecimals | salt | bundle | originHash | chainName | receiver | targetBridgeContractAddress`

OriginHash is a hash of `tx | eventId | current network`.
Some parameters can be empty.

On our bridge currently supports the following token types: Native / ERC20 / ERC721 / ERC1155 / MetaplexFT / MetaplexNFT / NearFT / NearNFT

----

## Smart contract requirements

- Smart-contract should store ECDSA t-n public key
- Changing of stored admins key requires t-n signature of the new key produced by the old key
- Withdrawal operations should accept Merkle node content, Merkle path, and admin t-n signature for Merkle root.
- Withdrawal operations should perform a check for Merkle path, signature, and other sensitive stuff

----

## Bridging flow

1. User wants to transfer the token from the current chain to the target one. For this, user submits a deposit transaction to the current chain bridge smart-contract.
2. Any other external services (or users by themselves) submit deposit information into the core service. Core service performs transaction validation before adding them to the ledger.
3. External services (signature producers) observe new operations in the core state and add them to the local pool.
4. Signature producers combine deposits into the Merkle tree, produce the threshold signature, and submit that info to the core. Core service performs the Merkle tree data and signature validation before adding it to the ledger.
5. After signed deposit information appears in the core state it becomes possible to call the withdrawal operation on the target network.
6. Smart-contract on the target network validates information about the deposit, Merkle path, and signature for the Merkle root, and performs a token withdrawal.




