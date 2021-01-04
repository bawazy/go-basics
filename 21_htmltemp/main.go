package main

import (
	"html/template"
	"log"
	"net/http"
)

//Person struct
type Person struct {
	Name string
}

func main() {
	http.HandleFunc("/", helloworld)
	http.ListenAndServe(":80", nil)
}
func helloworld(w http.ResponseWriter, r *http.Request) {
	person := Person{
		Name: "Hayatu",
	}

	tmpl := template.New("html template")

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal("parsefile: ", err)
	}

	err1 := tmpl.Execute(w, person)
	if err != nil {
		log.Fatal("execute: ", err1)
	}
}
