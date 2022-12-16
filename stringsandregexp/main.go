package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	product := "Kayak"

	for _, char := range product {
		fmt.Println(string(char), "Upper case:", unicode.IsUpper(char))
	}

	fmt.Println("Contains:", strings.Contains(product, "yak"))
	fmt.Println("ContainsAny:", strings.ContainsAny(product, "abc"))
	fmt.Println("ContainsRune:", strings.ContainsRune(product, 'K'))
	fmt.Println("EqualFold:", strings.EqualFold(product, "KAYAK"))
	fmt.Println("HasPrefix:", strings.HasPrefix(product, "Ka"))
	fmt.Println("HasSuffix:", strings.HasSuffix(product, "yak"))

	kayak := "KaYaK iS gOoD"
	fmt.Println(kayak, "to lower:", strings.ToLower(kayak))
	fmt.Println(kayak, "to upper:", strings.ToUpper(kayak))
	fmt.Println(kayak, "title:", strings.Title(kayak))
	fmt.Println(kayak, "to title:", strings.ToTitle(kayak))

}
