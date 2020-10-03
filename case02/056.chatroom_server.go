package main

import (
	"fmt"
	"net"
	"strings"
)

var onlineConns map[string]net.Conn
var messageQueue chan string
var quitChan chan bool

func checkError(err error){
	if err != nil{
		fmt.Println(err)
	}
}

func ProcessInfo(conn net.Conn){
	buff := make([]byte,1024)
	defer conn.Close()

	for{
		numOfBytes,_ := conn.Read(buff)

		if numOfBytes != 0 {
			message := string(buff[0:numOfBytes])
			messageQueue<-message
		}
	}

}

func ConsumerMessage(){

	quitChan = make(chan bool)

	for{
		select{
		case message := <-messageQueue:
		//	处理消息
			doProcessMessage(message)
		case <-quitChan:
			break;
		}
	}
}

func doProcessMessage(message string)  {
	contents := strings.Split(message,"#")

 	if len(contents)>1{
		addr := contents[0]
		sendMessage := contents[1]

		addr = strings.Trim(addr," ")

		if conn,ok :=onlineConns[addr];ok{
			conn.Write([]byte(sendMessage))
		}

	}
}

func main() {

	onlineConns = make(map[string]net.Conn)
	messageQueue = make(chan string, 1000)

	listen_socker,_ := net.Listen("tcp","127.0.0.1:8081")
	defer listen_socker.Close()

	fmt.Println("is waiting")

	go ConsumerMessage()

	//一旦连接上  就进行读取信息
	for{
		conn,_ := listen_socker.Accept();

		link := fmt.Sprintf("%s", conn.RemoteAddr())
		onlineConns[link] = conn

		for i, _ := range onlineConns {
			fmt.Println(i)
		}

		go ProcessInfo(conn)
	}
}
