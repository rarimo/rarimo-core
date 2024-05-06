---
layout: default
title: CSCA List proposal guide
---


# CSCA List proposal guide

## Overview

This guide is focused on working with `rarimo-cored` CLI app.

CSCA Master List is a list of root certificates of passport signing authorities.
The list itself is provided by [ICAO](https://www.icao.int/Pages/default.aspx) and contains
root certificates of passport issuing authorities for a lot of countries.
When obtained the passport signer's root certificate, you can check its presence in the Merkle tree,
stored on Rarimo chain in a decentralized manner.
This is how you ensure that the passport is signed by an eligible entity.

Technical details can be found in [CSCA List Rarimo Core module](../../../x/cscalist/README.md).

The keys of CSCA signers can be compromised or updated, ergo list updates must be handled.
This is done via proposals. The guide will help you in creating and verifying proposals on list changes.

## Pre-requisites

CLI is required in order to convert CSCA List from LDIF to leaves, which are sent to proposal.
You can also verify someone's CSCA list proposals with it to prevent fraud.

1. Install Go [https://go.dev/doc/install](https://go.dev/doc/install)
2. Install `rarimo-cored` CLI app:
```bash
go install github.com/rarimo/rarimo-core/cmd/rarimo-cored
```
This should install the `rarimo-cored` binary into your `$GOPATH/bin` directory.
You can call the CLI from that dir, or add it to your `$PATH`:
```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

All commands require Rarimo home directory set.
Also, some of them require network RPC in order to fetch state.
You can put these into environment variable for convenient command calls:
```bash
export RARIMO_HOME=/path/to/your/rarimo/home
export RPC=https://your-rpc
```

## Monitor changes

The proposal mechanism was implemented due to inability of automatic fetch of CSCA list.
This is why you have to monitor the list manually if you are interested in up-to-date storage.

### Download the list
1. Go to [ICAO website](https://pkddownloadsg.icao.int/)
2. Accept terms and pass captcha below
3. Find "The latest collection of CSCA Master Lists" on the page and download it
4. You should have a `.ldif` file in your downloads with name like `icaopkd-002-complete-*.ldif`

### Compare with state

```bash
rarimo-cored --home $RARIMO_HOME --node $RPC query cscalist ldif-tree-diff your-file.ldif
```
If roots are different, the list was changed and the proposal must be created to update the tree.

## Create proposal

Most likely, you have already filled basic proposal data in the block scan (title, description, etc.). You should have the file with CSCA list, see **Download the list** subsection. Assuming the file name is `your-file.ldif`.

Next, obtain primary data for your proposal, which depends on the proposal type.

### ReplaceCSCAListProposal

```bash
cd ~/Downloads # or wherever you have the file
rarimo-cored --home $RARIMO_HOME query cscalist parse-ldif your-file.ldif your-output-file.txt
```

`your-output-file.txt` will contain the big list of hashes of public keys from the LDIF file. Example:
```
"0x1d4dd579478a38c00f58a4d94263ff2bb0459992c073ebb7a6991194e44157f2",
"0x0dcc4019fccc7ad4fbb535a40633cc32f99a18096a736b21e695e35e964209ae",
"0x01a3d79b678d79a8f912b693c4d57b38cf0e44ef413b7684e92a664e98c911ed"
```
These are the hashes of compatible public keys from the provided LDIF file.
Just copy them to the `leaves` field of your draft proposal.

### EditCSCAListProposal

```bash
cd ~/Downloads # or wherever you have the file
rarimo-cored --home $RARIMO_HOME query cscalist prepare-proposal your-file.ldif your-output-file.txt
```

`your-output-file.txt` will contain two lists of hashes to add and remove. Example:
```
=== To add ===
"0x1d4dd579478a38c00f58a4d94263ff2bb0459992c073ebb7a6991194e44157f2",
"0x0dcc4019fccc7ad4fbb535a40633cc32f99a18096a736b21e695e35e964209ae"
=== To remove ===
"0x01a3d79b678d79a8f912b693c4d57b38cf0e44ef413b7684e92a664e98c911ed"
```
The command has fetched the current state, compared it to the tree from file and provided the differences to consume.
Copy the respective hashes to the `toAdd` and `toRemove` fields of your draft proposal.

### Submit proposal

Basically, you need a command:
```bash
rarimo-cored --home $RARIMO_HOME --node $RPC tx gov submit-proposal draft_proposal.json --from <your-rarimo-account>
```

Refer to Cosmos SDK documentation and other guides to get more details about:
- Account creation
- Generating and managing your keys
- Obtaining `draft_proposal.json`

## Verify proposal

In order to vote for the proposal, you need to ensure that its content corresponds to the actual state of CSCA List.

Firstly, download the list from [ICAO website](https://pkddownloadsg.icao.int/)
and prepare the list of hashes, like described in **Create proposal** section.
Secondly, check whether the proposal makes any difference to the current state:
```bash
rarimo-cored --home $RARIMO_HOME --node $RPC query cscalist ldif-tree-diff your-file.ldif
```
If there is no difference, the proposal either does not make sense, or is fraudulent, and you should vote 'No' or 'NoWithVeto'.
Again, the deeper validation depends on proposal type.

If this is `ReplaceCSCAListProposal`, just compare the `leaves` field with `your-output-file.txt`. It is possible that you could have the different order of hashes due to ICAO change of order. Therefore, you may wish to sort both lists, e.g. with `sort` command.

For `EditCSCAListProposal`, you should additionally perform:
```bash
rarimo-cored --home $RARIMO_HOME --node $RPC query cscalist prepare-proposal your-file.ldif your-output-file.txt
```
Then compare the `toAdd` and `toRemove` fields with the differences between `your-output-file.txt` and someone's proposal.

If you get the same results as in the proposal (of any type), it is correct, and you should support it.
