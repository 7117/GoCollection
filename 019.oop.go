package main

import "fmt"

type Interger struct {
}

// 这里的self  前面的别名  其实就相当于$this
func (self Interger) compare(a int, b int) (c bool) {
	return a < b
}

func main() {
	self := new(Interger)
	a := 1
	b := 2

	fmt.Printf("%v", self.compare(a, b))
	fmt.Printf("%v", self.compare(b, a))

}
