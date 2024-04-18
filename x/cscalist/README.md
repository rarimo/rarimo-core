---
layout: default
title: x/cscalist
---

# `x/cscalist`

## Abstract

The module is designed to store and manage [CSCA Master List](https://pkddownloadsg.icao.int/).
The list itself is provided by [ICAO](https://www.icao.int/Pages/default.aspx) and contains root certificates of passport issuing authorities for a lot of countries.
When obtained the passport signer's root certificate, you can check its presence in the Merkle tree,
stored on Rarimo chain in a decentralized manner.
Changing the list is possible via proposal.

----

## Concepts

The CSCA certificates can only be obtained from the ICAO website in form of an LDIF file.
The file is parsed by our [LDIF SDK](https://github.com/rarimo/ldif-sdk) onto the hash values of raw public keys (Poseidon hash with special splitting scheme for 4096-bit RSA keys).
The obtained hashes are used as leaves for the Merkle tree.

Dynamic treap-based Merkle tree is used to store the CSCA public key hashes, see [treap](https://en.wikipedia.org/wiki/Treap).
The reason is good consistency of treap with Cosmos KVStore, while updates of the list may be quite frequent.
For each node the priority is deterministically derived from the key with formula: `priority = poseidon_hash(key) mod MAX_UINT64`, where `key` is Poseidon hash of public key and `MAX_UINT64 = 2^64-1`.

In order to update the tree the proposal can be created. There are two types:
- `ReplaceCSCAListProposal` - replaces the whole tree with a new one, very gas-consuming, suits for initial setup.
- `EditCSCAListProposal` - add or remove certain CSCA public key hashes, flexible and efficient.

On update an operation `CSCARootUpdate` (see `rarimocore` module) is emitted in order to update the root key in external systems, e.g. EVM registration contract.

CLI is provided to work with LDIF (extract public keys, hash them or build a tree with root), query tree and module params, compare the current tree with the locally built from CSCA list.

### Architecture

Implementation of the Merkle tree is copied from LDIF SDK and optimized for using with KVStore. Base methods are defined in `x/cscalist/keeper/treap.go`:
```go
func (t Treap) Insert(ctx sdk.Context, key string) {}

func (t Treap) Remove(ctx sdk.Context, key string) {}

func (t Treap) split(ctx sdk.Context, root, key string) (string, string) {}

func (t Treap) merge(ctx sdk.Context, left, rirght string) string {}
```

The `Node` structure is used to store the tree nodes. The public key hash received from proposal is processed in the following way:
1. The hash must be a hex string with 0x prefix
2. Hex string is set in new `Node` along with its `Hash` and `ChildrenHash` values.
3. On node insertion, deeply in KVStore implementation, string is converted directly to bytes, like `[]byte(key)`.
4. Some operations on treap are performed to build a new tree

Processing in 3 results in redundant gas consumption, because we could store bytes directly that are 2 times smaller than bytes of hex strings themselves.
Also, algorithms of insertion and removal are not currently optimized for multiple values at once, resulting in a huge gas consumption.

## State examples

Models can be found in `proto/cscalist`. Example values are provided in JSON for convenient reading.

Module params
```json
{
  "params": {
    "rootKey": "0x36141b81b879c28068b3df0bbe9fad19c202b3ef7a140046e018c4153a8ce4c1",
    "rootUpdated": false,
    "updatedAtBlock": 854221,
    "lastUpdateOperationIndex": "0x2fd7af49f584db04cc8048fd09be7fccf01bd7efe6c93127c0dbae55e643d625"
  }
}
```

Node of the tree
```json
{
  "node": {
    "key": "0x28815c9a1c9d638886d6ac193df55f98824c491d09bbbd712f96b5adfeba742e",
    "priority": 1833377742,
    "left": "0x",
    "right": "0x022e071cf4eed456dc7f3a36b7b190c6a1f991ef8c5809f612cb193e9c28af78",
    "hash": "0x2fd7af49f584db04cc8048fd09be7fccf01bd7efe6c93127c0dbae55e643d625",
    "childrenHash": "0x022e071cf4eed456dc7f3a36b7b190c6a1f991ef8c5809f612cb193e9c28af78"
  }
}
```

Replace CSCA list proposal
```json
{
        "@type": "/rarimo.rarimocore.cscalist.ReplaceCSCAListProposal",
        "title": "Init",
        "description": "Lorem ipsum dolor sit amet concestetur!",
        "leaves": [
          "0x28815c9a1c9d638886d6ac193df55f98824c491d09bbbd712f96b5adfeba742e"
        ]
}
```

Edit CSCA list proposal
```json
{
        "@type": "/rarimo.rarimocore.cscalist.EditCSCAListProposal",
        "title": "Add",
        "description": "Lorem ipsum dolor sit amet concestetur!",
        "toAdd": [
          "0x28815c9a1c9d638886d6ac193df55f98824c491d09bbbd712f96b5adfeba742e"
        ],
        "toRemove": []
}
```

## Create proposal

You can create a draft proposal with `rarimo-cored tx gov draft-proposal`.
Select `other`, then choose type `cosmos.gov.v1.MsgExecLegacyContent`.
In `messages[0].content` section, fill `@type` (see proposals above) with module prefix `rarimo.rarimocore.cscalist`.
Depending on the proposal type, fill either `leaves` or `toAdd/toRemove` fields with the list of CSCA public key hashes.

Or you can directly edit this sample (replace `@type` and `leaves` for another type):
```json
{
  "messages": [
    {
      "@type": "/cosmos.gov.v1.MsgExecLegacyContent",
      "content": {
        "@type": "/rarimo.rarimocore.cscalist.ReplaceCSCAListProposal",
        "title": "Init",
        "description": "Lorem ipsum dolor sit amet concestetur!",
        "leaves": [
          "0x28815c9a1c9d638886d6ac193df55f98824c491d09bbbd712f96b5adfeba742e"
        ]
      },
      "authority": "rarimo10d07y265gmmuvt4z0w9aw880jnsr700jw44g3e"
    }
  ],
  "metadata": "ipfs://CID",
  "deposit": "10000000urmo"
}
```

When you have saved your proposal (assuming `draft_proposal.json`), you may submit it, make a deposit for it and vote, then retrieve your proposal.
Here is the script sample for testing:
```bash
# approximate gas limit required to build full list of 382 nodes
rarimo-cored --keyring-backend test --home config/validator tx gov submit-proposal --gas 9999999999 --yes --from "<address>" "draft_proposal.json"
sleep 5
rarimo-cored --keyring-backend test --home config/validator tx gov deposit --gas 9999999 --yes --from "<address>" "<id>" 10000000stake
sleep 5
rarimo-cored --keyring-backend test --home config/validator tx gov vote --yes --gas 9999999 --from "<address>" "<id>" yes
sleep 5
rarimo-cored query gov proposal "<id>"
```
