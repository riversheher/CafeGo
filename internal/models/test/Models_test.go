package models_test

import (
	"database/sql"
	"testing"

	"github.com/rainbowriverrr/CafeGo/internal/models"
	"github.com/rainbowriverrr/CafeGo/pkg/database"
	_ "modernc.org/sqlite"
)

// TestInitTables tests the InitTables function
func TestInitTables(t *testing.T) {
	var db *sql.DB = database.InitDB("test")
	defer db.Close()

	models.InitTables(db)

	// check if tables exist
	tables := []string{
		models.IngredientTable,
		models.MenuTable,
		models.ProductTable,
		models.UserTable,
		models.OrderTable,
		models.AlternativesTable,
		models.ProductToMenuTable,
		models.ProductToOrderTable,
		models.IngredientToProductTable,
	}

}
