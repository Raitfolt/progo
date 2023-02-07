package main

import (
	"fmt"
	"reflect"
	"strings"
)

func makeMapperFunc(mapper interface{}) interface{} {
	mapVal := reflect.ValueOf(mapper)
	if mapVal.Kind() == reflect.Func && mapVal.Type().NumIn() == 1 &&
		mapVal.Type().NumOut() == 1 {
		inType := reflect.SliceOf(mapVal.Type().In(0))
		inTypeSlice := []reflect.Type{inType}
		outType := reflect.SliceOf(mapVal.Type().Out(0))
		outTypeSlice := []reflect.Type{outType}
		funcType := reflect.FuncOf(inTypeSlice, outTypeSlice,
			false)
		funcVal := reflect.MakeFunc(funcType,
			func(params []reflect.Value) (results []reflect.Value) {
				srcSliceVal := params[0]
				resultsSliceVal := reflect.MakeSlice(outType, srcSliceVal.Len(), 10)
				for i := 0; i < srcSliceVal.Len(); i++ {
					r := mapVal.Call([]reflect.Value{
						srcSliceVal.Index(i)})
					resultsSliceVal.Index(i).Set(r[0])
				}
				results = []reflect.Value{resultsSliceVal}
				return
			})
		return funcVal.Interface()
	}
	Printfln("Unexpected types")
	return nil
}
func main() {
	lowerStringMapper := makeMapperFunc(strings.ToLower).(func([]string) []string)
	names := []string{"Alice", "Bob", "Charlie"}
	results := lowerStringMapper(names)
	Printfln("Lowercase Results: %v", results)

	incrementFloatMapper := makeMapperFunc(func(val float64) float64 {
		return val + 1
	}).(func([]float64) []float64)
	prices := []float64{279, 48.95, 19.50}
	floatResults := incrementFloatMapper(prices)
	Printfln("Increment Results: %v", floatResults)
	floatToStringMapper := makeMapperFunc(func(val float64) string {
		return fmt.Sprintf("$%.2f", val)
	}).(func([]float64) []string)
	Printfln("Price Results: %v",
		floatToStringMapper(prices))
}
