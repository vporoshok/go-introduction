package main

import "fmt"

func InsertSort(a []int) {
	for i := 1; i < len(a); i++ {
		j := i
		for ; j > 0 && a[i] < a[j-1]; j-- {
		}
		for ; i > j; i-- {
			a[i], a[i-1] = a[i-1], a[i]
		}
	}
}

func main() {
	a := []int{12, 8, 22, 11, 1, 3}
	InsertSort(a)
	fmt.Println(a) // [1 3 8 11 12 22]
}
