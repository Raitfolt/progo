package main

import (
	"reflect"
)

func getFieldValues(s interface{}) {
	structValue := reflect.ValueOf(s)
	if structValue.Kind() == reflect.Struct {
		for i := 0; i < structValue.NumField(); i++ {
			fieldType := structValue.Type().Field(i)
			fieldVal := structValue.Field(i)
			Printfln("Name: %v, Type: %v, Value: %v", fieldType.Name, fieldType.Type, fieldVal)
		}
	} else {
		Printfln("Not a struct")
	}
}

func main() {
	product := Product{Name: "Kayak", Category: "Watersports", Price: 279}
	customer := Customer{Name: "Acme", City: "Chicago"}
	purchase := Purchase{Customer: customer, Product: product, Total: 279, taxRate: 10}
	getFieldValues(purchase)
}
