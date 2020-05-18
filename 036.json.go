package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	x := [5]int{1,2,3,4,5}

	s,err := json.Marshal(x)

	if err!= nil {
		panic("error");
	}
}
