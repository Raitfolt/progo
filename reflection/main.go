package main

import (
	"reflect"
)

func inspectFuncType(f interface{}) {
	funcType := reflect.TypeOf(f)
	if funcType.Kind() == reflect.Func {
		Printfln("Function parameters: %v", funcType.NumIn())
		for i := 0; i < funcType.NumIn(); i++ {
			paramType := funcType.In(i)
			if i < funcType.NumIn()-1 {
				Printfln("Parameter #%v, Type: %v", i, paramType)
			} else {
				Printfln("Parameter #%v, Type: %v, Variadic: %v", i, paramType, funcType.IsVariadic())
			}
		}
		Printfln("Function results: %v", funcType.NumOut())
		for i := 0; i < funcType.NumOut(); i++ {
			resultType := funcType.Out(i)
			Printfln("Result #%v, Type: %v", i, resultType)
		}
	}
}

func main() {
	inspectFuncType(Find)
}
