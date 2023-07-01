package models

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type Product struct {
	ID          int64        `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Price       float64      `json:"price"`
	Ingredients []Ingredient `json:"ingredients"`
	Discount    float64      `json:"discount"`
	Type        string       `json:"type"`
}

const (
	ProductTable        = "products"
	IngredientToProduct = "ingredientToProduct"
)

func CreateProductTables(db *sql.DB) {
	products := `CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		description TEXT,
		price REAL,
		discount REAL,
		type TEXT
	);`

	ingredients := `CREATE TABLE IF NOT EXISTS ingredientToProduct (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		product_id INTEGER,
		ingredient_id INTEGER,
		FOREIGN KEY(product_id) REFERENCES products(id) ON DELETE CASCADE ON UPDATE CASCADE,
		FOREIGN KEY(ingredient_id) REFERENCES ingredients(id) ON DELETE CASCADE ON UPDATE CASCADE
	);`

	_, err := db.Exec(products)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(ingredients)
	if err != nil {
		panic(err)
	}
}
