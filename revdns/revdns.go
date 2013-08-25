package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

func init() {
	flag.Usage = usage
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: revdns IP...")
	fmt.Fprintln(os.Stderr, "Perform reverse DNS lookups on provided IP-addresses.")
}

func main() {
	flag.Parse()
	for _, rawIP := range flag.Args() {
		err := RevDNS(rawIP)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

// RevDNS performs a reverse DNS lookup on the provided IP-address.
func RevDNS(rawIP string) (err error) {
	names, err := net.LookupAddr(rawIP)
	if err != nil {
		return err
	}
	for _, name := range names {
		fmt.Printf("%s = %s\n", rawIP, name)
	}
	return nil
}
