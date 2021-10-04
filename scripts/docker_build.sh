#!/usr/bin/env bash

set -e

NAME=$1
TAG=$2

if [ -z "${NAME}" ]; then
  echo "[ERROR] name is missing!"
  exit 1
fi

if [ -z "${TAG}" ]; then
  echo "[ERROR] tag is missing!"
  exit 1
fi

docker build --compress --rm --tag "${NAME}":"${TAG}" --file ./build/Dockerfile .
