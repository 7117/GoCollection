package main

import "fmt"

func add(a *int) *int {
	*a = *a + 1
	return a
}

func main() {
	a := 1
	// 传递的指针
	add(&a)
	fmt.Printf("%v", a)
}
