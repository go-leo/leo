#!/bin/sh

set -o errexit
set -o nounset
set -o pipefail

if [ ! $(command -v wire) ]
then
	go install github.com/google/wire/cmd/wire@latest
	wire help
fi

echo "--- wire generate start ---"
wire ./...
echo "--- wire generate end ---"

