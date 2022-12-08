package main

import "fmt"

type Product struct {
	name, category string
	price          float64
}

func printDetails(product *Product) {
	fmt.Println("Name:", product.name, "Category:", product.category, "Price", product.price)
}

func main() {
	products := []*Product{
		{"Kayak", "Watersports", 275},
		{"Lifejacket", "Watersports", 48.95},
		{"Soccer Ball", "Soccer", 19.5},
	}

	for _, p := range products {
		printDetails(p)
	}
}

//355
