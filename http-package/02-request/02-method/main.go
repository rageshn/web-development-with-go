package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
)

var tpl *template.Template

type reqType int

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func (r reqType) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	data := struct {
		Method      string
		Submissions url.Values
	}{
		req.Method,
		req.Form,
	}

	tpl.ExecuteTemplate(rw, "index.gohtml", data)
}

func main() {
	var r reqType
	http.ListenAndServe(":8080", r)

}
