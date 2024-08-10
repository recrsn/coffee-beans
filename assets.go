//go:build !release
// +build !release

package main

import "embed"

var embedFS embed.FS
var buildMode = "debug"
