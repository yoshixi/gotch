package main

import (
	"bytes"
	"fmt"
	"github.com/jarcoal/httpmock"
	"io"
	"net/url"
	"os"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	siteURLString := "https://test111111.gotch"
	siteURI, _ := url.ParseRequestURI(siteURLString)

	tests := []struct {
		name          string
		useMetadata   bool
		inputSiteURLs []string
		outMetadata   MetaData
		outLog        string
	}{
		{
			name:          "default",
			useMetadata:   false,
			inputSiteURLs: []string{siteURLString},
			outMetadata:   MetaData{Site: siteURI.Hostname(), NumLinks: 1, Images: 1, LastFetch: time.Now()},
		},
		{
			name:          "metadata flag",
			useMetadata:   true,
			inputSiteURLs: []string{siteURLString},
			outMetadata:   MetaData{Site: siteURI.Hostname(), NumLinks: 1, Images: 1, LastFetch: time.Now()},
		},
	}

	for _, tt := range tests {

		// mock
		for _, u := range tt.inputSiteURLs {
			testhtml, _ := os.ReadFile("testdata/test.html")
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			httpmock.RegisterResponder("GET", u,
				httpmock.NewBytesResponder(200, testhtml),
			)
		}

		if tt.useMetadata {
			t.Run(tt.name, func(t *testing.T) {
				Run(tt.inputSiteURLs, false)
				out := captureStdout(func() { Run(tt.inputSiteURLs, true) })

				expected := fmt.Sprintf(MetdataOutputTemplate, tt.outMetadata.Site, tt.outMetadata.NumLinks, tt.outMetadata.Images, tt.outMetadata.LastFetch.UTC().Format(time.RFC1123))
				if out != expected {
					t.Errorf("get %s, expected %s", out, expected)
				}
			})
		}

		if tt.name == "default" {
			t.Run(tt.name, func(t *testing.T) {
				Run(tt.inputSiteURLs, false)
				for _, siteURL := range tt.inputSiteURLs {
					parsedURL, _ := url.ParseRequestURI(siteURL)

					// check creating file
					if !fileExists(fmt.Sprintf("%s.html", parsedURL.Hostname())) {
						t.Errorf("html file not created")
					}

					// check createing metadata
					if !fileExists(fmt.Sprintf(".metadata/%s.json", parsedURL.Hostname())) {
						t.Errorf("metadata file not created")
					}
				}
			})
		}
	}
}

func tearDownTestRun(Hostname string) {
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func captureStdout(f func()) string {
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	stdout := os.Stdout
	os.Stdout = w

	f()

	os.Stdout = stdout
	w.Close()

	var buf bytes.Buffer
	io.Copy(&buf, r)

	return buf.String()
}
