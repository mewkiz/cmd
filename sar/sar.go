package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
)

// flagInPlace specifies if a file should be edited in place.
var flagInPlace bool

func init() {
	flag.Usage = usage
	flag.BoolVar(&flagInPlace, "i", false, "Edit file in place.")
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: sar [OPTION]... SEARCH REPLACE [FILE]")
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "Flags:")
	flag.PrintDefaults()
}

func main() {
	flag.Parse()
	if flag.NArg() < 2 {
		flag.Usage()
		os.Exit(1)
	}

	search, err := strconv.Unquote(`"` + flag.Arg(0) + `"`)
	if err != nil {
		log.Fatalln(err)
	}
	replace, err := strconv.Unquote(`"` + flag.Arg(1) + `"`)
	if err != nil {
		log.Fatalln(err)
	}
	if flag.NArg() == 2 {
		// input from: stdin
		input, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalln(err)
		}
		output := sar(string(input), search, replace)
		fmt.Print(output)
	} else {
		// input from: FILE
		input, err := ioutil.ReadFile(flag.Arg(2))
		if err != nil {
			log.Fatalln(err)
		}
		output := sar(string(input), search, replace)
		if flagInPlace {
			fi, err := os.Stat(flag.Arg(2))
			if err != nil {
				log.Fatalln(err)
			}
			err = ioutil.WriteFile(flag.Arg(2), []byte(output), fi.Mode())
			if err != nil {
				log.Fatalln(err)
			}
		} else {
			fmt.Print(output)
		}
	}
}

// sar uses regexp to search and replace provided input.
func sar(input, search, replace string) (output string) {
	regSearch := regexp.MustCompile(search)
	output = regSearch.ReplaceAllString(input, replace)
	return output
}
