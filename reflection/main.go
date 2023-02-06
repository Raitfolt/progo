package main

import (
	"reflect"
	"strings"
)

var stringPtrType = reflect.TypeOf((*string)(nil))

func transformString(val interface{}) {
	elemValue := reflect.ValueOf(val)
	if elemValue.Type() == stringPtrType {
		upperStr := strings.ToUpper(elemValue.Elem().String())
		if elemValue.Elem().CanSet() {
			elemValue.Elem().SetString(upperStr)
		}
	}
}

func main() {
	name := "Alice"

	transformString(&name)
	Printfln("Follow pointer value: %v", name)
}
