package models

import (
	"database/sql"
	"fmt"
	"time"

	_ "modernc.org/sqlite"
)

type Order struct {
	ID         int64     `json:"id"`
	Customer   User      `json:"customer"`
	Products   []Product `json:"products"`
	Time       time.Time `json:"time"`
	Subtotal   float64   `json:"subtotal"`
	InProgress bool      `json:"inProgress"`
}

const (
	OrderTable          = "orders"
	ProductToOrderTable = "productToOrder"
)

func (o Order) Equals(other Order) bool {
	return false
}

func CreateOrderTables(db *sql.DB) {
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

func (app *Application) OrderExists(order Order) bool {
	return false
}

func (app *Application) GetOrder(id int64) (Order, error) {
	return Order{}, nil
}

func (app *Application) GetAllOrders() ([]Order, error) {
	return []Order{}, nil
}

func (app *Application) GetInProgressOrders() ([]Order, error) {
	return []Order{}, nil
}

func (app *Application) GetCompletedOrders() ([]Order, error) {
	return []Order{}, nil
}

func (app *Application) GetOrdersByUser(user User) ([]Order, error) {
	return []Order{}, nil
}

func (app *Application) UpdateOrder(order Order) error {
	return nil
}

func (app *Application) InsertOrder(order Order) (Order, error) {
	return Order{}, nil
}

func (app *Application) AddProductToOrder(order Order, product Product) error {
	return nil
}

func (app *Application) RemoveProductFromOrder(order Order, product Product) error {
	return nil
}

func (app *Application) GetProductsFromOrder(order Order) ([]Product, error) {
	return []Product{}, nil
}
