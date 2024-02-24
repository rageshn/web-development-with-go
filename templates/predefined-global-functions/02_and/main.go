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
	Optional bool
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	fireGiant := boss{
		Location: "Mountaintops of the Giants",
		Name:     "Fire Giant",
		Optional: false,
	}

	godskinDuo := boss{
		Location: "Crumbling Farum Azula",
		Name:     "Godskin Duo",
		Optional: false,
	}

	rykard := boss{
		Location: "Volcano Manor",
		Name:     "Rykard, Lord of Blasphemy",
		Optional: true,
	}

	radahn := boss{
		Location: "Caelid",
		Name:     "Starscourge Radahn",
		Optional: true,
	}

	bosses := []boss{fireGiant, godskinDuo, rykard, radahn}

	err := tpl.Execute(os.Stdout, bosses)
	if err != nil {
		fmt.Println("Error: ", err)
	}

}
