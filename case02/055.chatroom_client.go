package main

import (
	"fmt"
	"net"
)


func main() {

	conn,_:= net.Dial("tcp","127.0.0.1:8081")
	defer conn.Close()

	conn.Write([]byte("hello"))

	fmt.Println("had send")
}