package main

import "fmt"

func sendEmail() {
	// 相当于php的throw   发送邮件有错误进行报警panic
	panic(" error ")
}

func main() {

	// 相等于try
	defer func() {
		// 相等于catch  recover进行捕获错误
		if r := recover(); r != nil {
			fmt.Printf("执行修改数据表状态值的动作  END")
		}
	}()

	sendEmail()

}
