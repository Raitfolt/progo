package main

import (
	"os"
	"text/template"
)

func main() {
	t1, err1 := template.ParseFiles("templates/template.html")
	t2, err2 := template.ParseFiles("templates/extras.html")
	if err1 == nil && err2 == nil {
		t1.Execute(os.Stdout, &Kayak)
		os.Stdout.WriteString("\n")
		t2.Execute(os.Stdout, &Kayak)
		os.Stdout.WriteString("\n")
	} else {
		Printfln("Error: %v %v", err1.Error(), err2.Error())
	}
}
