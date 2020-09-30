package main

import "fmt"

type Animal interface {
	Fly()
	Run()
}

// 此处不需要写  interface
// 因为bird实现了Animal的所有方法
// 所以就默认与animal进行了连接了！
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

	bird.Fly()
	bird.Run()

}
