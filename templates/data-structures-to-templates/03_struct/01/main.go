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
	malekith := boss{
		Location: "Crumbling Farum Azula",
		Name:     "Malekith",
	}

	err := tpl.Execute(os.Stdout, malekith)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
