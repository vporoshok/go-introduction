package main

import (
	"testing"
)

func TestSort(t *testing.T) {
	if testing.Short() {
		t.Skip("Long test")
	}
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
			t.Parallel()
			InsertSort(c.input, BinSearch)
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
	if t.Failed() {
		t.Log("expected", expected, "actual", actual)
	}
}

func TestBinarySearch(t *testing.T) {
	cases := [...]struct {
		name string
		a    []int
		x    int
		i    int
	}{
		{
			"all equals",
			[]int{1, 1, 1, 1},
			1,
			4,
		},
		{
			"maximum",
			[]int{1, 2, 3, 4},
			5,
			4,
		},
		{
			"minimum",
			[]int{1, 2, 3, 4},
			0,
			0,
		},
		{
			"simple",
			[]int{1, 2, 4, 5},
			3,
			2,
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			i := BinSearch(c.a, c.x)
			if i != c.i {
				t.Errorf("expected %d instead of %d", c.i, i)
			}
		})
	}
}

func TestExpSearch(t *testing.T) {
	cases := [...]struct {
		name string
		a    []int
		x    int
		i    int
	}{
		{
			"all equals",
			[]int{1, 1, 1, 1},
			1,
			4,
		},
		{
			"maximum",
			[]int{1, 2, 3, 4},
			5,
			4,
		},
		{
			"minimum",
			[]int{1, 2, 3, 4},
			0,
			0,
		},
		{
			"simple",
			[]int{1, 2, 4, 5},
			3,
			2,
		},
		{
			"one element",
			[]int{2},
			1,
			0,
		},
		{
			"foo",
			[]int{2, 5, 6},
			3,
			1,
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			i := ExpSearch(c.a, c.x)
			if i != c.i {
				t.Errorf("expected %d instead of %d", c.i, i)
			}
		})
	}
}
