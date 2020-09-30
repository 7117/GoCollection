package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func(ch chan int) {
		ch <- 1
	}(ch)

	time.Sleep(1 * time.Second)

	select {
	case <-ch:
		fmt.Print("one")
	default:
		fmt.Print("default")
	}

}
