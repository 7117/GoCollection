package main

import "fmt"

func addNum(a *int) (int) {
	*a = *a + 1;
	return *a;
}

func main() {
	a := 1;
	c := addNum(&a);
	fmt.Println(c)
}
