package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

var onlineConns map[string]net.Conn
var messageQueue chan string
var quitChan chan bool

const (
	LUJING = "./t.log"
)

func checkError(err error){
	if err != nil{
		fmt.Println(err)
	}
}

func ProcessInfo(conn net.Conn){
	buff := make([]byte,1024)
	defer func(conn net.Conn) {
		addr := fmt.Sprintf("%s", conn.RemoteAddr())
		delete(onlineConns,addr)

		for i,_ :=range onlineConns{
			fmt.Println("end all"+i)
		}

		conn.Close()

	}(conn)

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
		sendMessage := strings.Join(contents[1:],"#")

		addr = strings.Trim(addr," ")

		if conn,ok :=onlineConns[addr];ok{
			conn.Write([]byte(sendMessage))
		}

	}else{
		contents := strings.Split(message,"*")

		if strings.ToUpper(contents[1]) == "LIST"{
			ips := ""
			for i,_ :=range onlineConns{
				ips = ips +"---" +i
			}

			addr := contents[0]
			addr = strings.Trim(addr," ")

			if conn,ok :=onlineConns[addr];ok{
				conn.Write([]byte(ips))
				fmt.Println("list info")
			}

		}
	}
}

func main() {

	LogFile,_ := os.OpenFile(LUJING,os.O_RDWR | os.O_CREATE,0);
	defer LogFile.Close()

	logger := log.New(LogFile,"\r\n",log.Ldate|log.Ltime|log.Llongfile)
	logger.Println("logloglog")

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
