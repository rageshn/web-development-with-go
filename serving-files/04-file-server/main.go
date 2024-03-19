package main

import (
	"io"
	"net/http"
)

func goku(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(rw, `<img src="goku.jpeg">`)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/goku", goku)
	http.ListenAndServe(":8080", nil)
}
