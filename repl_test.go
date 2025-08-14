package main

import (
	"testing"
)

func compareSlices(t *testing.T, actual, expected []string, input string) {
	if len(actual) != len(expected) {
		t.Errorf("Expected length %d, got %d for input '%s'", len(expected), len(actual), input)
		return
	}
	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("Expected '%s', got '%s' at index %d for input '%s'", expected[i], actual[i], i, input)
			return
		}
	}
}

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " It's Pikachu! ",
			expected: []string{"it's", "pikachu!"},
		},
		{
			input:    "  Bulbasaur, Charmander, Squirtle   ",
			expected: []string{"bulbasaur,", "charmander,", "squirtle"},
		},
		{
			input:    " Hello      Universe!",
			expected: []string{"hello", "universe!"},
		},
	}

	for _, c := range cases {
		actual := cleanInput((c.input))
		compareSlices(t, actual, c.expected, c.input)
	}
}
