package main

import (
	"fmt"
	"time"
)

var ch chan int

func main() {

	ch = make(chan int)

	go func() {
		for i := 1; i < 10; i++ {
			if i == 5 {
				<-ch
			}
			fmt.Println(i)
		}
	}()

	go func() {
		for i := 10; i <= 20; i++ {
			fmt.Println(i)
		}
		ch <- 1
	}()

	time.Sleep(2 * time.Second)
}
