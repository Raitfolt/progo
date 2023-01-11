package main

import (
	"os"
	"text/template"
)

func main() {
	allTemplates, err := template.ParseFiles("templates/template.html", "templates/extras.html")
	if err == nil {
		allTemplates.ExecuteTemplate(os.Stdout, "template.html", &Kayak)
		os.Stdout.WriteString("\n")
		allTemplates.ExecuteTemplate(os.Stdout, "extras.html", &Kayak)
		os.Stdout.WriteString("\n")
	} else {
		Printfln("Error: %v", err.Error())
	}
}
