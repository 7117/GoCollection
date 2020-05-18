package main

import (
	"fmt"
	"regexp"
)

func main() {
	a := []byte("abcdef|abcdefghijk")
	b := "abcdefghijk"
	reg1 := regexp.MustCompilePOSIX(b)
	// PS D:\gopath\src\goPractice> go run .\045.regexp.go
	// abcdef
	// 小的找大的  大的是一个字符型切片
	fmt.Printf("%v\n", string(reg1.Find(a)))

}
