package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	fmt.Fprintln(rw, "Implements handler interface")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
