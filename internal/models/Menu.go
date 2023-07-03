package models

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type Menu struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Products    []Product `json:"products"`
}

const (
	MenuTable          = "menus"
	ProductToMenuTable = "productToMenu"
)

func createMenuTables(db *sql.DB) {
	menu := `CREATE TABLE IF NOT EXISTS menus (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		description TEXT
	);`

	products := `CREATE TABLE IF NOT EXISTS productToMenu (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		menu_id INTEGER,
		product_id INTEGER,
		FOREIGN KEY(menu_id) REFERENCES menus(id) ON DELETE CASCADE ON UPDATE CASCADE,
		FOREIGN KEY(product_id) REFERENCES products(id) ON DELETE CASCADE ON UPDATE CASCADE
	);`

	_, err := db.Exec(menu)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(products)
	if err != nil {
		panic(err)
	}
}
