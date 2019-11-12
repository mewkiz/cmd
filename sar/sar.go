package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"

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
		// inPlace specifies whether to edit the file in place.
		inPlace bool
	)
	flag.BoolVar(&inPlace, "i", false, "Edit file in place.")
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() < 2 {
		flag.Usage()
		os.Exit(1)
	}
	search := flag.Arg(0)
	replace := flag.Arg(1)

	// Perform regexp search and replace.
	if flag.NArg() == 2 {
		// input from: stdin
		input, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalf("%+v", errors.WithStack(err))
		}
		output := sar(string(input), search, replace)
		fmt.Print(output)
	} else {
		// input from: FILE
		input, err := ioutil.ReadFile(flag.Arg(2))
		if err != nil {
			log.Fatalf("%+v", errors.WithStack(err))
		}
		output := sar(string(input), search, replace)
		if inPlace {
			fi, err := os.Stat(flag.Arg(2))
			if err != nil {
				log.Fatalf("%+v", errors.WithStack(err))
			}
			err = ioutil.WriteFile(flag.Arg(2), []byte(output), fi.Mode())
			if err != nil {
				log.Fatalf("%+v", errors.WithStack(err))
			}
		} else {
			fmt.Print(output)
		}
	}
}

// sar returns the result of a regexp search and replace on the given input.
func sar(input, search, replace string) string {
	regSearch := regexp.MustCompile(search)
	output := regSearch.ReplaceAllString(input, replace)
	return output
}
