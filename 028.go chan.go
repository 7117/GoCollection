package main

import "fmt"

func Read(ch chan int) {
	value := <-ch
	fmt.Println("value")
}

func main() {
	var ch chan int

	ch = make(chan int)
}
