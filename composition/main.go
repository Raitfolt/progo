package main

import (
	"composition/store"
	"fmt"
)

func main() {
	kayak := store.NewProduct("Kayak", "Watersports", 279)

	type OfferBundle struct {
		*store.SpecialDeal
		*store.Product
	}
	bundle := OfferBundle{store.NewSpecialDeal("Weekend Special", kayak, 50), store.NewProduct("Lifejacket", "Watersports", 48.95)}
	fmt.Println("Price:", bundle.Product.Price(0), bundle.SpecialDeal.Price(0))
}
