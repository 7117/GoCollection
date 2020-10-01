package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	var ch chan int
	ch = make(chan int)

	go read(ch)
	go write(ch)


	time.Sleep(2 * time.Second)
}

func read(ch chan int){
	value := <-ch
	fmt.Println("value is : " + strconv.Itoa(value) )
}

func write(ch chan int){
	ch <- 10
}