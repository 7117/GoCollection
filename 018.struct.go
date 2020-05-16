package main

import "fmt"

type Person struct{
	name string
	age int
}

type Student struct{
	Person
	num string
}

func main(){
	person := Person{"张三",12};
	fmt.Print(person);

	student:=Student{Person{"账务",11},"student"};
	fmt.Print(student);
	fmt.Print(student.num);
	fmt.Print(student.Person.age);
}
