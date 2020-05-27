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
	// json
	stu := Stu{"a", 20}
	json_type, _ := json.Marshal(stu)
	fmt.Printf(string(json_type))

	// unjson
	var arr_type interface{}
	json.Unmarshal([]byte(json_type), &arr_type)
	fmt.Printf("%v", arr_type)

}
