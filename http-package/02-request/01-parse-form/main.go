package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type reqType int

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func (r reqType) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

func main() {
	var rt reqType
	http.ListenAndServe(":8080", rt)
}
