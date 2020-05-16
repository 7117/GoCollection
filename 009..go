package main

import "fmt"

func swap(a int, b int) (int, int) {
	return b, a
}

func main() {
	a := 1
	b := 2

	a, b = swap(a, b)
	fmt.Printf("%v %v", a, b)
}
