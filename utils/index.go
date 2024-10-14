package utils

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
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

		start := time.Now()

		go func(batch []document, batchMap Index, batchIndex int) {
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

			fmt.Printf("batch %d completed in %s\n", batchIndex, time.Since(start))
		}(batch, batchMaps[batchIndex], batchIndex)
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

var CONCURRENT_THREADS = 50 * 2
var idx int32 = 0

func doWork(name string, wg *sync.WaitGroup, docs []document, index *sync.Map) {
	start := time.Now()
	defer wg.Done()
	for {
		x := atomic.AddInt32(&idx, 1) - 1
		if int(x) >= len(docs) {
			break
		}
		doc := docs[int(x)]

		for _, token := range analyse(doc.Text) {
			index.LoadOrStore(token, &sync.Map{})
			tokenMap, _ := index.Load(token)
			tokenMap.(*sync.Map).Store(doc.Id, struct{}{})
		}
	}
	fmt.Printf("thread %s completed in %s\n", name, time.Since(start))
}

func (index *Index) AddConcurrent(docs []document) {
	var wg sync.WaitGroup
	concurrentIndex := &sync.Map{}

	for i := 0; i < CONCURRENT_THREADS; i++ {
		wg.Add(1)
		go doWork(strconv.Itoa(i), &wg, docs, concurrentIndex)
	}

	wg.Wait()

	// Convert sync.Map back to regular map
	concurrentIndex.Range(func(key, value interface{}) bool {
		token := key.(string)
		ids := []int{}
		value.(*sync.Map).Range(func(k, v interface{}) bool {
			ids = append(ids, k.(int))
			return true
		})
		(*index)[token] = ids
		return true
	})
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
