package main

import (
	"fmt"
	"net"
	"strings"
)

var onlineConns map[string]net.Conn
var messageQueue chan string
var quitChan chan bool

//收集消息
func Product(conn net.Conn) {
	defer conn.Close()

	buff := make([]byte, 1024)

	for {
		numOfBytes, _ := conn.Read(buff)
		if numOfBytes != 0 {
			message := string(buff[0:numOfBytes])
			messageQueue <- message
		}
	}
}

func Consume() {

	quitChan = make(chan bool)

	for {
		select {
		case message := <-messageQueue:
			contents := strings.Split(message, "#")
			if len(contents) > 1 {
				addr := contents[0]
				mess := contents[1]

				if conn, ok := onlineConns[addr]; ok {
					conn.Write([]byte(mess))
				}
			}
		case <-quitChan:
			break
		}
	}
}

func main() {
	onlineConns = make(map[string]net.Conn)
	messageQueue = make(chan string, 1000)

	listen, _ := net.Listen("tcp", "127.0.0.1:8080")
	defer listen.Close()

	// 一直发送
	go Consume()

	// 一直接受
	for {
		conn, _ := listen.Accept()

		link := fmt.Sprintf("%s", conn.RemoteAddr())

		onlineConns[link] = conn

		for i, _ := range onlineConns {
			fmt.Println(i)
		}

		go Product(conn)
	}
}
