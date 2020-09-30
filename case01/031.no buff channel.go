package main

import (
	"fmt"
	"time"
)

func test(ch chan int) {
	fmt.Print("test1")
	ch <- 1
	fmt.Print("test2")
}

func main() {
	ch := make(chan int, 0)
	go test(ch)
	time.Sleep(3 * time.Second)
	fmt.Print("main end")
	<-ch
	time.Sleep(3 * time.Second)

}
