package main

import (
	"log"
	"os"
	"text/template"
)

type Person struct {
	Name   string
	Emails []string
}

const tmpl = `Helloo my name is{{.Name}} and {{range .Emails}} my email is {{.}}
{{end}}
`

func main() {
	person := Person{
		Name:   "Hayatu Bayero",
		Emails: []string{"hbawazy@gmail.com", "bayerowh@tetfund.gov.ng"},
	}
	t := template.New("person template")

	t, err := t.Parse(tmpl)
	if err != nil {
		log.Fatal("parse: ", err)
	}
	err = t.Execute(os.Stdout, person)
	if err != nil {
		log.Fatal("Execute: ", err)
	}
}
