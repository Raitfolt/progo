package main

import (
	"encoding/json"
	"os"
)

func main() {
	cheapProducts := []Product{}
	for _, p := range Products {
		if p.Price < 100 {
			cheapProducts = append(cheapProducts, p)
		}
	}
	//file, err := os.OpenFile("cheap.json", os.O_WRONLY|os.O_CREATE, 0666)
	file, err := os.CreateTemp(".", "tempfile-*.json")
	if err == nil {
		defer file.Close()
		encoder := json.NewEncoder(file)
		encoder.Encode(cheapProducts)
	} else {
		Printfln("Error: %v", err.Error())
	}
}
