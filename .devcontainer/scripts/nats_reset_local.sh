#!/bin/bash

# Remove the nsc directory which has the NATS creds
rm -rf .devcontainer/nsc/ .devcontainer/nats/resolver.conf

# Remove the NATS docker containers
sudo docker ps --format json | jq ' select(.Names | contains ("virtual-machine-api")) | .ID' -r | sudo xargs -I% 'docker rm $(docker kill %)'