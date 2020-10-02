package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	x := make(map[string]int)

	x["a"]=1;
	x["b"]=2;

	d,_:= json.Marshal(x);

	fmt.Println(string(d))

}

