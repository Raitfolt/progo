package main

import (
	"fmt"
	"reflect"
	"strings"
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
		fieldDetails := []string{}
		elemType := reflect.TypeOf(elem)
		elemValue := reflect.ValueOf(elem)
		if elemType.Kind() == reflect.Struct {
			for i := 0; i < elemType.NumField(); i++ {
				fieldName := elemType.Field(i).Name
				fieldVal := elemValue.Field(i)
				fieldDetails = append(fieldDetails, fmt.Sprintf("%v: %v", fieldName, fieldVal))
			}
			Printfln("%v: %v", elemType.Name(), strings.Join(fieldDetails, ", "))
		} else {
			Printfln("%v: %v", elemType.Name(), elemValue)
		}
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
