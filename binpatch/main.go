// The binpatch tool applies the patch to the binary file OLD to recreate the
// binary file NEW. The patch was produced by the bindiff tool based on the
// binary difference between OLD and NEW.
//
// Usage:
//
//    binpatch [OPTION]... OLD [PATCH]
//
// Flags:
//
//    -o string
//          output path of the binary file NEW
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

// dbg represents a logger with the "binpatch:" prefix, which logs debug
// messages to standard error.
var dbg = log.New(os.Stderr, term.MagentaBold("binpatch:")+" ", 0)

func usage() {
	const use = `
Apply the patch to the binary file OLD to recreate the binary file NEW.

Usage:

	binpatch [OPTION]... OLD [PATCH]

Flags:
`
	fmt.Fprintln(os.Stderr, use[1:])
	flag.PrintDefaults()
}

func main() {
	// Parse command line arguments.
	var (
		// output specifies the output path of the binary file NEW.
		output string
	)
	flag.StringVar(&output, "o", "", "output path of the binary file NEW")
	flag.Usage = usage
	flag.Parse()
	var (
		patchPath string
		patch     io.Reader
	)
	switch flag.NArg() {
	case 1:
		// Read patch from standard input.
		patchPath = "<stdin>"
		patch = os.Stdin
	case 2:
		patchPath = flag.Arg(1)
		f, err := os.Open(patchPath)
		if err != nil {
			log.Fatalf("unable to open patch %q; %+v", patchPath, errors.WithStack(err))
		}
		defer f.Close()
		patch = f
	default:
		flag.Usage()
		os.Exit(1)
	}
	oldPath := flag.Arg(0)
	newPath := oldPath + ".new"
	if len(output) > 0 {
		newPath = output
	}
	old, err := os.Open(oldPath)
	if err != nil {
		log.Fatalf("unable to open old binary %q; %+v", oldPath, errors.WithStack(err))
	}
	defer old.Close()
	new, err := os.Create(newPath)
	if err != nil {
		log.Fatalf("unable to create new binary %q; %+v", oldPath, errors.WithStack(err))
	}
	defer new.Close()
	dbg.Println("recreating binary file NEW:", newPath)
	if err := binarydist.Patch(old, new, patch); err != nil {
		log.Fatalf("unable to apply patch %q; %+v", patchPath, errors.WithStack(err))
	}
}
