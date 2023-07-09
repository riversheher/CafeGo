package models

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func InitTables(db *sql.DB) {
	CreateIngredientTables(db)
	CreateProductTables(db)
	CreateMenuTables(db)
	CreateUserTables(db)
	CreateAdminTables(db)
	CreateOrderTables(db)
}
