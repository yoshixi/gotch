package main

import (
	"os"
	"testing"
	"time"
)

func TestNewMetaData(t *testing.T) {
	site := "string"

	subject, err := NewMetaData(site)

	if err != nil {
		t.Errorf("TestNewMetaData: error %v", err)
	}

	if subject.Images != 0 {
		t.Errorf("TestNewMetaData: not expected Images")
	}

	if subject.NumLinks != 0 {
		t.Errorf("TestNewMetaData: not expected NumLinks")
	}
}

func TestSetMetaData(t *testing.T) {
	testhtml, _ := os.ReadFile("testdata/test.html")
	htmlString := string(testhtml)
	testhtml2, _ := os.ReadFile("testdata/test2.html")
	htmlString2 := string(testhtml2)

	tests := []struct {
		inputHtmlString string
		outMetadata     MetaData
	}{
		{
			inputHtmlString: htmlString,
			outMetadata:     MetaData{Site: "test", NumLinks: 1, Images: 1, LastFetch: time.Now()},
		},
		{
			inputHtmlString: htmlString2,
			outMetadata:     MetaData{Site: "test2", NumLinks: 2, Images: 2, LastFetch: time.Now()},
		},
	}

	for _, tt := range tests {
		metadata, _ := NewMetaData(tt.outMetadata.Site)
		metadata.SetMetaData(tt.inputHtmlString)

		if metadata.Images != tt.outMetadata.Images {
			t.Errorf("expected images to set %d, but setted %d", metadata.Images, tt.outMetadata.Images)
		}

		if metadata.NumLinks != tt.outMetadata.NumLinks {
			t.Errorf("expected images to set %d, but setted %d", metadata.NumLinks, tt.outMetadata.NumLinks)
		}
	}
}
