package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManipulate(t *testing.T) {
	cases := []struct {
		name string
		in   string
		out  []string
	}{
		{"applied a", "a", []string{"a"}},
		{"applied ab", "ab", []string{"ab", "ba"}},
		{"applied abc", "abc", []string{"abc", "acb", "bac", "bca", "cab", "cba"}},
		{"applied aabb", "aabb", []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"}},
		{"applied aabb by ignoring the order of the elements in result", "aabb", []string{"baba", "abab", "abba", "aabb", "bbaa", "baab"}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var got []string
			Manipulate(c.in, 0, &got)

			assert.ElementsMatch(t, got, c.out)
		})
	}
}
