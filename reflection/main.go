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

func invokeFunction(f interface{}, params ...interface{}) {
	paramVals := []reflect.Value{}
	for _, p := range params {
		paramVals = append(paramVals, reflect.ValueOf(p))
	}
	funcVal := reflect.ValueOf(f)
	if funcVal.Kind() == reflect.Func {
		results := funcVal.Call(paramVals)
		for i, r := range results {
			Printfln("Result #%v: %v", i, r)
		}
	}
}

func main() {
	names := []string{"Alice", "Bob", "Charlie"}
	invokeFunction(Find, names, "London", "Bob")
}
