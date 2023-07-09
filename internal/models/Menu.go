package models

import (
	"database/sql"
	"fmt"

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
	menu := `CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		description TEXT
	);`

	products := `CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		menu_id INTEGER,
		product_id INTEGER,
		FOREIGN KEY(menu_id) REFERENCES menus(id) ON DELETE CASCADE ON UPDATE CASCADE,
		FOREIGN KEY(product_id) REFERENCES products(id) ON DELETE CASCADE ON UPDATE CASCADE
	);`

	_, err := db.Exec(fmt.Sprintf(menu, MenuTable))
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(fmt.Sprintf(products, ProductToMenuTable))
	if err != nil {
		panic(err)
	}
}
