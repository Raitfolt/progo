package main

import (
	"reflect"
)

func enumerateStrings(arrayOrSlice interface{}) {
	arrayOrSliceVal := reflect.ValueOf(arrayOrSlice)
	if (arrayOrSliceVal.Kind() == reflect.Array ||
		arrayOrSliceVal.Kind() == reflect.Slice) &&
		arrayOrSliceVal.Type().Elem().Kind() == reflect.String {
		for i := 0; i < arrayOrSliceVal.Len(); i++ {
			Printfln("Element: %v, Value: %v", i, arrayOrSliceVal.Index(i).String())
		}
	}
}

func main() {
	name := "Alice"
	city := "London"
	hobby := "Running"

	slice := []string{name, city, hobby}
	array := [3]string{name, city, hobby}

	enumerateStrings(slice)
	enumerateStrings(array)
}
