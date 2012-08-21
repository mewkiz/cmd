package main

import "flag"
import "fmt"
import "io/ioutil"
import "log"
import "os"
import "regexp"

// flagInPlace specifies if a file should be edited in place.
var flagInPlace bool

func init() {
	flag.Usage = usage
	flag.BoolVar(&flagInPlace, "i", false, "edit file in place.")
	flag.Parse()
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s SEARCH REPLACE [FILE]\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	if flag.NArg() < 2 {
		flag.Usage()
		return
	} else if flag.NArg() == 2 {
		// input from: stdin
		input, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalln(err)
		}
		output := sar(string(input), flag.Arg(0), flag.Arg(1))
		fmt.Print(output)
	} else {
		// input from: FILE
		input, err := ioutil.ReadFile(flag.Arg(2))
		if err != nil {
			log.Fatalln(err)
		}
		output := sar(string(input), flag.Arg(0), flag.Arg(1))
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
