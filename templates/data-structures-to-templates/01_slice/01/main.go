package main

import (
	"fmt"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	bosses := []string{"Godrick", "Mohg Lord of Blood", "Godfrey", "Malekith"}
	err := tpl.Execute(os.Stdout, bosses)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
