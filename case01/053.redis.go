package main

import (
	"github.com/astaxie/goredis"
)

func main() {

	var client goredis.Client
	client.Addr = "127.0.0.1:6379"

	client.Zadd("zadd", []byte("bei"), 20)
}
