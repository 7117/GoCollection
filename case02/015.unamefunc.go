package main

import "fmt"

func main() {
	f:= func(x,y int)int{
		return x+y;
	}

	res := f(2,3)

	fmt.Println(res)
}
