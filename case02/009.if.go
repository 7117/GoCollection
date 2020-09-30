package main

import "fmt"

func main() {
	var x int
	var sum int

	for x=1;x<=100;x++{
		sum += x
	}

	fmt.Println(sum)
}
