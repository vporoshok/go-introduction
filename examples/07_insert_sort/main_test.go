package main

import (
	"testing"
)

func TestSort(t *testing.T) {
	cases := [...]struct {
		name   string
		input  []int
		result []int
	}{
		{
			"already sorted",
			[]int{1, 2, 3, 4, 5, 6},
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			"inversed",
			[]int{6, 5, 4, 3, 2, 1},
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			"duplicates",
			[]int{6, 5, 2, 3, 2, 1},
			[]int{1, 2, 2, 3, 5, 6},
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			InsertSort(c.input)
			CompareIntSlices(t, c.result, c.input)
		})
	}
}

func CompareIntSlices(t *testing.T, expected, actual []int) {
	t.Helper()
	if len(expected) != len(actual) {
		t.Errorf("Invalid slice length. Expected %d instead of %d",
			len(expected), len(actual))
		return
	}
	for i := range expected {
		if expected[i] != actual[i] {
			t.Errorf("Invalid element at position %d. Expected %d instead of %d",
				i, expected[i], actual[i])
		}
	}
}
