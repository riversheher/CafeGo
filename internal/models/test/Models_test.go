package models_test

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/rainbowriverrr/CafeGo/internal/models"
	"github.com/rainbowriverrr/CafeGo/pkg/database"
	_ "modernc.org/sqlite"
)

const (
	mainDB = "test"
)

type ModelsTestSuite struct {
	suite.Suite
	MainDB *sql.DB
	tables []string
}

func (suite *ModelsTestSuite) SetupTest() {
	suite.MainDB = database.InitDB(mainDB)
	suite.tables = []string{
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
}

func (suite *ModelsTestSuite) TearDownTest() {
	database.DeleteDB(mainDB)
}

// TestInitTables tests the InitTables function
func (suite *ModelsTestSuite) TestInitTables() {
	models.InitTables(suite.MainDB)

	for _, table := range suite.tables {
		exists, err := database.TableExists(suite.MainDB, table)
		if err != nil {
			suite.T().Errorf("Error checking if table exists: %s", err.Error())
		}
		assert.True(suite.T(), exists)
	}
}

func TestModels(t *testing.T) {
	suite.Run(t, new(ModelsTestSuite))
}
