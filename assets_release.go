//go:build release
// +build release

package main

import "embed"

//go:embed static/* templates/*
var embedFS embed.FS
var buildMode = "release"
