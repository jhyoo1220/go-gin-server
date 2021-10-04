#!/usr/bin/env bash

set -e

NAME=$1
TAG=$2
AWS_REGION=$3
AWS_ACCOUNT_ID=$4

if [ -z "${NAME}" ]; then
  echo "[ERROR] name is missing!"
  exit 1
fi

if [ -z "${TAG}" ]; then
  echo "[ERROR] tag is missing!"
  exit 1
fi

if [ -z "${AWS_REGION}" ]; then
  echo "[ERROR] AWS region is missing!"
  exit 1
fi

if [ -z "${AWS_ACCOUNT_ID}" ]; then
  echo "[ERROR] AWS account id is missing!"
  exit 1
fi

docker build --compress --rm --tag "${NAME}":"${TAG}" --file ./build/Dockerfile .

$(aws ecr get-login --no-include-email --region "${AWS_REGION}")

docker tag "${NAME}":"${TAG}" "${AWS_ACCOUNT_ID}".dkr.ecr.ap-northeast-2.amazonaws.com/"${NAME}":"${TAG}"
docker push "${AWS_ACCOUNT_ID}".dkr.ecr.ap-northeast-2.amazonaws.com/"${NAME}":"${TAG}"
