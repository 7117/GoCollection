package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int,1)

	go test(ch)

	time.Sleep(2*time.Second)
	fmt.Println("main")
}

func test(ch chan int){
	ch <- 1
	fmt.Println("come")
}

