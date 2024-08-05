package main

import (
	"strings"
	"regexp"
)

func IsPalindrome(str string) bool {
	str = strings.ToLower(str)
	str = regexp.MustCompile(`[[:punct:]]`).ReplaceAllString(str, "")
	str = strings.Join(strings.Fields(str), "")
	i, j := 0, len(str) - 1

	for i < len(str) && j > 0 {
		if str[i] != str[j]{
			return false
		}
		j -= 1
		i += 1
	}
	return true
}