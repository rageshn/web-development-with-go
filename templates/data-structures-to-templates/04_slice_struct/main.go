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

	err := tpl.Execute(os.Stdout, bosses)
	if err != nil {
		fmt.Println("Error: ", err)
	}

}
