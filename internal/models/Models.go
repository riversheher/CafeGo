package models

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func InitTables(db *sql.DB) {
	createIngredientTables(db)
	createProductTables(db)
	createMenuTables(db)
	createUserTables(db)
	createAdminTables(db)
	createOrderTables(db)
}
