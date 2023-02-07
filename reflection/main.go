package main

import (
	"reflect"
)

func inspectTags(s interface{}, tagName string) {
	structType := reflect.TypeOf(s)
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		tag := field.Tag
		valGet := tag.Get(tagName)
		valLookup, ok := tag.Lookup(tagName)
		Printfln("Field: %v, Tag %v: %v", field.Name, tagName, valGet)
		Printfln("Field: %v, Tag %v: %v, Set: %v", field.Name, tagName, valLookup, ok)
	}
}

func main() {
	stringType := reflect.TypeOf("string")

	structType := reflect.StructOf([]reflect.StructField{
		{Name: "Name", Type: stringType, Tag: `alias:"id"`},
		{Name: "City", Type: stringType, Tag: `alias:""`},
		{Name: "Country", Type: stringType},
	})

	inspectTags(reflect.New(structType), "alias")
}
