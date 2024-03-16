package main

import (
	"io"
	"net/http"
)

func defaultRoute(rw http.ResponseWriter, req *http.Request) {
	io.WriteString(rw, "default page")
}

func dog(rw http.ResponseWriter, req *http.Request) {
	io.WriteString(rw, "Doggo")
}

func me(rw http.ResponseWriter, req *http.Request) {
	io.WriteString(rw, "Ragesh")
}

func main() {
	http.HandleFunc("/", defaultRoute)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/me", me)

	http.ListenAndServe(":8080", nil)
}
