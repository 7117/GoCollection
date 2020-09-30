package main

import (
	"fmt"
	"time"
)

func test() {
	fmt.Print("test")
}

func main() {
	go test()

	time.Sleep(2)
}
