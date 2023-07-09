package models

import (
	"database/sql"
	"fmt"
	"time"

	_ "modernc.org/sqlite"
)

type Order struct {
	ID       int64     `json:"id"`
	Customer User      `json:"customer"`
	Products []Product `json:"products"`
	Time     time.Time `json:"time"`
	Subtotal float64   `json:"subtotal"`
}

const (
	OrderTable          = "orders"
	ProductToOrderTable = "productToOrder"
)

func createOrderTables(db *sql.DB) {
	order := `CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		customer_id INTEGER,
		time TIMESTAMP,
		subtotal REAL,
		FOREIGN KEY(customer_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
	);`

	products := `CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		order_id INTEGER,
		product_id INTEGER,
		FOREIGN KEY(order_id) REFERENCES orders(id) ON DELETE CASCADE ON UPDATE CASCADE,
		FOREIGN KEY(product_id) REFERENCES products(id) ON DELETE CASCADE ON UPDATE CASCADE
	);`

	_, err := db.Exec(fmt.Sprintf(order, OrderTable))
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(fmt.Sprintf(products, ProductToOrderTable))
	if err != nil {
		panic(err)
	}
}
