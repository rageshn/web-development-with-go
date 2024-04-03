package main

import (
	"io"
	"net/http"
	"strconv"
)

func incrementCookie(rw http.ResponseWriter, req *http.Request) {
	ck, err := req.Cookie("mycookie")
	if err == http.ErrNoCookie {
		ck = &http.Cookie{
			Name:  "mycookie",
			Value: "0",
		}
	}
	count, _ := strconv.Atoi(ck.Value)
	count++
	ck.Value = strconv.Itoa(count)

	http.SetCookie(rw, ck)
	io.WriteString(rw, ck.Value)
}

func main() {
	http.HandleFunc("/", incrementCookie)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
