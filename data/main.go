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

func queryDatabase(db *sql.DB) (products []Product, err error) {
	rows, err := db.Query(`SELECT Products.Id, Products.Name, Products.Price,
		Categories.Id as "Category.Id", Categories.Name as "Category.Name" 
		FROM Products, Categories
		WHERE Products.Category = Categories.Id`)
	if err != nil {
		return
	} else {
		results, err := scanIntoStruct(rows, &Product{})
		if err == nil {
			products = (results).([]Product)
		} else {
			Printfln("Scanning error: %v", err)
		}
	}
	return
}

func insertRow(db *sql.DB, p *Product) (id int64) {
	res, err := db.Exec(`
	INSERT INTO Products (Name, Category, Price)
	VALUES (?, ?, ?)`, p.Name, p.Category.Id, p.Price)
	if err == nil {
		id, err = res.LastInsertId()
		if err != nil {
			Printfln("Result error: %v", err.Error())
		}
	} else {
		Printfln("Exec error: %v", err.Error())
	}
	return
}

func insertAndUseCategory(db *sql.DB, name string, productIDs ...int) (err error) {
	tx, err := db.Begin()
	updatedFailed := false
	if err == nil {
		catResult, err := tx.Stmt(insertNewCategory).Exec(name)
		if err == nil {
			newID, _ := catResult.LastInsertId()
			preparedStatement := tx.Stmt(changeProductCategory)
			for _, id := range productIDs {
				changeResult, err := preparedStatement.Exec(newID, id)
				if err == nil {
					changes, _ := changeResult.RowsAffected()
					if changes == 0 {
						updatedFailed = true
						break
					}
				}
			}
		}
	}
	if err != nil || updatedFailed {
		Printfln("Aborting transaction %v", err)
		tx.Rollback()
	} else {
		tx.Commit()
	}
	return
}

func main() {
	db, err := openDatabase()
	if err == nil {
		products, _ := queryDatabase(db)
		for _, p := range products {
			Printfln("Product: %v", p)
		}
		db.Close()
	} else {
		panic(err)
	}
}
