package main

import "fmt"

type Interger struct {
	value int
}

func (a Interger) compare(c int) bool {
	return a.value < c;
}

func main() {

	a := Interger{1};
	fmt.Println(a.compare(2))
}
