package main

import (
	"fmt"
	"os"
	"text/template"
)

var tpl *template.Template

type boss struct {
	Location string
	Name     string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	godfrey := boss{
		Location: "Leyndell, Royal Capital",
		Name:     "Godfrey",
	}

	err := tpl.Execute(os.Stdout, godfrey)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
