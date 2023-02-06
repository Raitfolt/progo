package main

import (
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

func main() {
	pricesMap := map[string]float64{
		"Kayak": 279, "Lifejacket": 48.95, "Soccer Ball": 19.50,
	}
	describeMap(pricesMap)
}
