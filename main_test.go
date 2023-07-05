package main

import (
	_ "embed"
	"testing"
)

//go:embed go.mod
var goModContent string

func TestAAA(t *testing.T) {
	t.Log("hello world")
	t.Log(goModContent)
}
