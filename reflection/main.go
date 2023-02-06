package main

import (
	"reflect"
)

func inspectStructs(structs ...interface{}) {
	for _, s := range structs {
		structType := reflect.TypeOf(s)
		if structType.Kind() == reflect.Struct {
			inspectStructType(structType)
		}
	}
}

func inspectStructType(structType reflect.Type) {
	Printfln("--- Struct Type: %v", structType)
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		Printfln("Field: %v Name: %v, Type:%v, Exported: %v", field.Index, field.Name, field.Type, field.PkgPath == "")
	}
	Printfln("--- End Struct Type: %v", structType)
}

func main() {
	inspectStructs(Purchase{})
}
