package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func send(conn net.Conn) {
	var input string

	for {
		reader := bufio.NewReader(os.Stdin)
		data, _, _ := reader.ReadLine()
		input = string(data)

		// 退出
		if strings.ToUpper(input) == "EXIT" {
			conn.Close()
			break
		}

		// 写入到连接  进行发送
		conn.Write([]byte(input))
	}
}

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8080")
	defer conn.Close()

	//一直发送
	go send(conn)

	//一直接受
	buf := make([]byte, 1024)
	for {
		numOfBytes, _ := conn.Read(buf)

		fmt.Println("receive server message content:" + string(buf[0:numOfBytes]))
	}

}
