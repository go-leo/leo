//go:build tools
// +build tools

// following https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module

package tools

import (
	_ "github.com/ugorji/go"
	_ "github.com/ugorji/go/codec"
)
