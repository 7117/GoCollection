package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string `json:"aa"`
	Age int `json:"bb"`
}

func main() {

	stu := Student{"a", 20}

	res, err := json.Marshal(stu);
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(res))
}

