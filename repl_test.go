package main

import "testing"

func TestCleanInput(t *testing.T) {
	// Anonymous struct type, commonly used for table tests
	cases := []struct {
		input string
		expected []string
	}{
		{
		input: " hello world ",
		expected: []string{"hello", "world"},
		},
		{
		input: "my name  is sean!   ",
		expected: []string{"my", "name", "is", "sean!"},
		},
		{
		input: "ONE morE 	TEst",
		expected: []string{"one", "more", "test"},
		},
	}

	// Loop over each test case
	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check length, compare to actual
		if len(c.expected) != len(actual) {
			t.Errorf("Length of splice incorrect. Got: %v, Want: %v", len(c.expected), len(actual))
		}

		// Check each word in the slice, compare to actual
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Wrong word. Got: %q, Want:  %q ", word, expectedWord)
			}
		}
	}
}
