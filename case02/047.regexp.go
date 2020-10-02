package main

import (
	"fmt"
	"regexp"
)

func main() {
	a := []byte("abcdef|abcdefghijk")
	b := "abcdefghijk"
	reg1 := regexp.MustCompilePOSIX(b)
	//D:\gopath\src\goPractice\case02>go run 047.regexp.go
	//		abcdefghijk
	// 小的找大的  大的是一个字符型切片
	//MustCompilePOSIX最长匹配
	fmt.Printf("%v\n", string(reg1.Find(a)))

}
