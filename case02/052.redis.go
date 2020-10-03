package main

import (
	"fmt"
	"github.com/astaxie/goredis"
)

func main() {
	var client goredis.Client
	client.Addr = "127.0.0.1:6379"

	client.Set("test",[]byte("hello,befeng"))

	res,_ :=client.Get("test");

	fmt.Printf(string(res))
}
