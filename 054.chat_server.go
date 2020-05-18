package main

import (
	"fmt"
	"net"
)

func Deal(conn net.Conn) {
	defer conn.Close()

	buff := make([]byte, 1024)

	for {
		numOfBytes, _ := conn.Read(buff)
		if numOfBytes != 0 {
			remoteAddr := conn.RemoteAddr()
			fmt.Print(remoteAddr)
			fmt.Printf(" received mess：%s\n", string(buff[0:numOfBytes]))
		}
	}
}

func main() {

	listen, _ := net.Listen("tcp", "127.0.0.1:8080")
	defer listen.Close()

	for {
		conn, _ := listen.Accept()
		go Deal(conn)
	}
}
