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
	createAdminTable(db)
	createOrderTables(db)
}
