package main

import (
	"fmt"
	"net/http"
)

func set(rw http.ResponseWriter, req *http.Request) {
	ck := http.Cookie{
		Name:  "mycookie",
		Value: "cookie Value",
	}
	http.SetCookie(rw, &ck)

	fmt.Fprintln(rw, "Cookie written in browser")
}

func read(rw http.ResponseWriter, req *http.Request) {
	ck, err := req.Cookie("mycookie")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Fprintln(rw, "cookie value: ", ck)
}

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
