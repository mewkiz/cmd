// The bindiff tool produces a patch based on the binary difference between the
// two files OLD and NEW. The patch can be applied to OLD using the binpatch
// tool to recreate NEW.
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
	"io"
	"log"
	"os"

	"github.com/kr/binarydist"
	"github.com/mewkiz/pkg/term"
	"github.com/pkg/errors"
)

// dbg represents a logger with the "bindiff:" prefix, which logs debug messages
// to standard error.
var dbg = log.New(os.Stderr, term.MagentaBold("bindiff:")+" ", 0)

func usage() {
	const use = `
Produce a patch based on the binary difference between the two files OLD and NEW.

Usage:

	bindiff [OPTION]... OLD NEW

Flags:
`
	fmt.Fprintln(os.Stderr, use[1:])
	flag.PrintDefaults()
}

func main() {
	// Parse command line arguments.
	var (
		// output specifies the patch output path.
		output string
	)
	flag.StringVar(&output, "o", "", "patch output path")
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
	// Handle -o flag.
	if len(output) > 0 {
		patchPath = output
	}
	old, err := os.Open(oldPath)
	if err != nil {
		log.Fatalf("unable to open old binary %q; %+v", oldPath, errors.WithStack(err))
	}
	defer old.Close()
	new, err := os.Open(newPath)
	if err != nil {
		log.Fatalf("unable to open new binary %q; %+v", oldPath, errors.WithStack(err))
	}
	defer new.Close()
	patch, err := os.Create(patchPath)
	if err != nil {
		log.Fatalf("unable to create patch %q; %+v", patchPath, errors.WithStack(err))
	}
	defer patch.Close()
	dbg.Println("creating patch:", patchPath)
	if err := createDiff(old, new, patch); err != nil {
		log.Fatalf("unable to create patch %q; %+v", patchPath, errors.WithStack(err))
	}
}

// createDiff produces a patch based on the binary difference between OLD and
// NEW.
func createDiff(old, new io.Reader, patch io.Writer) error {
	if err := binarydist.Diff(old, new, patch); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
