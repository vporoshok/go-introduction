package main

import "fmt"

func main() {
	text := `Welcome to a tour of the Go programming language. The tour is
divided into a list of modules that you can access by clicking on A Tour of Go
on the top left of the page. You can also view the table of contents at any
time by clicking on the menu on the top right of the page. Throughout the tour
you will find a series of slides and exercises for you to complete.`
	counts := map[rune]int{}
	for _, c := range text {
		switch c {
		case ' ', '.':
			// не будем учитывать эти символы в статистике
		default:
			counts[c] += 1
		}
	}
	var (
		res rune
		max int
	)
	for c, count := range counts {
		if count > max {
			res = c
			max = count
		}
	}
	fmt.Println(string(res)) // преобразуем руну в строку, иначе будет выведен код
}
