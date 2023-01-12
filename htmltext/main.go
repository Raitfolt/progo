package main

import (
	"os"
	"text/template"
)

func Exec(t *template.Template) error {
	return t.Execute(os.Stdout, Products)
}

func main() {
	allTemplates, err := template.ParseGlob("templates/*.html")
	if err == nil {
		selectedTemplated := allTemplates.Lookup("mainTemplate")
		err = Exec(selectedTemplated)
		os.Stdout.WriteString("\n")
	}
	if err != nil {
		Printfln("Error: %v", err.Error())
	}
}
