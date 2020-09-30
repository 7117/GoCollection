package main

import "fmt"

func main() {
	x := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(x)

	y := x[1:3]
	fmt.Println(y)

	z := y[1]
	fmt.Println(z)
}
