package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

type boss struct {
	Location string
	Name     string
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func main() {
	fireGiant := boss{
		Location: "Mountaintops of the Giants",
		Name:     "Fire Giant",
	}

	godskinDuo := boss{
		Location: "Crumbling Farum Azula",
		Name:     "Godskin Duo",
	}

	rykard := boss{
		Location: "Volcano Manor",
		Name:     "Rykard, Lord of Blasphemy",
	}

	radahn := boss{
		Location: "Caelid",
		Name:     "Starscourge Radahn",
	}

	bosses := []boss{fireGiant, godskinDuo, rykard, radahn}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", bosses)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
