package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
)

func foo(rw http.ResponseWriter, req *http.Request) {
	io.WriteString(rw, "Foo ran")
}

func goku(rw http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("goku.gohtml")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	tpl.ExecuteTemplate(rw, "goku.gohtml", nil)

}

func gokuPic(rw http.ResponseWriter, req *http.Request) {
	http.ServeFile(rw, req, "goku.jpeg")
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/goku", goku)
	http.HandleFunc("/gokupic", gokuPic)
	http.ListenAndServe(":8080", nil)
}
