package main

import (
	"fmt"
)

type Product struct {
	name, category string
	price          float64
	*Supplier
}

type Supplier struct {
	name, city string
}

func newProduct(name, category string, price float64, supplier *Supplier) *Product {
	return &Product{name, category, price, supplier}
}

func copyProduct(product *Product) Product {
	p := *product
	s := *product.Supplier
	p.Supplier = &s
	return p
}

func main() {

	acme := &Supplier{"Acme Co", "New York"}

	products := [2]*Product{
		newProduct("Kayak", "Watersports", 275.0, acme),
		newProduct("Hat", "Skiing", 42.50, acme),
	}

	for _, p := range products {
		fmt.Println("Name:", p.name, "Supplier:", p.Supplier.name, p.Supplier.city)
	}

	fmt.Println("---------------------------")

	p1 := newProduct("Kayak", "Watersports", 275.0, acme)
	p2 := copyProduct(p1)

	p1.name = "Original Kayak"
	p1.Supplier.name = "BoatCo"

	for _, p := range []Product{*p1, p2} {
		fmt.Println("Name:", p.name, "Supplier:", p.Supplier.name, p.Supplier.city)
	}

	fmt.Println("---------------------------")

	var prod Product = Product{Supplier: &Supplier{}}
	var prodPtr *Product

	fmt.Println("Value:", prod.name, prod.category, prod.price, prod.Supplier.name)
	fmt.Println(&prod.Supplier)
	fmt.Println("Pointer:", prodPtr)

}
