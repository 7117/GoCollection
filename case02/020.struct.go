package main

import "fmt"

type Person struct {
	name string
	age int
}

func main() {
	person := Person{"zs",20}

	fmt.Println(person)
}
