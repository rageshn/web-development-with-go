package main

import (
	"fmt"
	"html/template"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", "Data passed to variable")
	if err != nil {
		fmt.Println("Error while executing template: ", err)
	}
}
