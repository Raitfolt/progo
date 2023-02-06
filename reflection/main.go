package main

import (
	"reflect"
)

func setValue(arrayOrSlice interface{}, index int, replacment interface{}) {
	arrayOrSliceVal := reflect.ValueOf(arrayOrSlice)
	replacementVal := reflect.ValueOf(replacment)
	if arrayOrSliceVal.Kind() == reflect.Slice {
		elemVal := arrayOrSliceVal.Index(index)
		if elemVal.CanSet() {
			elemVal.Set(replacementVal)
		}
	} else if arrayOrSliceVal.Kind() == reflect.Ptr &&
		arrayOrSliceVal.Elem().Kind() == reflect.Array &&
		arrayOrSliceVal.Elem().CanSet() {
		arrayOrSliceVal.Elem().Index(index).Set(replacementVal)
	}
}

func main() {
	name := "Alice"
	city := "London"
	hobby := "Running"

	slice := []string{name, city, hobby}
	array := [3]string{name, city, hobby}

	Printfln("Original slice: %v", slice)
	newCity := "Paris"
	setValue(slice, 1, newCity)
	Printfln("Modyfied slice: %v", slice)

	Printfln("Original slice: %v", array)
	newCity = "Rome"
	setValue(&array, 1, newCity)
	Printfln("Modyfied slice: %v", array)
}
