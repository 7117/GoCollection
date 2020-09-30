package main

import (
	"encoding/json"
	"fmt"
)

type Stu struct {
	Name string `json:"stuname"`
	Age  int
}

func main() {
	stu := Stu{"a", 20}

	res, _ := json.Marshal(stu)

	fmt.Println(string(res))
}
