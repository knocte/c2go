// Copyright 2014 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"rsc.io/cc"
)

var (
	inc = flag.String("I", "", "include directory")
)

func main() {
	log.SetFlags(0)
	flag.Parse()
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: c2go [options] *.c\n")
		flag.PrintDefaults()
		os.Exit(2)
	}

	if *inc != "" {
		cc.AddInclude(*inc)
	}

	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
	}

	var r []io.Reader
	files := args
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		r = append(r, f)
		defer f.Close()
	}
	prog, err := cc.ReadMany(files, r)
	if err != nil {
		log.Fatal(err)
	}
	_ = prog
}