package main

import (
	"reflect"
)

func findAndSplit(slice interface{}, target interface{}) interface{} {
	sliceVal := reflect.ValueOf(slice)
	targetType := reflect.TypeOf(target)
	if sliceVal.Kind() == reflect.Slice && sliceVal.Type().Elem() == targetType {
		for i := 0; i < sliceVal.Len(); i++ {
			if sliceVal.Index(i).Interface() == target {
				return sliceVal.Slice(0, i+1)
			}
		}
	}
	return slice
}

func main() {
	name := "Alice"
	city := "London"
	hobby := "Running"

	slice := []string{name, city, hobby}
	Printfln("Strings: %v", findAndSplit(slice, "London"))

	numbers := []int{1, 3, 5, 7, 9}
	Printfln("Numbers: %v", findAndSplit(numbers, 4))
}
