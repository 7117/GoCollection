package main

import (
	"fmt"
	"time"
)

func AddTo(x, y int) {
	z := x + y
	fmt.Println(z)
}

func main() {

	for i := 1; i < 10; i++ {
		go AddTo(i, i)
	}

	time.Sleep(2 * time.Second)
}
