# Running saver (oracle) service

## Broadcaster service

The broadcaster service should be the only one entrypoint for submitting transactions from certain rarimo account.
All your oracle services that uses the same account should submit transaction though the same roadcaster service.

To start `broadcaster-svc` service you need a hex private key that will be used to sign messages.

The following configuration .yaml file should be provided to launch your broadcaster service:
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
  chain_id: "rarimo_201411-2"
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

There are 2 commands execution required for launch:

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

## evm-identity-saver-svc

The following configuration .yaml file should be provided to launch your oracle:

```yaml
log:
  disable_sentry: true
  level: debug

## Port to listen incoming requests (used to trigger revote for some operation - rare flow)
listener:
  addr: :8000

evm:
  ## State contract for listening
  contract_addr: ""
  ## Websocket address
  rpc: ""
  ## Start block to listen events. (0 - listen from current). Used to catchup old events. Be careful to use.
  start_from_block:
  ## According to network name in Rarimo core. Example: Polygon
  network_name:
  ## How many blocks should be created after event to fetch it.
  block_window:

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
  issuer_id: ['', '']
  disable_filtration: true
```

Also, some environment variables will be needed to run
```yaml
-name: KV_VIPER_FILE
   value: /config/config.yaml # is the path to your config file
```

To start the service (in vote mode) use the following command:
```bash
evm-identity-saver-svc run state-update-voter
```

To run in full mode:
```bash
evm-identity-saver-svc run state-update-all
```

## Staking

To stake for the Oracle use the following CLI command: 

```shell
export ACCOUNT_HOME=/path/to/your/account/prv/data
export ACCOUNT_ADDR=rarimo...
rarimo-core tx oraclemanager stake [Network name, example: Solana] [Stake amount (urmo), example: 1000000] --chain-id rarimo --fees 0urmo --from=$ACCOUNT_ADDR --home=$ACCOUNT_HOME --keyring-backend=test
```
