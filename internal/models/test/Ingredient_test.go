package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/rainbowriverrr/CafeGo/internal/models"
	"github.com/rainbowriverrr/CafeGo/pkg/database"
)

type IngredientTestSuite struct {
	suite.Suite
	app                models.Application
	selectIngredients  string
	selectAlternatives string
}

func (suite *IngredientTestSuite) SetupTest() {
	suite.app = models.Application{
		DB: database.InitDB("testIngredient"),
	}
	models.CreateIngredientTables(suite.app.DB)
}

func (suite *IngredientTestSuite) TearDownTest() {
	database.DeleteDB("testIngredient")
}

func (suite *IngredientTestSuite) TestCreateIngredientTables() {
	//check if table exists
	exists, err := database.TableExists(suite.app.DB, models.IngredientTable)
	if err != nil {
		suite.T().Errorf("Error checking if table exists: %s", err.Error())
	}
	assert.True(suite.T(), exists)
}

func (suite *IngredientTestSuite) TestIngredientEquals() {
}

func (suite *IngredientTestSuite) TestIngredientExists() {

}

func (suite *IngredientTestSuite) TestGetIngredient() {

}

func (suite *IngredientTestSuite) TestUpdateIngredient() {

}

func (suite *IngredientTestSuite) TestInsertIngredient() {

}

func (suite *IngredientTestSuite) TestDeleteIngredient() {

}

func (suite *IngredientTestSuite) TestGetAlternatives() {

}

func (suite *IngredientTestSuite) TestInsertAlternative() {

}

func (suite *IngredientTestSuite) TestDeleteAlternative() {

}

func TestIngredient(t *testing.T) {
	suite.Run(t, new(IngredientTestSuite))
}
