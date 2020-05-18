package main

import (
	"io/ioutil"
	"net/http"
	"fmt"
)

func main() {
	res, _ := http.Get("http://www.baidu.com")
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
