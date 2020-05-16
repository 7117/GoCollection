package main

import "fmt"

func main() {
	type sum func(x, y int) int

	var f sum = func(x, y int) int {
		return x + y
	}

	fmt.Print(f(3, 4))
}
