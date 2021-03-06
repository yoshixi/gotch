package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	useMetadata bool
)

func init() {
	flag.BoolVar(&useMetadata, "metadata", false, "description")
}

func main() {
	flag.Parse()
	reqURLs := flag.Args()

	os.Exit(Run(reqURLs, useMetadata))
}

func Run(reqURLs []string, readMetadata bool) int {
	netURLs := make([]*NetURL, len(reqURLs))

	for i, reqURL := range reqURLs {
		netURL, err := NewNetURL(reqURL)
		if err != nil {
			fmt.Println("break")
			return 1
		}
		netURLs[i] = netURL
	}

	for _, netURL := range netURLs {
		if readMetadata {
			m, _ := NewMetaData(netURL.url.Hostname())
			m.ReadAndPrint()
		} else {
			err := netURL.FetchAndCreateHTML()
			if err != nil {
				fmt.Println(err)
				return 1
			}
		}
	}

	return 0
}
