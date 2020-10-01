package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go test(ch)

	fmt.Println("main")
	time.Sleep(2*time.Second)
	<-ch
	time.Sleep(1*time.Second)
}

func test(ch chan int){
	ch <- 1
	fmt.Println("come")
}

