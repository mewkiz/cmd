package main

import "flag"
import "fmt"
import "io/ioutil"
import "log"
import "os"
import "regexp"

func init() {
   flag.Parse()
   flag.Usage = usage
}

func usage() {
   fmt.Fprintf(os.Stderr, "usage: %s SEARCH REPLACE [FILE]\n", os.Args[0])
   flag.PrintDefaults()
}

func main() {
   var input string
   switch flag.NArg() {
   case 2:
      buf, err := ioutil.ReadAll(os.Stdin)
      if err != nil {
         log.Fatalln(err)
      }
      input = string(buf)
   case 3:
      buf, err := ioutil.ReadFile(flag.Arg(2))
      if err != nil {
         log.Fatalln(err)
      }
      input = string(buf)
   default:
      flag.Usage()
      return
   }
   fmt.Print(sar(input, flag.Arg(0), flag.Arg(1)))
}

func sar(input, search, replace string) (output string) {
   regSearch := regexp.MustCompile(search)
   output = regSearch.ReplaceAllString(input, replace)
   return output
}
