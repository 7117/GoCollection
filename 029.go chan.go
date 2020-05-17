package main

import (
	"fmt"
	"strconv"
	"time"
)

func Read(ch chan int) {
	value := <-ch
	fmt.Print("value" + strconv.Itoa(value))
}
func Write(ch chan int) {
	ch <- 10
}

func main() {
	ch := make(chan int)
	go Read(ch)
	go Write(ch)
	time.Sleep(10)
	fmt.Print("end")

}
