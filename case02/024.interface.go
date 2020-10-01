package main

import "fmt"

type Animal interface {
	Fly() bool
	Run() bool
}

type Bird struct {
}

func (bird Bird) Fly() bool {
	fmt.Println("bird is fly")
	return true;
}

func (Bird Bird) Run()bool{
	fmt.Println("bird is run")
	return true;
}

func main() {
	bird := new(Bird)

	bird.Fly();
	bird.Run();
}
