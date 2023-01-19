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

func queryDatabase(db *sql.DB, id int) (p Product) {
	row := db.QueryRow(`SELECT Products.Id, Products.Name, Products.Price,
		Categories.Id as Cat_Id, Categories.Name as CatName 
		FROM Products, Categories
		WHERE Products.Category = Categories.Id 
		AND Products.Id = ?`, id)
	if row.Err() == nil {
		scanErr := row.Scan(&p.Id, &p.Name, &p.Price, &p.Category.Id, &p.Category.Name)
		if scanErr != nil {
			Printfln("Scan error: %v", scanErr)
		}
	} else {
		Printfln("Row error: %v", row.Err().Error())
	}
	return
}

func main() {
	db, err := openDatabase()
	if err == nil {
		for _, id := range []int{1, 3, 10} {
			p := queryDatabase(db, id)
			Printfln("Product: %v", p)
		}
		db.Close()
	} else {
		panic(err)
	}
}
