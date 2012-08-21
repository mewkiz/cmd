package main

import "bufio"
import "flag"
import "fmt"
import "log"
import "os"
import "sort"

import "github.com/mewkiz/pkg/bufioutil"

func init() {
	flag.Usage = usage
	flag.Parse()
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [FILE]...\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	err := rawSort(flag.Args())
	if err != nil {
		log.Fatalln(err)
	}
}

// rawSort writes the sorted concatenation of all provided files to standard
// output. If no file is provided it will read from standard input.
func rawSort(filePaths []string) (err error) {
	lines := make([]string, 0)
	for _, filePath := range filePaths {
		// read the lines from all
		fr, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer fr.Close()
		br := bufio.NewReader(fr)
		l, err := bufioutil.ReadLines(br)
		if err != nil {
			return err
		}
		lines = append(lines, l...)
	}
	if len(filePaths) == 0 {
		br := bufio.NewReader(os.Stdin)
		l, err := bufioutil.ReadLines(br)
		if err != nil {
			return err
		}
		lines = append(lines, l...)
	}
	sort.Strings(lines)
	for _, line := range lines {
		fmt.Println(line)
	}
	return nil
}
