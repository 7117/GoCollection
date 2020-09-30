package main

import "fmt"

func main() {

	c := addNum(1, 4);
	fmt.Println(c)
}

func addNum(a, b int) (int) {
	return a + b;
}
