package main

import (
	"reflect"
)

func inspectMethods(s interface{}) {
	sType := reflect.TypeOf(s)
	if sType.Kind() == reflect.Struct ||
		(sType.Kind() == reflect.Ptr && sType.Elem().Kind() == reflect.Struct) {
		Printfln("Type: %v, Methods: %v", sType, sType.NumMethod())
		for i := 0; i < sType.NumMethod(); i++ {
			method := sType.Method(i)
			Printfln("Method name: %v, Type: %v", method.Name, method.Type)
		}
	}

}

func main() {
	inspectMethods(Purchase{})
	inspectMethods(&Purchase{})
}
