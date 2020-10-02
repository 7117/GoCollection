package main

import (
	"fmt"
	"regexp"
)

func main() {
	value, _ := regexp.Match("[a-zA-Z]{3}", []byte("ha1q"))

	fmt.Println(value)

}
