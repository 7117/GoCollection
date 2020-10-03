package main

import (
	"fmt"
	"github.com/astaxie/goredis"
)

func main() {
	var client goredis.Client

	client.Addr = "127.0.0.1:6379"

	mapping := make(map[string]interface{})

	mapping["a"] = 11;
	mapping["b"] = "22";
	mapping["c"] = "aa";

	client.Hmset("ggg",mapping)

	res,_:=client.Hget("ggg","a");

	fmt.Println(string(res))
}
