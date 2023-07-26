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

func (m Menu) Equals(other Menu) bool {
	return false
}

func CreateMenuTables(db *sql.DB) {
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

func (app *Application) MenuExists(menu Menu) bool {
	return false
}

func (app *Application) GetMenu(id int64) (Menu, error) {
	return Menu{}, nil
}

func (app *Application) GetMenus() ([]Menu, error) {
	return []Menu{}, nil
}

func (app *Application) UpdateMenu(menu Menu) error {
	return nil
}

func (app *Application) DeleteMenu(menu Menu) error {
	return nil
}

func (app *Application) InsertMenu(menu Menu) (Menu, error) {
	return Menu{}, nil
}

func (app *Application) AddProductToMenu(menu Menu, product Product) error {
	return nil
}

func (app *Application) RemoveProductFromMenu(menu Menu, product Product) error {
	return nil
}

func (app *Application) GetProductsFromMenu(menu Menu) ([]Product, error) {
	return []Product{}, nil
}
