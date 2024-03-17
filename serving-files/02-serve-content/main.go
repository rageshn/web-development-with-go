package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func goku(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(rw, `<img src="/goku.jpeg">`)
}

func gokuPic(rw http.ResponseWriter, req *http.Request) {
	f, err := os.Open("goku.jpeg")
	if err != nil {
		fmt.Println("Error: ", err)
		http.Error(rw, "file not found", 404)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(rw, "file not found", 404)
	}
	http.ServeContent(rw, req, f.Name(), fi.ModTime(), f)
}

func main() {
	http.HandleFunc("/", goku)
	http.HandleFunc("/goku", gokuPic)

	http.ListenAndServe(":8080", nil)
}
