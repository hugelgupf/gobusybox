// Copyright 2015-2018 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// makebbmain adds u-root command package imports to an existing main()
// template file.
package main

import (
	"flag"
	"log"
	"os"

	"github.com/u-root/gobusybox/src/pkg/bb/bbinternal"
	"github.com/u-root/gobusybox/src/pkg/uflag"
)

var (
	pkg      = flag.String("template_pkg", "", "Go import package path")
	destDir  = flag.String("dest_dir", "", "Destination directory")
	pkgFiles uflag.Strings
	commands uflag.Strings
)

func init() {
	flag.Var(&pkgFiles, "package_file", "package files")
	flag.Var(&commands, "command", "Go package path for command to import")
}

func main() {
	flag.Parse()

	fset, astp, _, err := bbinternal.ParseAST("main", pkgFiles)
	if err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll(*destDir, 0755); err != nil {
		log.Fatal(err)
	}
	if err := bbinternal.CreateBBMainSource(fset, astp, commands, *destDir); err != nil {
		log.Fatal(err)
	}
}
