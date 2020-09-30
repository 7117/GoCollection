package main

import "fmt"

func main() {
	var x [5]int

	x[0] = 1
	x[1] = 2

	fmt.Print(x)
	fmt.Printf("%v", x)

	f := [5]int{1, 2, 3}
	fmt.Printf("%v", f)

}
