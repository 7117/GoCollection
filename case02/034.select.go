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

	time.Sleep(time.Second)

	select{
	case <-ch:
		fmt.Println("<-ch")
	default:
		fmt.Println("default")
	}

}

