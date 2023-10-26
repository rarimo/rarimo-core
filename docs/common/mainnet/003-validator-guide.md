---
layout: default
title: Validators' Guide
---

# Validators' Guide

This instruction tells how to start one or another system service.
It is assumed that you are already familiar with Linux and are fluent in it

Service start order:
1. node (rarimo-core)
2. broadcaster-svc
3. evm-identity-saver-svc
4. tss-svc

## Node (rarimo-core)

Chain id: `rarimo_201411-1`

Genesis file: "<https://storage.cloud.google.com/rarimo-mainnet/core/genesis.json>"

Address book: "<https://storage.googleapis.com/rarimo-mainnet/core/addrbook.json>"

APP file: "<https://storage.googleapis.com/rarimo-mainnet/core/app.toml>"

Rarimo core binary (linux/amd64): "<https://storage.googleapis.com/rarimo-mainnet/core/rarimo-core>"

To, generate validator configs, use:
- `rarimo-core` file in link.
- node ip connect: `34.66.205.183`. Use for submitting TX as well as for p2p connection.

Also use the following P2P connection to our RPC node: `39072210913bb52a6444543ad160b04668d58210@34.66.205.183:26656`

Execute te following scripts:

Set Envs
```bash
export MONIKER_NAME=YOU_VALIDATOR_NAME
export RARIMO_HOME=YOU_RARIMO_HOME_PATH
export RARIMO_NODE=tcp://34.66.205.183:26657
```

Init folder structure:
```bash
rarimo-core init $MONIKER_NAME --chain-id=rarimo_201411-1 --home=$RARIMO_HOME
```

Paste custom `genesis.json` and `app.toml` in `$RARIMO_HOME/config/` folder.

Create validator private key:
```bash
rarimo-core keys add <key-name> --keyring-backend test --home=$RARIMO_HOME
```

Save your mnemonic and address. That address will be used for your validator staking and also for your oracles.

Send you address (rarimo...) to `yp@distributedlab.com` or `vl@distributedlab.com`.

To setup broadcaster service later you will need a private key. Use your mnemonic to get ECDSA private key (0x...):
```bash
rarimo-core tx rarimocore parse-mnemonic 'mnemonic phrase'
```

Also, add environment var with created address:
```bash
export LOCAL_VALIDATOR_ADDRESS=rarimo...
```

Please, backup the following files and folders:
```bash
$RARIMO_HOME/config/priv_validator_key.json
$RARIMO_HOME/config/node_key.json
$RARIMO_HOME/keyring-test
```

Check validator seed:
```bash
rarimo-core tendermint show-node-id --home=$RARIMO_HOME
```

Check key exists in keystore:
```bash
rarimo-core keys show $LOCAL_VALIDATOR_ADDRESS --keyring-backend test --home=$RARIMO_HOME
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

----

## Broadcaster

Broadcaster service binary (linux/amd64): "<https://storage.googleapis.com/rarimo-mainnet/modules/broadcaster-svc>"

To start `broadcaster-svc` service you need a hex private key that will be used to sign messages.

The following configuration `config.yaml` file should be provided to launch your broadcaster service:

```yaml
log:
  disable_sentry: true
  level: debug

## Port to listen incoming requests
listener:
  addr: :80

## PostgreSQL database connect
db:
  url: "postgres://broadcaster:broadcaster@broadcaster-db/broadcaster?sslmode=disable"

key:
  ## Rarimo account private key in Hex format: 0x123fab...
  sender_prv_hex: ""
  ## Rarimo chain id
  chain_id: "rarimo_201411-1"
  ## Base coin name
  coin_name: "urmo"

cosmos:
  addr: "validator:9090"
```

You will also need some environment variables to run:
```yaml
- name: KV_VIPER_FILE
  value: /config/config.yaml # The path to your config file
```

The execution of 2 following commands is required for launch:

1. To perform migrations
```bash
broadcaster-svc migrate up
```

2. To start the service
```bash
broadcaster-svc run all
```

Also, you can run these commands together like this:
```bash
broadcaster-svc migrate up && broadcaster-svc run all
```

----

## EVM identity saver

EVM identity saver service binary (linux/amd64): "<https://storage.googleapis.com/rarimo-mainnet/1.0.4/evm-identity-saver-svc>"

Currently, it should be one instance per account only for Polygon.

The following configuration `config.yaml` file should be provided to launch your oracle:

```yaml
log:
  disable_sentry: true
  level: debug

## Port to listen incoming requests (used to trigger revote for some operation - rare flow)
listener:
  addr: :8000

evm:
  ## State contract for listening
  contract_addr: "0x624ce98D2d27b20b8f8d521723Df8fC4db71D79D"
  ## Polygon websocket address
  rpc: ""
  ## Start block to listen events. (0 - listen from current). Used to catchup old events. Be careful to use.
  start_from_block: 0
  ## According to network name in Rarimo core. Example: Polygon
  network_name: Polygon
  ## How many blocks should be created after event to fetch it.
  block_window: 20

broadcaster:
  ## address of broadcaster-svc
  addr: ""
  ## sender account address (rarimo..)
  sender_account: ""

## Rarimo chain connections

core:
  addr: tcp://validator:26657

cosmos:
  addr: validator:9090

## Profiling

profiler:
  enabled: false
  addr: :8080

## Issuer information

state_contract_cfg:
  issuer_id: ['']
  disable_filtration: true
```

Also, some environment variables will be needed to run
```yaml
- name: KV_VIPER_FILE
  value: /config/config.yaml # is the path to your config file
```

Oracle service requires staking of some RMO tokens in Rarimo chain.
Please, **do not start** the service before we confirm your oracle staking.

To start the service (in vote mode) use the following command:
```bash
evm-identity-saver-svc run state-update-voter
```

----

## TSS

TSS service binary (linux/amd64): "<https://storage.googleapis.com/rarimo-mainnet/1.0.4/tss-svc>"

To become an active TSS you need to follow that steps:

1. Generate TSS account (it is recommended to generate the new one and do not use it in other services):

  ```shell
  rarimo-core keys add <key-name> --keyring-backend test --home=$TSS_HOME
  ```

  Also, you need to parse mnemonic to get corresponding private key:

  ```shell
  rarimo-core tx rarimocore parse-mnemonic 'mnemonic phrase'
  ```

2. Pre-setup secret parameters, execute and save response JSON:

  ```shell
  tss-svc run paramgen
  ```

3. Generate trial ECDSA private key, execute and store results:

  ```shell
  tss-svc run prvgen
  ```

4. Setup the Vault service and create secret for your tss (type KV version 2). Secret should contain the following credentials:

  * "data": "Leave empty"

  * "pre": "Generated pre params JSON"

  * "account": "Your Rarimo account hex key"

  * "trial": "Generated Trial ECDSA private key hex"

  JSON example:

  ```json
  {
    "tls": true,
    "data": "",
    "pre": "pre-generated-secret-data",
    "account": "rarimo-account-private-key-hex-leading-0x",
    "trial": "trial-ecdsa-private-key-hex-leading-0x"
  }
  ```

5. Create a configuration file `config.yaml` with the following structure:

  ```yaml
  log:
    disable_sentry: true
    level: debug

  ## PostreSQL connection

  db:
    url: "postgres://tss:tss@tss-2-db:5432/tss?sslmode=disable"

  ## Port to listen for incoming GRPC requests

  listener:
    addr: :9000

  ## Core connections

  core:
    addr: tcp://validator:26657

  cosmos:
    addr: validator:9090

  ## Session configuration (should be the same for all services accross the system)

  session:
    start_block: 100
    start_session_id: 0

  ## Swagger doc configuration

  swagger:
    addr: :1313
    enabled: false

  ## Chain configuration

  chain:
    chain_id: "rarimo_201411-1"
    coin_name: "urmo"
  ```

  Set up host environment:

  ```yaml
    - name: KV_VIPER_FILE
    value: /config/config.yaml # is the path to your config file
    - name: VAULT_PATH
    value: http://vault-internal:8200 # your vault endpoint
    - name: VAULT_TOKEN
    value: "" # your vault token ("root"/"read/write")
    - name: MOUNT_PATH
    value: secret
    - name: SECRET_PATH
    value: tss # name of the secret path vault (type KV version 2)
  ```

6. To launch the service use the following command.

  ```shell
  tss-svc migrate up && tss-svc run service
  ```

7. After launching of your service, share your tss-svc URL, rarimo account address and ECDSA public key with `yp@distributedlab.com` or `vl@distributedlab.com`.

  Note, that your TSS service should be accessible only using secure TLS connection.

  After some period your TSS will generate new keys with other active parties and become an active party.

----

## Finishing with validator

After receiving confirmation from ourside about token transfer, execute command to stake tokens and become a validator.
You need to stake exactly 1000000urmo that is equal to 1 RMO.

```bash
rarimo-core tx staking create-validator --amount 1000000urmo --commission-max-change-rate "0.01"  --commission-max-rate "0.2" --commission-rate "0.1" --min-self-delegation "1" --details "Meet new Rarimo validator" --pubkey $(rarimo-core tendermint show-validator --home=$RARIMO_HOME) --moniker $MONIKER_NAME --chain-id rarimo_201411-1 --fees 0urmo --from $LOCAL_VALIDATOR_ADDRESS --home=$RARIMO_HOME --node=$RARIMO_NODE --keyring-backend=test
```

Congratulations, you are the Rarimo Mainnet validator!

----

## Useful commands

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
