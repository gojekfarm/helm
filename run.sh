#!/bin/bash

export HELM_DEBUG=1
export HELM_REPOSITORY_CONFIG=~/.helm/repository HELM_NAMESPACE=hermes 
export HELM_CONFIG=~/.kube/config
export HELM_NO_PLUGINS=1
make build
ln -F -s ./bin/helm helm3

#./helm3 delete something
#./helm3 install something stable/redis
./helm3 ls -v=7

