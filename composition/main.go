package main

import (
	"composition/store"
	"fmt"
)

func main() {
	boats := []*store.Boat{
		store.NewBoat("Kayak", 275, 1, false),
		store.NewBoat("Canoe", 400, 3, false),
		store.NewBoat("Tender", 650.25, 2, true),
	}
	for _, b := range boats {
		fmt.Println("Boat:", b.Name, "Price:", b.Price(0.2))
	}

	rentals := []*store.RentalBoat{
		store.NewRentalBoat("Rubber Ring", 10, 1, false, false, "N/A", "N/A"),
		store.NewRentalBoat("Yacht", 50000, 5, true, true, "Bob", "Alice"),
		store.NewRentalBoat("Super Yacht", 100000, 15, true, true, "Dora", "Charlie"),
	}
	for _, r := range rentals {
		fmt.Println("Rental boat:", r.Name, "Rental Price:", r.Price(0.2), "Capitan:", r.Capitan)
	}

	fmt.Println("------------------------------------------")

	product := store.NewProduct("Kayak", "Watersports", 279)
	deal := store.NewSpecialDeal("Weekend Special", product, 50)

	Name, price, Price := deal.GetDetails()

	fmt.Println("Name:", Name)
	fmt.Println("Price field:", price)
	fmt.Println("Price method:", Price)
}
