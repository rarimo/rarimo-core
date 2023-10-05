# Rarimo core

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## [Mainnet information](./docs/common/mainnet/001-mainnet.md)

## Introduction

**rarimo-core** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

Based on Tendermint + Cosmos SDK blockchain core the main goal is to provide validated information about different cross-chain operations.

Documentation:
* [Rarimo core modules business logic](./x/README.md)
* [Bridge contract architecture](./docs/common/contracts/001-contracts.md)
* [Bridging flow overview](./docs/common/bridging/001-bridging.md)
* [Oracles architecture overview](./docs/common/oracles/001-oracles.md)

## Build

### Install

Use
```shell
ignite chain build
```

If it fails in generate proto stage with error on etheremint .proto files just re-execute that command or try to add `--clear-cache` flag.

### Generate proto files

Use
```shell
ignite generate proto-go
```

If it fails in generate proto stage with error on etheremint .proto files remove `./proto/ethermint` package and execute command again.
DO NOT FORGET TO RETURN DELETED PACKAGE.

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)
