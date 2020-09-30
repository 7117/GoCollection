package main

import (
	// "fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello nihao"))
	})

	http.ListenAndServe("127.0.0.1:8080", nil)

}
