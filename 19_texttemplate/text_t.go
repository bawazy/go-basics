package main

import (
	"log"
	"os"
	"text/template"
)

type Student struct {
	Name string
}

func main() {

	s := Student{Name: "Hayatu"}

	tmpl := template.New("test")

	tmpl, err := tmpl.Parse("Hello {{.Name}}!")
	if err != nil {
		log.Fatal("Parse: ", err)
	}

	err1 := tmpl.Execute(os.Stdout, s)
	if err1 != nil {
		log.Fatal("Execute: ", err1)
	}
}
