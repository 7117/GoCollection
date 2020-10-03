package main

import (
	"fmt"
	"github.com/astaxie/goredis"
)

func main() {
	var client goredis.Client

	client.Addr = "127.0.0.1:6379"

	res,err := client.Zadd("wwwa", []byte("bbb"),10)

	if err!=nil{
		panic(err)
	}

	fmt.Println(res)
}
