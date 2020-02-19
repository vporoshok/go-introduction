package main

import "fmt"

func ExpSearch(a []int, x int) int {
	for k, n := 0, len(a); ; n = n - k/2 {
		for k = 1; k <= n && a[n-k] > x; k *= 2 {
		}
		if k == 1 {
			return n
		}
	}
}

func BinSearch(a []int, x int) int {
	l, r := 0, len(a)
	for l < r {
		if a[(r+l)/2] > x {
			r = (r + l) / 2
		} else {
			l = (r+l)/2 + 1
		}
	}
	return r
}

func InsertSort(a []int, search func([]int, int) int) {
	for i := 1; i < len(a); i++ {
		j := search(a[:i], a[i])
		buf := a[i]
		copy(a[j+1:i+1], a[j:i])
		a[j] = buf
	}
}

func main() {
	a := []int{12, 8, 22, 11, 1, 3}
	InsertSort(a, BinSearch)
	fmt.Println(a) // [1 3 8 11 12 22]
}
