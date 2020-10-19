package main

import "fmt"

type Animal interface {
	Fly() bool
	Run() bool
}

type Animale2 interface {
	Fly() bool
}
type Bird struct {
}

func (bird Bird) Fly() bool {
	fmt.Println("bird is fly")
	return true;
}

func (Bird Bird) Run() bool {
	fmt.Println("bird is run")
	return true;
}

func main() {
	//bird := new(Bird)
	//
	//bird.Fly();
	//bird.Run();

	var animal Animal
	var animal2 Animale2
	bird := new(Bird)

	animal = bird
	animal2 = bird

	animal.Fly();
	animal.Run();
	animal2.Fly();
}
