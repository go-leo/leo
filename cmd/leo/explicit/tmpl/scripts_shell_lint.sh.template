#!/bin/sh

set -o errexit
set -o nounset
set -o pipefail

if [ ! $(command -v golangci-lint) ]
then
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh  | sh -s -- -b $(go env GOPATH)/bin
	golangci-lint --version
fi

echo "--- golangci lint start ---"

golangci-lint run -v

echo "--- golangci lint end ---"

