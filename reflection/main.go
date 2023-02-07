package main

import (
	"reflect"
)

func inspectStructs(structs ...interface{}) {
	for _, s := range structs {
		structType := reflect.TypeOf(s)
		if structType.Kind() == reflect.Struct {
			inspectStructType([]int{}, structType)
		}
	}
}

func inspectStructType(baseIndex []int, structType reflect.Type) {
	Printfln("--- Struct Type: %v", structType)
	for i := 0; i < structType.NumField(); i++ {
		fieldIndex := append(baseIndex, i)
		field := structType.Field(i)
		Printfln("Field: %v Name: %v, Type:%v, Exported: %v", field.Index, field.Name, field.Type, field.PkgPath == "")
		if field.Type.Kind() == reflect.Struct {
			field := structType.FieldByIndex(field.Index)
			inspectStructType(fieldIndex, field.Type)
		}
	}
	Printfln("--- End Struct Type: %v", structType)
}

func main() {
	inspectStructs(Purchase{})
}
