# rarimocore
**rarimocore** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Install

Use
```shell
ignite chain build
```

If it fails in generate proto stage with error on etheremint .proto files just re-execute that command or try to add `--clear-cache` flag.

Also try use `ignite chain build --skip-proto`

For release build:
```shell
ignite chain build --release -t linux:amd64 --skip-proto
```

## Generate proto files

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
