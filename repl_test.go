package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  you shall not pass  ",
			expected: []string{"you", "shall", "not", "pass"},
		},
		{
			input:    "  did  we just become best friends  ",
			expected: []string{"did", "we", "just", "become", "best", "friends"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Expected length %d, got %d", len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Expected work %q, got %q", expectedWord, word)
			}
		}
	}
}
