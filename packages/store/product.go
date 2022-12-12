// Package store provides types and methods
// commonly required for online sales
package store

// Product describes an item for sale
type Product struct {
	Name, Category string // Name and type of the product
	price          float64
}
