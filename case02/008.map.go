package main

import "fmt"

func main() {
	var stu map[string]float32
	stu = make(map[string]float32)
	stu["zs"] = 11.11
	fmt.Printf("%v", stu)

	fmt.Println()

	tea := make(map[string]float32)
	tea["te"] = 22.22
	fmt.Printf("%v", tea)
}
