package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()
	fmt.Println(currentLocation)

	//time.Time格式
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	firstOfYear := time.Date(currentYear, time.January, 1, 0, 0, 0, 0, currentLocation)
	lastOfYear := firstOfMonth.AddDate(1, 0, -1)
	fmt.Println(firstOfMonth)
	fmt.Println(lastOfMonth)
	fmt.Println(firstOfYear)
	fmt.Println(lastOfYear)

	//时间戳格式（time.Time格式后面加.Unix()就行了）
	intType := firstOfMonth.Unix()
	fmt.Println(intType)
}

//2020-10-19 17:20:12.5198564 +0800 CST m=+0.006832101
//Local
//2020-10-01 00:00:00 +0800 CST
//2020-10-31 00:00:00 +0800 CST
//2020-01-01 00:00:00 +0800 CST
//2021-09-30 00:00:00 +0800 CST
//1601481600
