package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	usemetadata bool
)

func init() {
	flag.BoolVar(&usemetadata, "metadata", false, "description")
}

func main() {
	os.Exit(run())
}

func run() int {
	flag.Parse()

	fmt.Println(flag.Args()) // 残りの引数
	fmt.Println(usemetadata)

	reqURLs := flag.Args()
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
		if usemetadata {
			m, _ := NewMetaData(netURL.url.Hostname())
			m.Read()
		} else {
			err := netURL.fetchAndCreateHTML()
			if err != nil {
				fmt.Println(err)
				return 1
			}
		}
	}

	return 0
}
