package utils

import (
	"sync"
)

type Index map[string][]int

func (index *Index) Add(docs []document) {
	batchSize := 20000
	var wg sync.WaitGroup

	// Create and initialize concurrent maps for each batch
	batchMaps := make([]Index, (len(docs)-1)/batchSize+1)
	for i := range batchMaps {
		batchMaps[i] = make(Index)
	}

	// Process documents in batches
	for i := 0; i < len(docs); i += batchSize {
		end := i + batchSize
		if end > len(docs) {
			end = len(docs)
		}

		batch := docs[i:end]
		batchIndex := i / batchSize
		wg.Add(1)

		go func(batch []document, batchMap Index) {
			defer wg.Done()

			// Process each document in the batch
			for _, doc := range batch {
				for _, token := range analyse(doc.Text) {
					ids := batchMap[token]
					if ids == nil || ids[len(ids)-1] != doc.Id {
						batchMap[token] = append(ids, doc.Id)
					}
				}
			}
		}(batch, batchMaps[batchIndex])
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Merge batch maps into the main index
	for _, batchMap := range batchMaps {
		for token, ids := range batchMap {
			(*index)[token] = append((*index)[token], ids...)
		}
	}
}

func (index *Index) Search(query string) []int {
	tokens := analyse(query)
	var result []int
	for _, token := range tokens {
		if ids, ok := (*index)[token]; ok {
			if result == nil {
				result = ids
			} else {
				result = intersection(result, ids)
			}
		} else {
			return nil
		}
	}
	return result
}

func intersection(a, b []int) []int {
	var result []int

	mapA := make(map[int]bool)

	for _, v := range a {
		mapA[v] = true
	}
	for _, v := range b {
		if _, ok := mapA[v]; ok {
			result = append(result, v)
		}
	}
	return result
}
