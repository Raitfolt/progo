package main

import "database/sql"

type Category struct {
	Id   int
	Name string
}

type Product struct {
	Id   int
	Name string
	Category
	Price float64
}

func queryDatabase(db *sql.DB, categoryName string) []Product {
	products := []Product{}
	rows, err := db.Query(`SELECT Products.Id, Products.Name, Products.Price,
		Categories.Id as Cat_Id, Categories.Name as CatName from Products, Categories
		WHERE Products.Category = Categories.Id 
		AND Categories.Name = ?`, categoryName)
	if err == nil {
		for rows.Next() {
			p := Product{}
			scanErr := rows.Scan(&p.Id, &p.Name, &p.Price, &p.Category.Id, &p.Category.Name)
			if scanErr == nil {
				products = append(products, p)
			} else {
				Printfln("Scan error: %v", scanErr)
				break
			}
		}
	} else {
		Printfln("Error: %v", err)
	}
	return products
}

func main() {
	db, err := openDatabase()
	if err == nil {
		for _, cat := range []string{"Soccer", "Watersports"} {
			Printfln("--- %v Results ---", cat)
			products := queryDatabase(db, cat)
			for i, p := range products {
				Printfln("#%v: %v %v %v", i, p.Name, p.Category.Name, p.Price)
			}
		}
		db.Close()
	} else {
		panic(err)
	}
}
