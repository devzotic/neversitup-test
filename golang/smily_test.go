package main

import "testing"

func TestCountSmileys(t *testing.T) {

	cases := []struct {
		name     string
		input    []string
		expected int
	}{
		{"found 2", []string{":)", ";(", ";}", ":-D"}, 2},
		{"found 3", []string{";D", ":-(", ":-)", ";~)"}, 3},
		{"found 1", []string{";]", ":[", ";*", ":$", ";-D"}, 1},
		{"not found if all faces is invalid", []string{";}", ":-(", ";]", ":[", ";*", ":$"}, 0},
		{"not found if faces is empty", []string{}, 0},
		{"not found if faces is nil", nil, 0},
	}

	for _, c := range cases {

		t.Run(c.name, func(t *testing.T) {
			got := countSmileys(c.input)
			if got != c.expected {
				t.Errorf("expected %v but got %v", c.expected, got)
			}
		})

	}
}
