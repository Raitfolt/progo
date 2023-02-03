package main

import (
	"reflect"
)

func swap(first interface{}, second interface{}) {
	firstValue, secondValue := reflect.ValueOf(first), reflect.ValueOf(second)
	if firstValue.Type() == secondValue.Type() &&
		firstValue.Kind() == reflect.Ptr &&
		firstValue.Elem().CanSet() &&
		secondValue.Elem().CanSet() {
		tmp := reflect.New(firstValue.Elem().Type())
		tmp.Elem().Set(firstValue.Elem())
		firstValue.Elem().Set(secondValue.Elem())
		secondValue.Elem().Set(tmp.Elem())
	}
}

func main() {
	name := "Alice"
	price := 279
	city := "London"

	for _, val := range []interface{}{name, price, city} {
		Printfln("Old value: %v", val)
	}
	swap(&name, &city)
	for _, val := range []interface{}{name, price, city} {
		Printfln("Value: %v", val)
	}

}
