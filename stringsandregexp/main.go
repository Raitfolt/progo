package main

import (
	"fmt"
	"regexp"
)

func getSubstring(s string, indices []int) string {
	return string(s[indices[0]:indices[1]])
}

func main() {
	pattern := regexp.MustCompile("A (?P<type>[A-z]*) for (?P<capacity>[A-z]*) person")

	description := "Kayak. A boat for one person."

	replaced := pattern.ReplaceAllStringFunc(description, func(s string) string {
		return "This is the replacement content"
	})
	fmt.Println(replaced)
}
