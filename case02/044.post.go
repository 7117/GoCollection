package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	res, _ := http.Post("http://www.baidu.com", "application/x-www-form-urlencoded", strings.NewReader("info"))
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
