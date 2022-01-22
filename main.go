package main

import (
	"fmt"
  "time"
  "flag"
	"os"
)

var (
  usemetadata bool
)

type MetaData  struct {
  Site string
  NumLinks int64
  images int64
  lastFetch time.Time
}

func init() {
  flag.BoolVar(&usemetadata, "metadata", false, "description")
}


func main() {
  os.Exit(run())
}

func run () int {
  flag.Parse()

  fmt.Println(flag.Args()) // 残りの引数
  fmt.Println(usemetadata)


  reqURLs := flag.Args()
  netURLs := make([]*NetURL, len(reqURLs))

  for i, reqURL := range reqURLs {
    netURL, err :=  NewNetURL(reqURL)
    if err != nil {
      fmt.Println("break")
      return 1
    }
    netURLs[i]= netURL
  }


  for _, netURL := range netURLs {
    fmt.Println("req")
    err := netURL.fetchAndCreateHTML()
    if (err != nil) {
      fmt.Println(err)
      return 1
    }
  }

  return 0
}


