package utils

import (
	"strings"

	snowballeng "github.com/kljensen/snowball/english"
)

func stemmingFilter(tokens []string) []string {
	words := make([]string, len(tokens))
	for i, token := range tokens {
		words[i] = snowballeng.Stem(token, true)
	}
	return words
}

func stopwordsFilter(tokens []string) []string {
	var stopWords = map[string]struct{}{
		// add some common stopwords
		"a":       {},
		"the":     {},
		"of":      {},
		"and":     {},
		"or":      {},
		"to":      {},
		"with":    {},
		"at":      {},
		"on":      {},
		"in":      {},
		"for":     {},
		"from":    {},
		"by":      {},
		"i":       {},
		"that":    {},
		"as":      {},
		"is":      {},
		"this":    {},
		"it":      {},
		"he":      {},
		"she":     {},
		"him":     {},
		"we":      {},
		"you":     {},
		"me":      {},
		"himself": {},
	}

	words := make([]string, 0, len(tokens))

	for _, token := range tokens {
		if _, ok := stopWords[token]; !ok {
			words = append(words, token)
		}
	}
	return words
}

func lowercaseFilter(tokens []string) []string {
	words := make([]string, len(tokens))

	for i, token := range tokens {
		words[i] = strings.ToLower(token)
	}
	return words
}
