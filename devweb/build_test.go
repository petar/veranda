package main

import (
	"testing"
)

func TestBuild(t *testing.T) {
	rootPackage = "github.com/petar/p/blog"
	buildProxy()
}
