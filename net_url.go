package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type NetURL struct {
	url *url.URL
}

func NewNetURL(reqURL string) (*NetURL, error) {
	u, err := url.ParseRequestURI(reqURL)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &NetURL{url: u}, nil
}

func (p *NetURL) fetchHtml() (string, error) {
	reqURL := p.url.String()
	res, err := http.Get(reqURL)

	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	buf := bytes.NewBuffer(body)
	html := buf.String()

	if res.StatusCode > 299 {
		return html, fmt.Errorf("StatusCode %s requestURL: %s", res.Status, reqURL)
	}

	return html, nil
}

func (p *NetURL) fetchAndCreateHTML() error {
	htmlString, err := p.fetchHtml()
	if err != nil {
		return err
	}

	fmt.Println("create  file start")
	err = writeFile(htmlString, fmt.Sprintf("%s.html", p.url.Hostname()))

	if err != nil {
		return err
	}

	fmt.Println("create  metadata start")
	m, _ := NewMetaData(p.url.Hostname())
	m.SetMetaData(htmlString)
	m.Store()
	fmt.Printf("%v", m)

	return nil
}

func writeFile(htmlString string, fileName string) error {
	fp, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer fp.Close()
	fp.WriteString(htmlString)

	return nil
}
