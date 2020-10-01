package main

import (
	"fmt"
	"time"
)

func max() {
	fmt.Println("info")
}

func main() {

	go max();
	time.Sleep(1)
}
