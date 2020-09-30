package main

import "fmt"

func getSum(num []int) int {
	sum := 0
	for i := 0; i < len(num); i++ {
		sum += num[i]
	}

	return sum
}

func main() {

	num := []int{1, 2, 3, 4, 5}
	res := getSum(num)
	fmt.Print(res)
}
