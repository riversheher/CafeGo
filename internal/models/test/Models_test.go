package models_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/rainbowriverrr/CafeGo/internal/models"
	"github.com/rainbowriverrr/CafeGo/pkg/database"
	_ "modernc.org/sqlite"
)

const (
	testDB = "test"
)

var (
	app    models.Application
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
	//setup
	app = models.Application{
		DB: database.InitDB(testDB),
	}
	defer app.DB.Close()

	//tests
	exitVal := m.Run()

	//clean up
	clean()

	os.Exit(exitVal)
}

func clean() {
	//delete test database
	app.DB.Close()
	database.DeleteDB(testDB)
}

// TestInitTables tests the InitTables function
func TestInitTables(t *testing.T) {
	models.InitTables(app.DB)

	query := "SELECT name FROM sqlite_schema WHERE type='table' AND name NOT  LIKE 'sqlite_%';"
	rows, err := app.DB.Query(query)
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
		_, err := app.DB.Query(fmt.Sprintf("SELECT * FROM %s", table))
		if err != nil {
			t.Errorf("Error querying table %s: %s", table, err.Error())
		}
	}

	//drop all tables
	for _, table := range tables {
		_, err := app.DB.Exec(fmt.Sprintf("DROP TABLE %s", table))
		if err != nil {
			t.Errorf("Error dropping table %s: %s", table, err.Error())
		}
	}
}
