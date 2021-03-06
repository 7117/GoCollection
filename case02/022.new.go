package main

import "fmt"

type Point struct {
	px float32
	py float32
}

func (point *Point) setxy(px, py float32) {
	point.px = px;
	point.py = py;
}

func (point *Point) getxy() (px, py float32) {
	return point.px, point.py;
}

func main() {
	obj := new(Point)
	obj.setxy(1.1, 2.2);
	a, b := obj.getxy();
	fmt.Println(a, b)

	//直接使用类的实例化不行  要先实例化才可以！
	//point.setxy(1.11,2.22);
	//a,b := point.getxy();
	//fmt.Println(a,b);
}
