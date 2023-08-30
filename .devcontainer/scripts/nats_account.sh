#!/bin/bash


sudo chown -Rh vscode:vscode $WORKSPACE_ROOT/.devcontainer/nsc

echo "Dumping NATS user creds file"
nsc --data-dir=$WORKSPACE_ROOT/.devcontainer/nsc/nats/nsc/stores generate creds -a VMAAS -n USER > /tmp/user.creds

echo "Dumping NATS sys creds file"
nsc --data-dir=$WORKSPACE_ROOT/.devcontainer/nsc/nats/nsc/stores generate creds -a SYS -n sys > /tmp/sys.creds

