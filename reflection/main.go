package main

import (
	"fmt"
	"reflect"
)

func setMap(m interface{}, key interface{}, val interface{}) {
	mapValue := reflect.ValueOf(m)
	keyValue := reflect.ValueOf(key)
	valValue := reflect.ValueOf(val)
	if mapValue.Kind() == reflect.Map &&
		mapValue.Type().Key() == keyValue.Type() &&
		mapValue.Type().Elem() == valValue.Type() {
		mapValue.SetMapIndex(keyValue, valValue)
	} else {
		Printfln("Not a map or mosmatched types")
	}
}

func removeFromMap(m interface{}, key interface{}) {
	mapValue := reflect.ValueOf(m)
	keyValue := reflect.ValueOf(key)
	if mapValue.Kind() == reflect.Map &&
		mapValue.Type().Key() == keyValue.Type() {
		mapValue.SetMapIndex(keyValue, reflect.Value{})
	}
}

func main() {
	pricesMap := map[string]float64{
		"Kayak": 279, "Lifejacket": 48.95, "Soccer Ball": 19.50,
	}
	for k, v := range pricesMap {
		Printfln("Key: %v, Value: %v", k, v)
	}
	fmt.Println()
	setMap(pricesMap, "Kayak", 100.00)
	for k, v := range pricesMap {
		Printfln("Key: %v, Value: %v", k, v)
	}
	fmt.Println()
	setMap(pricesMap, "Hat", 10.00)
	for k, v := range pricesMap {
		Printfln("Key: %v, Value: %v", k, v)
	}
	fmt.Println()
	removeFromMap(pricesMap, "Lifejacket")
	for k, v := range pricesMap {
		Printfln("Key: %v, Value: %v", k, v)
	}
}
