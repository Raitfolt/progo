package main

import (
	"fmt"
)

func Sum(array []int) int {
	result := 0
	for _, value := range array {
		result += value
	}
	return result
}

func main() {
	for index, char := range []rune("€48.95") {
		fmt.Println(index, char, string(char))
	}
	fmt.Println()
	for index, char := range "€48.95" {
		fmt.Println(index, char, string(char))
	}
	fmt.Println()
	for index, char := range []byte("€48.95") {
		fmt.Println(index, char)
	}
}
