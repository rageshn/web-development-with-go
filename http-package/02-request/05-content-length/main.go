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
		Method        string
		URL           *url.URL
		Submissions   map[string][]string
		Header        http.Header
		ContentLength int64
	}{
		req.Method,
		req.URL,
		req.Form,
		req.Header,
		req.ContentLength,
	}

	tpl.ExecuteTemplate(rw, "index.gohtml", data)
}

func main() {
	var r reqType
	http.ListenAndServe(":8080", r)

}
