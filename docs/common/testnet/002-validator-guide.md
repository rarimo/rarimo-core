---
layout: default
title: Validators' Guide
---

# Validators' Guide

This instruction tells how to start one or another system service.
It is assumed that you are already familiar with Cosmos SDK, Linux and are fluent in them.

## Node (rarimo-core)

Chain id: `rarimo_201411-2`

### Build

Clone the repository "<https://github.com/rarimo/rarimo-core.git>" and checkout the `chain/mainnet-beta` branch. 

Use the following commit hash to build: `20bcfef13ce6d3eaa77a63fa8f1d141601c27ba5`.

You can use [Dockerfile.vendor](../../../Dockerfile.vendor), Ignite CLI (do not use vendor package in that case)
or just a ```go build ./cmd/rarimo-cored```.

### Setup configs

To, generate validator configs, use:
- `rarimo-core` (or `rarimo-cored`) executable
- RPC from [Testnet page](./001-testnet.md)

Also use the following P2P connection to our RPC node: `f11612cdb2d705eaf7d0820f3e125a1d0ad9e225@104.196.227.66:26656`.

Init folder structure:
```bash
rarimo-core init $MONIKER_NAME --chain-id=rarimo_201411-2 --home=$RARIMO_HOME
```

Paste custom `genesis.json` and `app.toml` in `$RARIMO_HOME/config/` folder.

Create validator private key:
```bash
rarimo-core keys add <key-name> --keyring-backend test --home=$RARIMO_HOME
```

To run use env variables:
```yaml
- name: DAEMON_NAME
value: "rarimo-core"

- name: DAEMON_HOME
value: $RARIMO_HOME

- name: DAEMON_ALLOW_DOWNLOAD_BINARIES
value: "true"
```

Use the following command to start your node:
```bash
mkdir -p $DAEMON_HOME/cosmovisor/genesis/bin && cp YOU_STORE_CORE_BIN(name rarimo-core) $DAEMON_HOME/cosmovisor/genesis/bin && cosmovisor run start --home=$RARIMO_HOME --rpc.laddr tcp://0.0.0.0:26657
```

To stake the tokens for new validator execute:

```bash
rarimo-core tx staking create-validator --amount 1000000urmo --commission-max-change-rate "0.01"  --commission-max-rate "0.2" --commission-rate "0.1" --min-self-delegation "1" --details "Meet new Rarimo validator" --pubkey $(rarimo-core tendermint show-validator --home=$RARIMO_HOME) --moniker $MONIKER_NAME --chain-id rarimo_201411-1 --fees 0urmo --from $LOCAL_VALIDATOR_ADDRESS --home=$RARIMO_HOME --node=$RARIMO_NODE --keyring-backend=test
```

Congratulations, you are the Rarimo Mainnet validator!

----

### Useful commands

Query delegator rewards for all validators:
```shell
rarimo-core query distribution rewards [delegator address rarimo...] --node=https://rpc.mainnet.rarimo.com:443
```

or for certain validator:
```shell
rarimo-core query distribution rewards [delegator address rarimo...] [valdidator address rarimovaloper...] --node=https://rpc.mainnet.rarimo.com:443
```

Query un-withdrawn rewards:
```shell
rarimo-core query distribution validator-outstanding-rewards [valdidator address rarimovaloper...] --node=https://rpc.mainnet.rarimo.com:443
```

Query validator commission:
```shell
rarimo-core query distribution commission rarimovaloper... --node=https://rpc.mainnet.rarimo.com:443
```

Withdraw all validator rewards for a delegator:
```shell
rarimo-core tx distribution withdraw-all-rewards --from=rarimo... --node=https://rpc.mainnet.rarimo.com:443 --home=path-to-home-with-keyring
```

Withdraw certain validator rewards for a delegator (add `--commission` flag to withdraw commissions also):
```shell
rarimo-core tx distribution withdraw-rewards rarimovaloper1... --from rarimo... --commission --node=https://rpc.mainnet.rarimo.com:443 --home=path-to-home-with-keyring
```

Check validator seed:
```bash
rarimo-core tendermint show-node-id --home=$RARIMO_HOME
```

Check key exists in keystore:
```bash
rarimo-core keys show $LOCAL_VALIDATOR_ADDRESS --keyring-backend test --home=$RARIMO_HOME
```