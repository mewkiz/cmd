package main

import "flag"
import "fmt"
import "log"
import "net"
import "os"

func init() {
	flag.Usage = usage
	flag.Parse()
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: revdns IP...")
}

func main() {
	for _, rawIp := range flag.Args() {
		err := RevDNS(rawIp)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

// RevDNS performs a reverse DNS lookup on the provided IP-address.
func RevDNS(rawIp string) (err error) {
	names, err := net.LookupAddr(rawIp)
	if err != nil {
		return err
	}
	for _, name := range names {
		fmt.Println("name:", name)
	}
	return nil
}
