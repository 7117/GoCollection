package main

import (
	"fmt"
	"regexp"
)

func main() {
	v, _ := regexp.Match("[a-zA-Z]{3}", []byte("hl"))

	fmt.Println(v)
