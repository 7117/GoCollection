package main

import "fmt"

func main() {
	var v1 interface{}

	v1 = 3.33;

	if v, ok := v1.(float32); ok {
		fmt.Println(v,ok)
	}else{
		fmt.Println(v,ok)
	}
}
