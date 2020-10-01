package main

import "fmt"

func main() {
	var v1 interface{}

	v2 := make(map[string]int)
	v2["ss"] = 11;
	v1 = v2;
	switch v1.(type) {
	case int:
		fmt.Println("int");
	case float32:
		fmt.Println("float32");
	case float64:
		fmt.Println("float64");
	default:
		fmt.Println("other")
	}
}
