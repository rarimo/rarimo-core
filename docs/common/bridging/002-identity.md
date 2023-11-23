---
layout: default
title: Identity transfers
---

# Identity transfers

Identity transfers currently implemented in the following way:

- There are an operation types `IDENTITY_GIST_TRANSFER` and `IDENTITY_STATE_TRANSFER` that have to be used fore
  transferring state transitions from one chain to another.

- An [evm-identity-oracle-svc](https://github.com/rarimo/evm-identity-saver-svc) exists, that is responsible
  for delivering information from certain Iden3 state contract into the Rarimo chain.

- After delivering information about state or GIST update, oracles (`evm-identity-saver-svc`) vote for information
  correctness.

- After successful voting, the threshold signature producers provides the ECDSA signature for state update information
  hash.

- There is a [modified state smart contracts](https://github.com/rarimo/polygonid-integration-contracts) that accepts
  state updates with ECDSA signature instead of ZK proof of state update validity.

- Such smart contract should be deployed into every chain that we have to support.

- The information about state update with its witness (ECDSA signature) can be delivered into modified state smart
  contracts on the target chain by everyone.

- Also, the [relayer service](https://github.com/rarimo/identity-relayer-svc) exists, that tracks new state and
  GIST update operations, their signatures and submits if requested transactions to the target chain. For the DApp it
  can
  be used in the following way:
    - PolygonID wallet generates on-chain zkp based on information from Polygon chain state smart contact.
    - After receiving the proof on your application, using corresponding public signals you can trigger the transfer of
      required state and GIST hashes.

----

## Iden3

To explore how the Iden3 protocol works please visit their [documentation](https://docs.iden3.io/).

----

## [Identity saver (oracle)](https://github.com/rarimo/evm-identity-saver-svc)

In the current implementation the identity saver subscribes to the certain EVM chain state smart contract and tracks
the `StateLibStateUpdated` events.

Event definition:

```go
package main

import "math/big"

type StateLibStateUpdated struct {
	Id        *big.Int
	BlockN    *big.Int
	Timestamp *big.Int
	State     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

```

The following configuration ___.yaml___ file should be provided to launch your oracle:

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
  issuer_id: "issuer id in decimal format"
```

After fetching of some event, oracle will create the corresponding transaction and submit it to the Rarimo blockchain.
After transaction appears in Rarimo blockchain all chain oracles have to vote for its correctness.
So they will fetch the information about state update from Rarimo chain, verify it and submit their votes (YES/NO
answers).

For submitting transactions, savers (oracles) uses the [broadcaster-svc](https://github.com/rarimo/broadcaster-svc).
It accepts the transaction by GRPC endpoint, signs it with configured private key and submits it to the Rarimo chain.

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

Explore the simple docker-compose file to run the described services:

```yaml
version: "3.7"

services:

  broadcaster:
    image: ghcr.io/rarimo/broadcaster-svc:v1.0.2
    restart: unless-stopped
    depends_on:
      - broadcaster-db
    volumes:
      - ./config/broadcaster.yaml:/config.yaml
    environment:
      - KV_VIPER_FILE=/config.yaml
    entrypoint: sh -c "broadcaster-svc migrate up && broadcaster-svc run all"

  broadcaster-db:
    image: postgres:13
    restart: unless-stopped
    environment:
      - POSTGRES_USER=broadcaster
      - POSTGRES_PASSWORD=broadcaster
      - POSTGRES_DB=broadcaster
      - PGDATA=/pgdata
    volumes:
      - broadcaster-data:/pgdata

  evm-identity-saver:
    image: ghcr.io/rarimo/evm-identity-saver-svc:v1.0.6
    restart: unless-stopped
    depends_on:
      - broadcaster
    volumes:
      - ./config/evm-saver.yaml:/config.yaml
    environment:
      - KV_VIPER_FILE=/config.yaml
    entrypoint: sh -c "evm-identity-saver-svc run state-update-all"
```

----

## [Relayer service](https://github.com/rarimo/identity-relayer-svc)

Relayer service is responsible for fetching information from Rarimo chain about new signatures
for `IDENTITY_GIST_TRANSFER` and `IDENTITY_STATE_TRANSFER` operations.

The following configuration .yaml file should be provided to launch your raleyer service:

```yaml
log:
  disable_sentry: true
  level: debug

# The port to run on
listener:
  addr: :8000

# PostgreSQL DB connect
db:
  url: "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

# Rarimo core RPCs
core:
  addr: tcp://validator:26657

cosmos:
  addr: validator:9090

evm:
  chains:
    - name: "Ethereum"
      ## Address of the modified state contract on target chain
      contract_address: ""
      ## Private key HEX (without leading 0x) that will pay the tx fee
      submitter_private_key: ""
      ##  RPC address. Example https://mainnet.infura.io/v3/11111
      rpc:
      ## Target chain id
      chain_id: 1

relay:
  # Flag the indicates should service iterate over all existing transfer operation and fill the database
  catchup_disabled: true
```

Explore the simple docker-compose file to run described services:

```yaml
version: "3.7"

services:
  redis:
    image: 'bitnami/redis:latest'
    environment:
      - ALLOW_EMPTY_PASSWORD=yes

  relayer:
    image: registry.gitlab.com/rarimo/polygonid/relayer-svc:v1.0.2
    restart: on-failure
    depends_on:
      - redis
    volumes:
      - ./config/relayer.yaml:/config.yaml
    environment:
      - KV_VIPER_FILE=/config.yaml
    entrypoint: sh -c "relayer-svc run all"
```

----

## Notes

1. The issuer service required to be launched by your own to publish users' claims.

2. Currently, running of your oracle service requires staking of RMO tokens in Rarimo chain.

3. Support of certain smart contract and issuer by oracles requires the voting between all active oracles for that
   chain. It means that for your chain all other oracles have to update their configuration to support your smart
   contract and issuer. Also, the chain-clone can be added to the Rarimo to split oracle groups.