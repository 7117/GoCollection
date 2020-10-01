package main

import "fmt"

func main() {
	var v1 interface{}

	v1 = "张三";
	v1 = "111";
	v2 := make(map[int]string)
	v2[5] = "rf"
	v1 =v2

	fmt.Println(v1);
	fmt.Println(v2)
}
