package main

import (
	"strings"
	"regexp"
)

func WordFrquencyCounter(s string) map[string]int {
	s = regexp.MustCompile(`[[:punct:]]`).ReplaceAllString(s, "")
	words := strings.Fields(s)
	counts := make(map[string]int)

	for i := 0; i < len(words); i++ {
		counts[words[i]] += 1
	}

	return counts
}