---
layout: default
title: Commission overview
---

# Commission overview

The bridge feature is the main component of Rarimo cross-chain messaging protocol.
To maintain successful work and reward validators and other people that support the core system every bridge
should charge protocol usage fee.

Protocol fee in Rarimo will be charged for every type of cross-chain token transfer:
- Native
- FT
- and NFT transfers.

For every transfer we will charge a constant amount on start and prepare ground for percentage charging in the future.

----

## Architecture

In parameters, fee smart contract should contain the acceptable token addresses list,
amount for every token to be charged and bridge contract address. Those parameters should be managed by threshold signature.
To check the signature fee contract should execute a cross-contract call to the bridge contract to fetch the threshold public key.

----

## Deposit

For the token deposit, fee smart-contract should accept the same parameters as bridge contract and also contract
address of selected fee token. Deposit method should check if the provided token address is legal.
If everything is well, deposit method performs charging of the corresponding fee amount and executes
deposit method on the bridge contract.


Lets described flow more accurately:
1. User selects the fee token and calls the deposit method on the fee contract.
2. Fee contract checks if the token is available to be charged fee in, and gets the fee amount.
3. Fee contract charges a fee amount from the user. (If it is EVM or Near, the approval token method should be called before by the user).
4. If charging was successful, the fee contract calls a bridge contract with provided deposit arguments.

Also bridge smart-contract should be extended to receive deposit calls only from corresponding fee smart-contract.
Fee smart-contract address also should be stored on the bridge contract and be changeable by signed call.

----

## Collected fee withdrawal

The withdrawal method can be referred to the management methods, but we will describe its logic separately:
Withdrawal methods should be also managed by core proposals and threshold signature.
The public key should be taken from the bridge contract state.

----

## Management

For managing fee tokens, fee smart-contract should implement the **AddFeeToken**, **UpdateFeeToken**, **RemoveFeeToken**,
**Withdraw** methods that accept token information and signature for corresponding hash.

Let’s define the operation type enum with the following fields:
```
0 => AddFeeToken
1 => RemoveFeeToken
2 => UpdateFeeToken
3 => Withdraw
```

Let’s described flow more accurately:

1. Core provides signature and data to be changed on the fee contract.
  For the AddFeeToken, UpdateFeeToken and RemoveFeeToken operation it should be:
  __HASH(
  nonce 32 byte,
  contract addr,
  network name,
  operation type,
  token addr if non-native otherwise none (in evm 20 zero bytes for native),
  amount
  )__

2. For the withdrawal operation it should be:
  __HASH(
  nonce 32 byte,
  receiver address,
  contract addr,
  network name,
  operation type,
  token addr if non-native otherwise none (in evm 20 zero bytes for native),
  amount
  )__

3. Someone submits that information to the fee smart-contract.

4. Fee contract calculates hash of provided data and checks that signature public key corresponds to the key on bridge smart contract.

5. If all checks were correct, fee smart-contract updates state according to the provided information.

----

## Conclusion

On different networks the smart contracts architecture can be may differ from what is described cause of chain smart contracts peculiarities.
Explore every contract by yourself to gent better understanding.
