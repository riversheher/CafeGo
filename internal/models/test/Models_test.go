package models_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/rainbowriverrr/CafeGo/internal/models"
	"github.com/rainbowriverrr/CafeGo/pkg/database"
	_ "modernc.org/sqlite"
)

const (
	testDB = "test"
)

var (
	db     *sql.DB
	tables = []string{
		models.IngredientTable,
		models.MenuTable,
		models.ProductTable,
		models.UserTable,
		models.AdminTable,
		models.OrderTable,
		models.AlternativesTable,
		models.ProductToMenuTable,
		models.ProductToOrderTable,
		models.IngredientToProductTable,
	}
)

func TestMain(m *testing.M) {
	db = database.InitDB(testDB)
	defer db.Close()

	m.Run()

	clean()
}

func clean() {
	//delete test database
	db.Close()
	database.DeleteDB(testDB)
}

// TestInitTables tests the InitTables function
func TestInitTables(t *testing.T) {
	models.InitTables(db)

	query := "SELECT name FROM sqlite_schema WHERE type='table' AND name NOT  LIKE 'sqlite_%';"
	rows, err := db.Query(query)
	if err != nil {
		t.Errorf("Error querying database: %s", err.Error())
	}

	defer rows.Close()

	//Print out all tables
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			t.Errorf("Error scanning rows: %s", err.Error())
		}
		t.Logf("Table: %s", name)
	}

	// check if tables exist
	for _, table := range tables {
		_, err := db.Query(fmt.Sprintf("SELECT * FROM %s", table))
		if err != nil {
			t.Errorf("Error querying table %s: %s", table, err.Error())
		}
	}

	//drop all tables
	for _, table := range tables {
		_, err := db.Exec(fmt.Sprintf("DROP TABLE %s", table))
		if err != nil {
			t.Errorf("Error dropping table %s: %s", table, err.Error())
		}
	}
}
