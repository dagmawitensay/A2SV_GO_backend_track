package main

import "testing"


func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input string
		expected bool
	}{
		{
			input: "Hello World!",
			expected: false,
		},
		{
			input: "noon",
			expected: true,
		},
		{
			input: "madam",
			expected: true,
		},
		{
			input: "random string",
			expected: false,
		},
	}

	for _, test := range tests {
		result := IsPalindrome(test.input)
		if result != test.expected {
			t.Errorf("IsPalindrome(%q) = %t; expected %t", test.input, result, test.expected)
		}
	}
}