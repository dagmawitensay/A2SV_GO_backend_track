package main

import (
	"testing"
)

func TestWordFreqeuncyCounter(t *testing.T) {
	tests := []struct{
		input string
		expected map[string]int
	}{
		{
			input: "Hello World!",
			expected: map[string]int{"Hello": 1, "World": 1},
		},
		{
			input:    "No punctuation here",
			expected: map[string]int{"No": 1, "punctuation": 1, "here": 1},
		},
		{
			input:    "Test input, with some punctuation.",
			expected: map[string]int{"Test": 1, "input": 1, "with": 1, "some": 1, "punctuation": 1},
		},
	}

	for _, test := range tests {
		result := WordFrquencyCounter(test.input)
		if len(result) != len(test.expected) {
			t.Errorf("WordFrequencyCounter(%q) = %v; expected %v", test.input, result, test.expected)
			continue
		}

		for word, count := range test.expected {
			if count != result[word] {
				t.Errorf("Count of Word in result = %d; expected %d", result[word], count)
			}
		}
	}
}