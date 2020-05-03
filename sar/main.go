/*
The sar tool performs regexp search and replace on the input.

Installation:

	go get -u github.com/mewkiz/cmd/sar

Usage:

	sar [OPTION]... SEARCH REPLACE [FILE]

Flags:

	-fixed-search
	      interpret SEARCH as a fixed string, not a regular expression
	-i    edit file in place
	-unescape-replace
	      unescape REPLACE string
*/
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func usage() {
	const use = `
Usage: sar [OPTION]... SEARCH REPLACE [FILE]

Flags:`
	fmt.Fprintln(os.Stderr, use[1:])
	flag.PrintDefaults()
}

func main() {
	// Parse command line arguments.
	var (
		// Interpret SEARCH as a fixed string, not a regular expression.
		fixed bool
		// Edit the file in place.
		inPlace bool
		// Unescape REPLACE string.
		unescape bool
	)
	flag.BoolVar(&fixed, "fixed-search", false, "interpret SEARCH as a fixed string, not a regular expression")
	flag.BoolVar(&inPlace, "i", false, "edit file in place")
	flag.BoolVar(&unescape, "unescape-replace", false, "unescape REPLACE string")
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() < 2 {
		flag.Usage()
		os.Exit(1)
	}
	search := flag.Arg(0)
	replace := flag.Arg(1)
	if unescape {
		quoted := `"` + replace + `"`
		s, err := strconv.Unquote(quoted)
		if err != nil {
			log.Fatalf("%+v", errors.WithStack(err))
		}
		replace = s
	}

	// Perform regexp search and replace.
	switch flag.NArg() {
	case 2:
		// input from: stdin
		input, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalf("%+v", errors.WithStack(err))
		}
		output := sar(string(input), search, replace, fixed)
		fmt.Print(output)
	default:
		// input from: FILE
		path := flag.Arg(2)
		input, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatalf("%+v", errors.WithStack(err))
		}
		output := sar(string(input), search, replace, fixed)
		if inPlace {
			fi, err := os.Stat(path)
			if err != nil {
				log.Fatalf("%+v", errors.WithStack(err))
			}
			err = ioutil.WriteFile(path, []byte(output), fi.Mode())
			if err != nil {
				log.Fatalf("%+v", errors.WithStack(err))
			}
		} else {
			fmt.Print(output)
		}
	}
}

// sar returns the result of a regexp search and replace on the given input.
func sar(input, search, replace string, fixed bool) string {
	if fixed {
		return strings.ReplaceAll(input, search, replace)
	}
	regSearch := regexp.MustCompile(search)
	output := regSearch.ReplaceAllString(input, replace)
	return output
}
