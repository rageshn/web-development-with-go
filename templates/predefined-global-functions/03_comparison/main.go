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

	s1 := struct {
		Score1 int
		Score2 int
	}{
		7,
		9,
	}
	err := tpl.Execute(os.Stdout, s1)
	if err != nil {
		fmt.Println("Error: ", err)
	}

}
