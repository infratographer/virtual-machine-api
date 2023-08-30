#!/bin/sh -x

# script to bootstrap a nats operator environment

if nsc describe operator; then
    echo "operator exists, not overwriting config"
    exit 0
fi

echo "Cleaning up NATS environment"
rm -rf /nsc/*

echo "Creating NATS operator"
nsc add operator --generate-signing-key --sys --name LOCAL
nsc edit operator -u 'nats://nats:4222'
nsc list operators
nsc describe operator

export OPERATOR_SIGNING_KEY_ID=`nsc describe operator -J | jq -r '.nats.signing_keys | first'`

echo "Creating NATS account for virtual-machine-api"
nsc add account -n VMAAS -K ${OPERATOR_SIGNING_KEY_ID}
nsc edit account VMAAS --sk generate --js-mem-storage -1 --js-disk-storage -1 --js-streams -1 --js-consumer -1
nsc describe account VMAAS

export ACCOUNTS_SIGNING_KEY_ID=`nsc describe account VMAAS -J | jq -r '.nats.signing_keys | first'`

echo "Creating NATS user for virtual-machine-api"
nsc add user -n USER -K ${ACCOUNTS_SIGNING_KEY_ID}
nsc describe user USER

echo "Generating NATS resolver.conf"
# WORKSPACE_ROOT is possibly nil if this is in the init container, or it's set in the app container
if [ ! -z "$WORKSPACE_ROOT" ]; then
    NATS_ROOT="$WORKSPACE_ROOT/.devcontainer"
    mkdir $NATS_ROOT/nats
fi
nsc generate config --mem-resolver --sys-account SYS --config-file $NATS_ROOT/nats/resolver.conf --force