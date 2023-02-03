package main

import (
	"reflect"
)

func convert(src, target interface{}) (result interface{}, assigned bool) {
	srcVal := reflect.ValueOf(src)
	targetVal := reflect.ValueOf(target)
	if srcVal.Type().ConvertibleTo(targetVal.Type()) {
		result = srcVal.Convert(targetVal.Type()).Interface()
		assigned = true
	} else {
		result = src
	}
	return
}

func main() {
	name := "Alice"
	price := 279

	newVal, ok := convert(price, 100.00)
	Printfln("Converted %v: %v, %T", ok, newVal, newVal)
	newVal, ok = convert(name, 100.00)
	Printfln("Converted %v: %v, %T", ok, newVal, newVal)
}
