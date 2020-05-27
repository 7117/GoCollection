package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		for i := 1; i < 10; i++ {
			if i == 5 {
				runtime.Gosched()
			}
			fmt.Println(i)
		}
	}()

	go func() {
		for i := 10; i <= 20; i++ {
			fmt.Println(i)
		}
	}()

	time.Sleep(2 * time.Second)
}
