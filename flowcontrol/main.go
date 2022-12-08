package main

import (
	"fmt"
	"strconv"
)

func main() {

	priceString := "275"

	if kayakPrice, err := strconv.Atoi(priceString); err == nil {
		fmt.Println("Price:", kayakPrice)
	} else {
		fmt.Println("Error:", err)
	}

}
