// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/petar/p/devweb/slave"
	_ "github.com/petar/p/blog/post"
)

func main() {
	slave.Main()
}
