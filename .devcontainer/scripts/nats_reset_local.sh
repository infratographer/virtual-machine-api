#!/bin/bash

# Remove the nsc directory which has the NATS creds
rm -rf .devcontainer/nsc/ .devcontainer/nats/resolver.conf

# Remove the NATS docker containers
for ID in $(sudo docker ps --format json | jq ' select(.Names | contains ("virtual-machine-api")) | .ID' -r); do
  sudo docker kill $ID
  sudo docker rm $ID
done