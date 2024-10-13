package utils

import (
	"strings"
	"unicode"
)

func tokenize(text string) []string {
	return strings.FieldsFunc(text, func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	})
}

func analyse(text string) []string {
	tokens := tokenize(text)
	tokens = lowercaseFilter(tokens)
	tokens = stopwordsFilter(tokens)
	tokens = stemmingFilter(tokens)
	return tokens
}
