#!/bin/bash

KEY="mykey"
CHAINID="test-1"
MONIKER="localtestnet"
KEYALGO="secp256k1"
KEYRING="test"
LOGLEVEL="info"
# to trace evm
#TRACE="--trace"
TRACE=""

# remove existing daemon
rm -rf ~/.gravity*

# if $KEY exists it should be deleted
echo "decorate bright ozone fork gallery riot bus exhaust worth way bone indoor calm squirrel merry zero scheme cotton until shop any excess stage laundry" | gravity keys add $KEY --keyring-backend $KEYRING --algo $KEYALGO --recover

gravity init $KEY --chain-id $CHAINID 

# Allocate genesis accounts (cosmos formatted addresses)
gravity add-genesis-account $KEY 1000000000000stake --keyring-backend $KEYRING

# Sign genesis transaction
gravity gentx $KEY 1000000stake 0x1d65BCC107689Fb9c35Ae403d028E29C1C90C36A gravity18ytfr4s8lfccy048zl00y3akujxqvq75sfpuzq --keyring-backend $KEYRING --chain-id $CHAINID

# Collect genesis tx
gravity collect-gentxs 


if [[ $1 == "pending" ]]; then
  echo "pending mode is on, please wait for the first block committed."
fi

# update request max size so that we can upload the light client
# '' -e is a must have params on mac, if use linux please delete before run
sed -i'' -e 's/max_body_bytes = /max_body_bytes = 1/g' ~/.centauri/config/config.toml

# Start the node (remove the --pruning=nothing flag if historical queries are not needed)
gravity start --pruning=nothing  --minimum-gas-prices=0.0001stake --rpc.laddr tcp://0.0.0.0:26657