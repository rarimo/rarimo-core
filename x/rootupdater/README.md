---
layout: default
title: x/rootupdater
---

# `x/rootupdater`

## Abstract

This module is designed to extract the passport root from Ethereum Virtual Machine (EVM) contract events and 
subsequently create a new operation for `rarimo-core` module .

----

## Concepts

The core functionality of this module involves monitoring and processing events emitted by EVM contracts
to capture the passport root. Once the passport root is accurately extracted, the module proceeds to create
a new operation within the rarimo-core module.

An operation includes the following fields, such as `contract_address`, `root`, `root_timestamp`, `block_height` 

## State

Model can be found in `proto/rootupdater`. Module data are stored inside Params

####  Module params proto

```protobuf
message Params {
  string contract_address = 1;
  string root = 2;
  string last_signed_root = 3;
  string last_signed_root_index = 4;
  string event_name = 5;
  int64 root_timestamp = 6;
  uint64 block_height = 7;
}
```

Example values are provided in JSON for convenient reading.

####  Module params
```json
{
  "params": {
    "contract_address": "0x65D51e50453371392b4c1280BE9B75Cbe52F950e",
    "root": "0x000000",
    "last_signed_root": "0x000000",
    "last_signed_root_index": "0x1c710a1e732c4178b110e01096f26c83c81a5a9118d8219ef8d204954ca09dd9",
    "event_name": "RootUpdated",
    "root_timestamp": "1724316208",
    "block_height": "0"
  }
}
```