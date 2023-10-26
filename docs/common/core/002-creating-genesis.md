---
layout: default
title: Setup local genesis state
---

# Setup local genesis state 

### Creating validator key
```shell
rarimo-cored keys add validator_key --keyring-backend test --home=./genesis
```

### Exporting validator address
```shell
export MY_VALIDATOR_ADDRESS=rarimo10efu3md78z8qhlzjsx8u0kq3p7j3uhhf80yp5u
```

### Init chain
```shell
rarimo-cored init main --chain-id rarimo-core --home=./genesis
```

### Adding account to genesis
```shell
rarimo-cored add-genesis-account $MY_VALIDATOR_ADDRESS 1000000000000stake --home=./genesis
```

### Generating add-validator transaction
```shell
rarimo-cored gentx validator_key 10000000000stake --chain-id rarimo-core --keyring-backend test --home=./genesis
```

### Importing transaction into genesis file
```shell
rarimo-cored collect-gentxs --home=./genesis
```
