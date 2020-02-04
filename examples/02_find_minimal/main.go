package main

import "fmt"

func findMinimal(check func(int) bool) int {
	i := 1
	for ; !check(i); i++ {
	}
	return i
}

func greatThan100(x int) bool {
	return x > 100
}

func dividedBy11(x int) bool {
	return x%11 == 0
}

func dividedBy11And13(x int) bool {
	return x%11 == 0 && x%13 == 0
}

func main() {
	fmt.Println("Минимальное число, больше 100:", findMinimal(greatThan100))
	fmt.Println("Минимальное число, делимое на 11:", findMinimal(dividedBy11))
	fmt.Println("Минимальное число, делимое на 11 и 13:", findMinimal(dividedBy11And13))
}
