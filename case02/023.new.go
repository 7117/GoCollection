package main

import "fmt"

type Personone struct {
	name string
	age  int
}

type Stuone struct {
	Personone
	specially string
}

func (person Personone) getNameadnAge() (string, int) {
	return person.name, person.age
}

func main() {
	studentone := new(Stuone)
	studentone.name = "zs";
	studentone.age = 12;
	name, age := studentone.getNameadnAge()
	fmt.Println(name, age)
}
