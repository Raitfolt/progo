package main

import (
	"fmt"
	"math"
)

func main() {
	var intVal = math.MaxInt64
	var floatVal = math.MaxFloat64

	fmt.Println(intVal + 1)
	fmt.Println(floatVal)
	fmt.Println(floatVal + 1)
	fmt.Println(floatVal + 2)
	fmt.Println(floatVal + 3)
	fmt.Println(floatVal + 4)
	fmt.Println(math.IsInf((floatVal * 2), 0))
}
