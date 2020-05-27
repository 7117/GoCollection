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
	ch := make(chan int, 1)
	go test(ch)
	time.Sleep(1 * time.Second)
	fmt.Print("mainend")
	<-ch
	time.Sleep(1 * time.Second)

}
