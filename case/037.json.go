package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	x := [5]int{1, 2, 3, 4, 5}

	s, _ := json.Marshal(x)

	fmt.Println(string(s))

	y := make(map[string]float64)

	y["a"] = 1.1
	y["b"] = 2.2

	q, _ := json.Marshal(y)

	fmt.Println(string(q))

}
