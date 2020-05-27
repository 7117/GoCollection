package main

import "fmt"

func main() {
	var v1 interface{} = 2
	var v2 interface{} = "a"

	fmt.Print(v1)
	fmt.Print(v2)
}
