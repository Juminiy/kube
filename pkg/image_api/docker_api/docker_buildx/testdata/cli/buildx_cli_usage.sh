#!/bin/bash

# create a buildx instance as buildx build --builder
docker buildx create --name multiple-x --driver docker-container --bootstrap --use

# equal buildx build args
# 1.
# --load
# --output=type=docker
# 2.
# --push
# --output=type=registry

docker buildx build \
    --build-arg k1=v1 \
    --build-arg k2=v2 \
    --builder multiple-x \
    --debug \
    --file Dockerfile \
    --label labelk1=labelv1 \
    --label l1belk2=labelv2 \
    --load \
    --platform linux/amd64,linux/arm64,linux/arm64v7 \
    --tag refStr1 \
    --tag refStr2 \
    .

