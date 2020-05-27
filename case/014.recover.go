package main

import "fmt"

func main() {
	defer func() {
		fmt.Print("1")
		recover()
		fmt.Print("2")
	}()
	panic("panic")

	fmt.Print("3")
}
