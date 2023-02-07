package main

import (
	"fmt"
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

func describeField(s interface{}, fieldName string) {
	structType := reflect.TypeOf(s)
	field, found := structType.FieldByName(fieldName)
	if found {
		Printfln("Found: %v, Type: %v, Index: %v", field.Name, field.Type, field.Index)
		index := field.Index
		for len(index) > 1 {
			index = index[0 : len(index)-1]
			field = structType.FieldByIndex(index)
			Printfln("Parent: %v, Type: %v, Index: %v", field.Name, field.Type, field.Index)
		}
		Printfln("Top-Level Type: %v", structType)
	} else {
		Printfln("Field %v not found", fieldName)
	}
}

func main() {
	inspectStructs(Purchase{})
	fmt.Println("-----------------------------------------------------")
	describeField(Purchase{}, "Price")
}
