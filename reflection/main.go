package main

import (
	"reflect"
	"strings"
)

func mapSlice(slice interface{}, mapper interface{}) (mapped []interface{}) {
	sliceVal := reflect.ValueOf(slice)
	mapperVal := reflect.ValueOf(mapper)
	mapped = []interface{}{}
	if sliceVal.Kind() == reflect.Slice && mapperVal.Kind() == reflect.Func &&
		mapperVal.Type().NumIn() == 1 &&
		mapperVal.Type().In(0) == sliceVal.Type().Elem() {
		for i := 0; i < sliceVal.Len(); i++ {
			result := mapperVal.Call([]reflect.Value{sliceVal.Index(i)})
			for _, r := range result {
				mapped = append(mapped, r.Interface())
			}
		}
	}
	return
}

func main() {
	names := []string{"Alice", "Bob", "Charlie"}
	results := mapSlice(names, strings.ToUpper)
	Printfln("Results: %v", results)
}
