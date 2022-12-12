package main

import (
	"fmt"
	"packages/store"
)

func main() {
	product := store.Product{
		Name:     "Kayak",
		Category: "Watersports",
	}
	fmt.Println("Name:", product.Name)
	fmt.Println("Category:", product.Category)
}
