package main

import (
	"io"
	"net/http"
)

type dog int

func (d dog) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	io.WriteString(rw, "Doggo")
}

type cat string

func (c cat) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	io.WriteString(rw, "Kitty")
}

func main() {
	var d dog
	var c cat

	http.Handle("/dog", d)
	http.Handle("/cat", c)

	http.ListenAndServe(":8080", nil)
}
