package main

import "fmt"

type Animal1 interface {
	Fly()
	Run()
}

type Animal2 interface {
	Fly()
}

type Bird struct {
}

func (bird Bird) Fly() {
	fmt.Print("bird is flying")
}

func (bird Bird) Run() {
	fmt.Print("bird is runing")
}

func main() {

	bird := new(Bird)

	var Animal1 Animal1
	var Animal2 Animal2

	// Animal1 = bird;
	// Animal1.Fly();
	// Animal2 = bird;
	// Animal2.Fly();

	// 报错
	// Animal2 = bird;
	// Animal1 = Animal2;
	// Animal1.Fly();

	// 正常
	Animal1 = bird
	Animal2 = Animal1
	Animal2.Fly()
}
