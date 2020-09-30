package main

import "fmt"

func getSum(num []int) int {
	sum := 0;
	for _, v := range num {
		sum += v;
	}

	return sum;
}

func main() {
	num := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(getSum(num))
}
