package main

import (
	"reflect"
	"strings"
)

func incrementOrUpper(values ...interface{}) {
	for _, elem := range values {
		elemValue := reflect.ValueOf(elem)
		if elemValue.Kind() == reflect.Ptr {
			elemValue = elemValue.Elem()
		}
		if elemValue.CanSet() {
			switch elemValue.Kind() {
			case reflect.Int:
				elemValue.SetInt(elemValue.Int() + 1)
			case reflect.String:
				elemValue.SetString(strings.ToUpper(elemValue.String()))
			}
			Printfln("Modified Value: %v", elemValue)
		} else {
			Printfln("Cannot set %v: %v", elemValue.Kind(), elemValue)
		}
	}
}

func selectValue(data interface{}, index int) (result interface{}) {
	dataVal := reflect.ValueOf(data)
	if dataVal.Kind() == reflect.Slice {
		result = dataVal.Index(index).Interface()
	}
	return
}

func getTypePath(t reflect.Type) (path string) {
	path = t.PkgPath()
	if path == "" {
		path = "(built-in)"
	}
	return
}

var intPtrType = reflect.TypeOf((*int)(nil))
var byteSliceType = reflect.TypeOf([]byte(nil))

func printDetails(values ...interface{}) {
	for _, elem := range values {
		elemType := reflect.TypeOf(elem)
		elemValue := reflect.ValueOf(elem)
		if elemType == intPtrType {
			Printfln("Pointer to Int: %v", elemValue.Elem().Int())
		} else if elemType == byteSliceType {
			Printfln("Byte slice: %v", elemValue.Bytes())
		} else {
			switch elemValue.Kind() {
			case reflect.Bool:
				var val bool = elemValue.Bool()
				Printfln("Bool: %v", val)
			case reflect.Int:
				var val int64 = elemValue.Int()
				Printfln("Int: %v ", val)
			case reflect.Float32, reflect.Float64:
				var val float64 = elemValue.Float()
				Printfln("Float: %v", val)
			case reflect.String:
				var val string = elemValue.String()
				Printfln("String: %v", val)
			case reflect.Ptr:
				var val reflect.Value = elemValue.Elem()
				if val.Kind() == reflect.Int {
					Printfln("Pointer to Int: %v", val.Int())
				}
			default:
				Printfln("Other: %v", elemValue.String())
			}
		}
	}
}

func setAll(src interface{}, targets ...interface{}) {
	srcVal := reflect.ValueOf(src)
	for _, target := range targets {
		targetVal := reflect.ValueOf(target)
		if targetVal.Kind() == reflect.Ptr &&
			targetVal.Elem().Type() == srcVal.Type() &&
			targetVal.Elem().CanSet() {
			targetVal.Elem().Set(srcVal)
		}
	}
}

func containsOld(slice interface{}, target interface{}) (found bool) {
	sliceVal := reflect.ValueOf(slice)
	targetType := reflect.TypeOf(target)
	if sliceVal.Kind() == reflect.Slice &&
		sliceVal.Type().Elem().Comparable() &&
		targetType.Comparable() {
		for i := 0; i < sliceVal.Len(); i++ {
			if sliceVal.Index(i).Interface() == target {
				found = true
			}
		}
	}
	return
}

func contains(slice interface{}, target interface{}) (found bool) {
	sliceVal := reflect.ValueOf(slice)
	if sliceVal.Kind() == reflect.Slice {
		for i := 0; i < sliceVal.Len(); i++ {
			if reflect.DeepEqual(sliceVal.Index(i).Interface(), target) {
				found = true
			}
		}
	}
	return
}

func main() {
	product := Product{
		Name: "Kayak", Category: "Watersports", Price: 279,
	}
	number := 100
	slice := []byte("Alice")
	printDetails(true, 10, 23.30, "Alice", &number, product, slice)
	Printfln("----------------------------------------------")
	names := []string{"Alice", "Bob", "Charlie"}
	val := selectValue(names, 1).(string)
	Printfln("Selected: %v", val)
	Printfln("----------------------------------------------")
	name := "Alice"
	price := 279
	city := "London"

	incrementOrUpper(&name, &price, &city)
	for _, val := range []interface{}{name, price, city} {
		Printfln("Value: %v", val)
	}
	Printfln("----------------------------------------------")
	setAll("New String", &name, &price, &city)
	setAll(10, &name, &price, &city)
	for _, val := range []interface{}{name, price, city} {
		Printfln("Value: %v", val)
	}
	Printfln("----------------------------------------------")
	city = "London"
	citiesSlice := []string{"Paris", "Rome", "London"}
	Printfln("Found #1: %v", contains(citiesSlice, city))

	sliceOfSlices := [][]string{
		citiesSlice, {"First", "Second", "Thrid"}}
	Printfln("Found #2: %v", contains(sliceOfSlices, citiesSlice))
}
