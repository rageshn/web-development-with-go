package main

import (
	"fmt"
	"net/http"
)

type reqType int

func (r reqType) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("ragesh-key", "This is custom header")
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(rw, "<h1>Any code in this function</h1>")
}

func main() {
	var r reqType
	http.ListenAndServe(":8080", r)

}
