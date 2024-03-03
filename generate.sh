#!/bin/bash

export GO_POST_PROCESS_FILE="/usr/local/bin/gofmt -w"

OPEN_API_FILE="https://www.cysend.com/openapi/doc/cysend_v050203_prod.json"

openapi-generator-cli validate -i "${OPEN_API_FILE}"
openapi-generator-cli generate \
    -i "${OPEN_API_FILE}" \
    -g go \
    -p packageName=cysend \
    -p enumClassPrefix=true \
    --git-host github.com \
    --git-repo-id cysend-go \
    --git-user-id pourcheragh
