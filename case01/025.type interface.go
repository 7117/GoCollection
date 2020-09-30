package main

import "fmt"

func main() {
	var v1 interface{} = 11.2

	v1, ok := v1.(float64)

	// 11.2 true
	fmt.Print(v1, ok)
}
