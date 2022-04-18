package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	// return map[string]int{"x": 1}
	wordCount := make(map[string]int)
	for _, i := range strings.Fields(s) {
		wordCount[i]++
	}
	return wordCount
}

func main() {
	wc.Test(WordCount)
}