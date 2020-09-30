package main

import "fmt"

func main() {
	x := 3;
	switch x {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("default")
	}

	r := 1;
	switch r {
	case 1:
		//会进行顺序执行
		fallthrough
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("default")
	}

}
