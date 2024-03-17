package main

import (
	"io"
	"net/http"
)

func goku(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(rw, `<img src="/goku.jpeg">`)
}

func gokuPic(rw http.ResponseWriter, req *http.Request) {
	http.ServeFile(rw, req, "goku.jpeg")
}

func main() {
	http.HandleFunc("/", goku)
	http.HandleFunc("/goku", gokuPic)

	http.ListenAndServe(":8080", nil)
}
