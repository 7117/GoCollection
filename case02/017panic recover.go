package main

import "fmt"

func main() {
	defer fmt.Println("after")

	panic("i am wrong")



}
