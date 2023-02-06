package main

import (
	"fmt"
	"reflect"
)

func describeMap(m interface{}) {
	mapType := reflect.TypeOf(m)
	if mapType.Kind() == reflect.Map {
		Printfln("Key type: %v, Value type: %v", mapType.Key(), mapType.Elem())
	} else {
		Printfln("Not a map")
	}
}

func printMapContents(m interface{}) {
	mapValue := reflect.ValueOf(m)
	if mapValue.Kind() == reflect.Map {
		for _, keyVal := range mapValue.MapKeys() {
			reflectedVal := mapValue.MapIndex(keyVal)
			Printfln("Map Key: %v, Value: %v", keyVal, reflectedVal)
		}
	} else {
		Printfln("Not a map")
	}
}

func printMapContents2(m interface{}) {
	mapValue := reflect.ValueOf(m)
	if mapValue.Kind() == reflect.Map {
		iter := mapValue.MapRange()
		for iter.Next() {
			Printfln("Map Key: %v, Value: %v", iter.Key(), iter.Value())
		}
	} else {
		Printfln("Not a map")
	}
}

func main() {
	pricesMap := map[string]float64{
		"Kayak": 279, "Lifejacket": 48.95, "Soccer Ball": 19.50,
	}
	describeMap(pricesMap)
	fmt.Println()
	printMapContents(pricesMap)
	fmt.Println()
	printMapContents2(pricesMap)
}
