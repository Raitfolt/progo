package main

import (
	"reflect"
)

func pickValues(slice interface{}, indices ...int) interface{} {
	sliceVal := reflect.ValueOf(slice)
	if sliceVal.Kind() == reflect.Slice {
		newSlice := reflect.MakeSlice(sliceVal.Type(), 0, 10)
		for _, index := range indices {
			newSlice = reflect.Append(newSlice, sliceVal.Index(index))
		}
		return newSlice
	}
	return nil
}

func main() {
	name := "Alice"
	city := "London"
	hobby := "Running"

	slice := []string{name, city, hobby, "Bob", "Paris", "Soccer"}
	picked := pickValues(slice, 0, 3, 5)
	Printfln("Picked values: %v", picked)
}
