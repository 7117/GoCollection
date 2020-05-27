package main

import "fmt"

func main() {

	defer fmt.Print("defer1")

	panic("panic ")

	defer fmt.Print("defer2")

}
