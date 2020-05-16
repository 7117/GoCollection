package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Student struct {
	Person
	speciality string
}

func (person Person) getNameAndAge() (string, int) {
	return person.name, person.age
}

func (student Student) getSpeciality() string {
	return student.speciality
}

func main() {
	student := new(Student)

	student.name = "张三"
	student.age = 20
	student.speciality = "生物"

	name, age := student.getNameAndAge()
	fmt.Print(name, age)

	speciality := student.getSpeciality()
	fmt.Print(speciality)

}
