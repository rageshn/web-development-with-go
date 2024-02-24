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

type car struct {
	Manufacturer string
	Model        string
	Doors        string
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

	tata := car{
		Manufacturer: "Tata Motors",
		Model:        "Tata Nexon",
		Doors:        "5",
	}
	ford := car{
		Manufacturer: "Ford",
		Model:        "Ecosport",
		Doors:        "5",
	}

	bosses := []boss{fireGiant, godskinDuo}
	cars := []car{tata, ford}

	data := struct {
		Bosses    []boss
		Transport []car
	}{
		bosses,
		cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
