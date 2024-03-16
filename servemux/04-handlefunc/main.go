package main

import (
	"io"
	"net/http"
)

func dog(rw http.ResponseWriter, req *http.Request) {
	io.WriteString(rw, "Doggo")
}

func cat(rw http.ResponseWriter, req *http.Request) {
	io.WriteString(rw, "Kitty")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/cat", cat)
	mux.HandleFunc("/dog", dog)

	http.ListenAndServe(":8080", mux)
}
