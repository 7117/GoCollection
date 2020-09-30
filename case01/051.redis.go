package main

import (
	"fmt"

	"github.com/astaxie/goredis"
)

func main() {

	var client goredis.Client
	client.Addr = "127.0.0.1:6379"

	client.Set("redis", []byte("hello redis"))

	res, _ := client.Get("redis")

	fmt.Println(string(res))

}
