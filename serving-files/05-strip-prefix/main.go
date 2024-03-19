package main

import (
	"io"
	"net/http"
)

func getAssets(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(rw, `<img src="goku.jpeg">`)
}

func main() {
	http.HandleFunc("/", getAssets)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}
