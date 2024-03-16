package main

import (
	"io"
	"net/http"
)

type reqType int

func (r reqType) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(rw, "Doggo")
	case "/cat":
		io.WriteString(rw, "Kitty")
	}
}

func main() {
	var rt reqType
	http.ListenAndServe(":8080", rt)
}
