package main

import (
	"composition/store"
	"fmt"
)

func main() {

	products := map[string]store.ItemForSale{
		"Kayak": store.NewBoat("Kayak", 279, 1, false),
		"Ball":  store.NewProduct("Soccer Ball", "Soccer", 19.50),
	}

	for key, p := range products {
		switch item := p.(type) {
		case store.Describable:
			fmt.Println(
				"Name:", item.GetName(),
				"Category:", item.GetCategory(),
				"Price:", item.Price(0.2))
		default:
			fmt.Println("Key:", key, "Price:", p.Price(0.2))
		}
	}

	kayak := store.NewProduct("Kayak", "Watersports", 279)

	type OfferBundle struct {
		*store.SpecialDeal
		*store.Product
	}
	bundle := OfferBundle{store.NewSpecialDeal("Weekend Special", kayak, 50), store.NewProduct("Lifejacket", "Watersports", 48.95)}
	fmt.Println("Price:", bundle.Product.Price(0), bundle.SpecialDeal.Price(0))
}
