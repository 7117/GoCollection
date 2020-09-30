package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		defer fmt.Print(i)
	}
}
