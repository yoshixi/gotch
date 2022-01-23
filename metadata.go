package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/html"
	"net/url"
	"os"
	"strings"
	"time"
)

const MetadataDir = ".metadata"
const MetdataOutputTemplate = `site: %s
num_links: %d
images: %d
last_fetch: %s
`

type MetaData struct {
	Site      string    `json:"site"`
	NumLinks  int       `json:"num_links"`
	Images    int       `json:"images"`
	LastFetch time.Time `json:"last_fetch"`
}

func NewMetaData(site string) (*MetaData, error) {
	return &MetaData{Site: site, NumLinks: 0, Images: 0, LastFetch: time.Now()}, nil
}

func (m *MetaData) SetMetaData(htmlString string) {
	m.setMetadataFromHTMLString(htmlString)
}

func (m *MetaData) setMetadataFromHTMLString(htmlString string) {
	r := strings.NewReader(htmlString)
	tokenizer := html.NewTokenizer(r)
	end := false

	for {
		tt := tokenizer.Next()
		switch {
		case tt == html.StartTagToken || tt == html.SelfClosingTagToken:
			token := tokenizer.Token()
			switch token.Data {
			case "a":

				for _, attr := range token.Attr {

					if attr.Key == "href" {
						link := attr.Val

						_, parseLinkErr := url.Parse(link)
						if parseLinkErr == nil {
							m.NumLinks++
						}

						if parseLinkErr != nil {
							fmt.Println("Can't parse: " + token.Data)
						}
					}
				}
				break
			case "img":
				m.Images++
			}
		case tt == html.ErrorToken:
			end = true
			break
		}
		if end {
			break
		}
	}
}

func (m *MetaData) Store() error {
	data, _ := json.Marshal(m)
	os.Mkdir(MetadataDir, 0777)
	return writeFile(string(data), m.getFilePath())
}

func (m *MetaData) ReadAndPrint() {
	jsonFromFile, err := os.ReadFile(m.getFilePath())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	json.Unmarshal(jsonFromFile, m)

	fmt.Printf(MetdataOutputTemplate, m.Site, m.NumLinks, m.Images, m.LastFetch.UTC().Format(time.RFC1123))
}

func (m *MetaData) getFilePath() string {
	return fmt.Sprintf("%s/%s.json", MetadataDir, m.Site)
}
