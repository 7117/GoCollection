package main

import (
	"github.com/astaxie/goredis"
)

func main() {

	var client goredis.Client
	client.Addr = "127.0.0.1:6379"

	f := make(map[string]interface{})

	f["name"] = "zhangsan"
	f["age"] = 12
	f["sex"] = "女"

	client.Hmset("hash", f)

}
