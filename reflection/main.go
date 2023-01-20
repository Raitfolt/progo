package main

import (
	"reflect"
)

func getTypePath(t reflect.Type) (path string) {
	path = t.PkgPath()
	if path == "" {
		path = "(built-in)"
	}
	return
}

func printDetails(values ...interface{}) {
	for _, elem := range values {
		elemType := reflect.TypeOf(elem)
		Printfln("Name: %v, PkgPath: %v, Kind: %v",
			elemType.Name(), getTypePath(elemType), elemType.Kind())
	}
}

type Payment struct {
	Currency string
	Amount   float64
}

type Payment2 struct {
	Payment
	Product
	Customer
}

func main() {
	product := Product{
		Name: "Kayak", Category: "Watersports", Price: 279,
	}
	customer := Customer{Name: "Alice", City: "New York"}
	payment := Payment{Currency: "USD", Amount: 100.50}
	p2 := Payment2{Payment: payment, Product: product, Customer: customer}
	printDetails(product, customer, payment, 10, true, p2)
}
