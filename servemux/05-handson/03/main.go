package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
)

func defaultRoute(rw http.ResponseWriter, req *http.Request) {
	io.WriteString(rw, "default page")
}

func dog(rw http.ResponseWriter, req *http.Request) {
	io.WriteString(rw, "Doggo")
}

func me(rw http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("test.gohtml")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	tpl.ExecuteTemplate(rw, "test.gohtml", "Ragesh")
}

func main() {
	http.HandleFunc("/", http.HandlerFunc(defaultRoute))
	http.HandleFunc("/dog", http.HandlerFunc(dog))
	http.HandleFunc("/me", http.HandlerFunc(me))

	http.ListenAndServe(":8080", nil)
}
