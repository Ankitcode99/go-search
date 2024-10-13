package utils

import (
	"encoding/xml"
	"os"
)

type document struct {
	Title string `xml:"title"`
	Url   string `xml:"url"`
	Text  string `xml:"abstract"`
	Id    int
}

func LoadDocuments(path string) ([]document, error) {

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	// gzipReader, err := gzip.NewReader(f)
	// if err != nil {
	// 	return nil, err
	// }
	// defer gzipReader.Close()

	decorder := xml.NewDecoder(f)
	dump := struct {
		Documents []document `xml:"doc"`
	}{}

	if err := decorder.Decode(&dump); err != nil {
		return nil, err
	}

	docs := dump.Documents
	for i := range docs {
		docs[i].Id = i
	}
	return docs, nil
}
