package main

import "fmt"

type Interger struct {
	value int
}

func (a Interger) compare(b Interger) bool {
	return a.value < b.value;
}

func main() {

	a := Interger{1};
	b := Interger{2};
	fmt.Println(a.compare(b))
}
