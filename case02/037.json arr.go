package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	x := [5]int{1,2,3,4,5};

	s,err := json.Marshal(x)

	if err!=nil{
		panic("error")
	}

	fmt.Println(string(s))

}

