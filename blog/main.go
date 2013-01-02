// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/petar/veranda/devweb/slave"
	_ "github.com/petar/veranda/blog/post"
	//_"code.google.com/veranda/rsc/appfs/server"
)

func main() {
	slave.Main()
}
