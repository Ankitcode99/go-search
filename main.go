package main

import (
	"flag"
	"log"
	"time"

	"github.com/Ankitcode99/full-text-search-engine/utils"
)

func main() {
	var dataDumpPath, searchQuery string

	flag.StringVar(&dataDumpPath, "path", "enwiki-latest-abstract1.xml", "Input file path")
	flag.StringVar(&searchQuery, "query", "Small wild cat", "Search query")
	flag.Parse()

	log.Println("Starting full text search!")
	start := time.Now()
	docs, err := utils.LoadDocuments(dataDumpPath)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Loaded %d documents in %v", len(docs), time.Since(start))

	start = time.Now()
	invertedIndex := make(utils.Index)

	invertedIndex.Add(docs)
	log.Printf("Indexed %d documents in %v", len(docs), time.Since(start))

	start = time.Now()
	matchedIds := invertedIndex.Search(searchQuery)
	log.Printf("Search found %d documents in %v", len(matchedIds), time.Since(start))

}
