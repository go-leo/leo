#!/bin/sh

set -o errexit
set -o nounset
set -o pipefail

if [ ! $(command -v gofumpt) ]
then
	go install mvdan.cc/gofumpt@latest
	gofumpt -version
fi

echo "--- go format start ---"
gofumpt -w .
echo "--- go format end ---"

