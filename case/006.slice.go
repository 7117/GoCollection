package main

import "fmt"

func main() {
	y := make([]int, 3, 5)
	fmt.Print(len(y))

	c := make([]int, 3, 5)
	c = append(c, 8, 8, 8, 10)

	fmt.Print(c)
	fmt.Print(cap(c))

	e := [5]int{1, 2, 3, 4, 5}
	fmt.Print(e)

}
