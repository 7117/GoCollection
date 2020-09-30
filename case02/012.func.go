package main

import "fmt"

func main() {
	a,b :=swap(1,2);
	fmt.Println(a,b)

	c :=add(1,4);
	fmt.Println(c)
}


func swap(a int,b int)(int,int){
	return b,a;
}

func add(a,b int)(int){
	return a+b;
}