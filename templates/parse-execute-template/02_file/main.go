package main

import (
	"fmt"
	"html/template"
	"os"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		fmt.Println("Error while parsing file: ", err)
	}

	f, err := os.Create("index.html")
	if err != nil {
		fmt.Println("Error while creating file: ", err)
	}
	defer f.Close()

	err = tpl.Execute(f, nil)
	if err != nil {
		fmt.Println("Error while executing: ", err)
	}
}
