package main

import "fmt"

func main() {
	//有些类似于栈的感觉
	for i := 1; i < 4; i++ {
		defer fmt.Println(i)
	}

	defer func() {
		fmt.Println("after")
	}()

	fmt.Println("before")
}
