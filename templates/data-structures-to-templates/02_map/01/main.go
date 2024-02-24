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

	bosses := map[string]string{
		"Stormveil Castle":        "Godrick",
		"Mohgwyn Palace":          "Mohg Lord of Blood",
		"Leyndell, Royal Capital": "Godfrey",
		"Crumbling Farum Azula":   "Malekith",
	}

	err := tpl.Execute(os.Stdout, bosses)
	if err != nil {
		fmt.Println("Error: ", err)
	}

}
