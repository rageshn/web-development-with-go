package main

import (
	"fmt"
	"math"
	"os"
	"text/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"fdbl":  double,
	"fsq":   square,
	"fsqrt": sqRoot,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func double(i int) int {
	return i + i
}

func square(i int) int {
	return i * i
}

func sqRoot(i float64) float64 {
	return math.Sqrt(i)
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 3)
	if err != nil {
		fmt.Println("Error: ", err)
	}

}
