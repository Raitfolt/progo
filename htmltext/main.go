package main

import (
	"html/template"
	"os"
	"strings"
)

func GetCategories(products []Product) (categories []string) {
	catMap := map[string]string{}
	for _, p := range products {
		if catMap[p.Category] == "" {
			catMap[p.Category] = p.Category
			categories = append(categories, p.Category)
		}
	}
	return
}

func Exec(t *template.Template) error {
	return t.Execute(os.Stdout, Products)
}

func main() {
	allTemplates := template.New("allTemplates")
	allTemplates.Funcs(map[string]interface{}{
		"getCats": GetCategories,
		"lower":   strings.ToLower,
	})
	allTemplates, err := allTemplates.ParseGlob("templates/*.html")
	if err == nil {
		selectedTemplated := allTemplates.Lookup("mainTemplate")
		err = Exec(selectedTemplated)
		os.Stdout.WriteString("\n")
	}
	if err != nil {
		Printfln("Error: %v", err.Error())
	}
}
