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
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		t.Logf("input: %q", c.input)
		t.Logf("expected: %v", c.expected)
		t.Logf("actual: %v", actual)

		// Check the length
		if len(actual) != len(c.expected) {
			t.Errorf("length mismatch for input %q: expected %d, got %d",
				c.input, len(c.expected), len(actual))
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			// Check each word
			if word != expectedWord {
				t.Errorf("for input %q at index %d: expected %q, got %q",
					c.input, i, expectedWord, word)
			}
		}
	}
}