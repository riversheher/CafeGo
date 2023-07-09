package models

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

type Application struct {
	DB     *sql.DB
	ErrLog log.Logger
}

func InitTables(db *sql.DB) {
	CreateIngredientTables(db)
	CreateProductTables(db)
	CreateMenuTables(db)
	CreateUserTables(db)
	CreateAdminTables(db)
	CreateOrderTables(db)
}
