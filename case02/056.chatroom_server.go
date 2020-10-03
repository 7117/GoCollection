package main

import (
"fmt"
"net"
)

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
			fmt.Println("process info" + string(numOfBytes))
		}
	}

}

func main() {

	listen_socker,_ := net.Listen("tcp","127.0.0.1:8081")
	defer listen_socker.Close()

	fmt.Println("is waiting")

	for{
		conn,_ := listen_socker.Accept();

		go ProcessInfo(conn)
	}
}
