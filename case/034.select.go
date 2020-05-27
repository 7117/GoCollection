package main

import (
	"fmt"
	"time"
)

// 超时控制
func main() {
	ch := make(chan int)
	timeout := make(chan int, 1)

	go func() {
		time.Sleep(time.Second)
		timeout <- 1
	}()

	// 此处用来设置ch的值
	// go func(ch chan int){
	// 	ch<-1;
	// }(ch)

	select {
	case <-ch:
		fmt.Print("ch")
	case <-timeout:
		fmt.Print("time out")
	}

	fmt.Print("end code")

}
