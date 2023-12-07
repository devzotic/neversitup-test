package main

import (
	"testing"
)

func TestFindOdd(t *testing.T) {

	cases := []struct {
		name  string
		in    []int
		out   int
		found bool
	}{
		{name: "found 7", in: []int{7}, out: 7, found: true},
		{name: "found 0", in: []int{0}, out: 0, found: true},
		{name: "found 2", in: []int{1, 1, 2}, out: 2, found: true},
		{name: "found 0", in: []int{0, 1, 0, 1, 0}, out: 0, found: true},
		{name: "found 4", in: []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}, out: 4, found: true},
		{name: "not found if input is empty", in: []int{}, out: 0, found: false},
		{name: "not found if all data for the input has an even number of times", in: []int{1, 1, 2, 3, 4, 2, 3, 4}, out: 0, found: false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {

			got, ok := findOdd(c.in)
			if got != c.out {
				t.Errorf("got %v want %v for output", got, c.out)
			}
			if ok != c.found {
				t.Errorf("got %v want %v for found state", c.in, c.found)
			}
		})
	}
}
