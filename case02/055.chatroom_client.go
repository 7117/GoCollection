package main

import (
	"bufio"
	"fmt"
	_ "github.com/astaxie/beego"
	"net"
	"os"
	"strings"
)

func MessageSend(conn net.Conn){

	var input string

	for{
		reader := bufio.NewReader(os.Stdin)

		data,_,_ :=reader.ReadLine()

		input =  string(data)

		if strings.ToUpper(input) == "EXIT"{
			conn.Close()
			break
		}
		conn.Write([]byte(input));
		fmt.Println("input success")


	}

}

func main() {

	conn,_:= net.Dial("tcp","127.0.0.1:8081")
	defer conn.Close()

	//一直在等待输入信息
	go MessageSend(conn)

	//这里的话一直在等待读取信息
	buf := make([]byte,1024)
	for{
		_,err := conn.Read(buf)
		if err != nil{
			fmt.Println("你已经退出")
			os.Exit(0)
		}
		fmt.Println("shoudao:"+ string(buf))
	}

	fmt.Println("client end")
}