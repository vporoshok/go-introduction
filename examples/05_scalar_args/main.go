package main

import "fmt"

func increment(x int) {
	x++
}

func main() {
	n := 12
	increment(n)
	fmt.Println(n) // 12
}
