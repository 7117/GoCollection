package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	person := Person{"张三", 12}

	fmt.Print(person)
}
