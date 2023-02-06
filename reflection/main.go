package main

import (
	"reflect"
)

func checkElemType(val interface{}, arrOrSlice interface{}) bool {
	elemType := reflect.TypeOf(val)
	arrOrSliceType := reflect.TypeOf(arrOrSlice)
	return (arrOrSliceType.Kind() == reflect.Array ||
		arrOrSliceType.Kind() == reflect.Slice) &&
		arrOrSliceType.Elem() == elemType
}

func main() {
	name := "Alice"
	city := "London"
	hobby := "Running"

	slice := []string{name, city, hobby}
	array := [3]string{name, city, hobby}

	Printfln("Slice (string): %v", checkElemType("testString", slice))
	Printfln("Array (string): %v", checkElemType("testString", array))
	Printfln("Array (int): %v", checkElemType(10, array))
}
