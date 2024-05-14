---
layout: default
title: List of Modules
---

# List of Modules

* [x/rarimocore](./rarimocore/README.md) - Base transfer logic (signing, storing operations, tss parties staking).
* [x/tokenmanager](./tokenmanager/README.md) - Managing supported tokens and chains.
* [x/bridge](./bridge/README.md) - Bridging native token from/to Rarimo chain.
* [x/multisig](./multisig/README.md) - Multisig module (Cosmos `x/group` module with extended functionality: can execute any message and hold tokens like a simple account).
* [x/oraclemanager](./oraclemanager/README.md) - Managing oracles that submit events to Rarimo chain.
* [x/evm and x/feemarket](./evm/README.md) - EVM compatibility feature.
* [x/identity](./identity/README.md) - Storing aggregated identity data.
* [x/cscalist](./cscalist/README.md) - Storing and managing CSCA Master List

## EVM

Read more about writing smart contracts with solidity at the official [`evm` documentation page](https://docs.evmos.org/modules/evm/).