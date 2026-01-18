package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: " Hello  WORLD",
			expected: []string{"hello", "world"},
		},
	}


	for _, c := range cases {
		actual := cleanInput(c.input)

		// Check length of actual slice vs expected slice
		if len(actual) != len(c.expected) {
			t.Errorf("Returned wrong amount of strings")
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			// Check each word in slice
			if word != expectedWord {
				t.Errorf("Word %d does not match", i)
			}
		}
	}
}