package main

import "fmt"

func main() {

	f := func(x int, y int) int {
		return x + y
	}

	fmt.Print(f(2, 43))

}
