#!/bin/sh

set -o errexit
set -o nounset
set -o pipefail

if [ ! $(command -v gors) ]
then
    go install github.com/go-leo/gors/cmd/gors@latest
	gors --version
fi

echo "--- go generate start ---"
go generate ./...
echo "--- go generate end ---"

