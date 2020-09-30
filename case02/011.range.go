package main

import "fmt"

func main() {
	arr := [5]int{0,1,2,3,4}
	for i,v := range arr{
		fmt.Println(i,v)
	}

	x := make(map[string]float32)
	x["z"] = 11.11
	x["y"] = 22.22
	x["r"] = 33.33
	for i,v := range x{
		fmt.Println(i,v)
	}


	r := "asd";
	for _,v := range r{
		fmt.Printf("%c\n",v)
	}


}
