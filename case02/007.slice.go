package main

import "fmt"

func main() {

	var slice []int
	slice = append(slice, 1, 2, 4)
	fmt.Println(slice)

	a := make([]int, 3, 5);
	a[1] = 22;
	fmt.Println(a)

	a = append(a, 4, 5, 6)
	fmt.Println(a)
}
