package main

import (
	"fmt"
)

func main() {
	chs := make([]chan int,10)


	for i:=0;i<10;i++{
		chs[i] = make(chan int);
		go adds(i,i,chs[i])
	}

	for _, v := range chs {
		<-v
	}

}

func adds(x int,y int,quit chan int){
	z:= x+y;
	fmt.Println(z)
	quit <- 1;
}



