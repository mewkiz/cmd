// The bindiff tool produces a patch based on the binary difference between two
// files, OLD and NEW. The patch can be applied to OLD using the binpatch tool
// to reproduce NEW.
//
// Usage:
//
//    bindiff [OPTION]... OLD NEW
//
// Flags:
//
//    -o string
//          patch output path
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/kr/binarydist"
)

func usage() {
	// Parse command line arguments.
	var (
		// output specifies the patch output path.
		output string
	)
	flag.StringVar(&output, "o", "", "patch output path")
	const use = `
Produce a patch based on the binary difference between two files.

Usage:

	bindiff [OPTION]... OLD NEW

Flags:
`
	fmt.Fprintln(os.Stderr, use[1:])
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() != 2 {
		flag.Usage()
		os.Exit(1)
	}
	var (
		oldPath   = flag.Arg(0)
		newPath   = flag.Arg(1)
		patchPath = oldPath + ".patch"
	)
	old, err := os.Open(oldPath)
	if err != nil {
		log.Fatalf("unable to open old binary %q; %v", oldPath, err)
	}
	defer old.Close()
	new, err := os.Open(newPath)
	if err != nil {
		log.Fatalf("unable to open new binary %q; %v", oldPath, err)
	}
	defer new.Close()
	patch, err := os.Create(patchPath)
	if err != nil {
		log.Fatalf("unable to create patch %q; %v", patchPath, err)
	}
	defer patch.Close()
	if err := binarydist.Diff(old, new, patch); err != nil {
		log.Fatalf("unable to create patch %q; %v", patchPath, err)
	}
}
