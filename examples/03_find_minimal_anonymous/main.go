package main

import "fmt"

func findMinimal(check func(int) bool) int {
	i := 1
	for ; !check(i); i++ {
	}
	return i
}

func main() {
	fmt.Println("Минимальное число, больше 100:", findMinimal(func(x int) bool {
		return x > 100
	}))
	fmt.Println("Минимальное число, делимое на 11:", findMinimal(func(x int) bool {
		return x%11 == 0
	}))
	fmt.Println("Минимальное число, делимое на 11 и 13:", findMinimal(func(x int) bool {
		return x%11 == 0 && x%13 == 0
	}))
}
