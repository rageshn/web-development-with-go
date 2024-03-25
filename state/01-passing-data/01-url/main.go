package main

import (
	"io"
	"net/http"
)

func foo(rw http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	io.WriteString(rw, "Search for "+v)
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
