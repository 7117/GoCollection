package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Human struct {
	Person
	specially string
}

func main() {
	human := Human{Person{"zs", 11}, "qqq"}

	fmt.Println(human)
	fmt.Println(human.Person)
	fmt.Println(human.Person.name)
}
