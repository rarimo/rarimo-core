---
layout: default
title: Mainnet upgrades
---

# Mainnet upgrades

## V1.0.7

Core binary: (alpine-linux/amd64): "<https://storage.googleapis.com/rarimo-mainnet/1.0.7/rarimo-core>".

Also, you can build core from sources by yourself: Use <"https://github.com/rarimo/rarimo-core"> repo
and `chains/mainnet` branch.

Upgrade will perform automatically if you are using `cosmovisor` under Alpine linux.
Also, if you are using Ubuntu linux, please install `musl-dev` using `sudo apt install musl-dev` command to be able to
use Alpine binary on your machine.

Upgrade v1.0.7 introduces fixes to the `rarimocore` module:

- Adding feature to clear old TSS violation reports.
- Manual unfreeze of all TSS parties without necessity to reshare keys.

## V1.0.6

Core binary: (alpine-linux/amd64): "<https://storage.googleapis.com/rarimo-mainnet/1.0.6/rarimo-core>"

Upgrade will perform automatically if you are using `cosmovisor` under Alpine linux.
Also, if you are using Ubuntu linux, please install `musl-dev` using `sudo apt install musl-dev` command to be able to
use Alpine binary on your machine.

Upgrade V1.0.6 introduces:

- Adding message for operation resign by Threshold signature producers.
- Changing of stored TSS pub-key format: removing constant 0x04 prefix.
- Adding `admin` field in supported network BridgeParams that will contain the pub-key of the bridge admin on Solana.
- Some minor fixes in: several operation creation entrypoints, updating params of the supported networks,

## V1.0.5

Core binary: (alpine-linux/amd64): "<https://storage.googleapis.com/rarimo-mainnet/1.0.5/rarimo-core>"

Upgrade will perform automatically if you are using `cosmovisor` under Alpine linux.
Also, if you are using Ubuntu linux, please install `musl-dev` using `sudo apt install musl-dev` command to be able to
use Alpine binary on your machine.

Upgrade v1.0.5 introduces several fixes:

- Adding `createdAtBlock` field to `StateInfo` entry in `identity` module.
- Fixed events emitting for multisig messages execution.

## V1.0.4

Core binary (linux/amd64): "<https://storage.googleapis.com/rarimo-mainnet/1.0.4/rarimo-core>"

Upgrade will perform automatically if you are using `cosmovisor` under Alpine linux.
Also, if you are using Ubuntu linux, please install `musl-dev` using `sudo apt install musl-dev` command to be able to
use Alpine binary on your machine.

Upgrade v1.0.4 introduces `identity` core module that is responsible for storing aggregated information about
identity state transitions published into Rarimo chain. It uses deployed into Rarimo EVM original Iden3 state
smart-contracts.

Also, there was a couple of fixes in Rarimo evm part and other modules.

`rarimocore` and `oraclemanager` modules genesis `freezedBlock` parameters was set to `103680` that is equal to 2x `gov`
voting time.

`feemanager` module `baseFee` parameter was set to 0 to disable fee for EVM transactions.

Also, we introduce new version of the following services:

- tss-svc:v1.0.4

  Link:"<https://storage.googleapis.com/rarimo-mainnet/1.0.4/tss-svc>"
- evm-identity-saver-svc:v1.0.4

  Link: "<https://storage.googleapis.com/rarimo-mainnet/1.0.4/evm-identity-saver-svc>"

Note, that there is a couple of changes in `evm-identity-saver-svc` config file:

Old:

```yaml
state_contract_cfg:
  issuer_id: "24681524151353338075533135666067797260271579497198610066639546696060309762"
```

New:

```yaml
state_contract_cfg:
  issuer_id: [ '' ]
  disable_filtration: true
```

Also, there are some updates in Rarimo node configuration:

config.toml:

```yaml
max_subscriptions_per_client = 50
```

----

## v1.0.1

Mainnet launched!
