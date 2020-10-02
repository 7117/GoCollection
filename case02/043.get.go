package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	res,_ :=http.Get("http://www.baidu.com")
	defer res.Body.Close()

	resbody,_ := ioutil.ReadAll(res.Body)

	fmt.Println(string(resbody))
}
