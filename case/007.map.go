package main

import "fmt"

func main() {
	var aa map[string]int
	aa = make(map[string]int)

	aa["a"] = 10
	aa["b"] = 20

	fmt.Printf("%v", aa)

	bb := make(map[string]int)
	bb["f"] = 10
	bb["w"] = 20
	fmt.Printf("%v", bb)

}
