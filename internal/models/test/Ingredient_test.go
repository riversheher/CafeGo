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
	insertIngredients  string
	insertAlternatives string
	deleteIngredients  string
	deleteAlternatives string
	chives             models.Ingredient
	leek               models.Ingredient
}

func (suite *IngredientTestSuite) SetupTest() {
	suite.app = models.Application{
		DB: database.InitDB("testIngredient"),
	}
	models.CreateIngredientTables(suite.app.DB)
	suite.chives = models.Ingredient{
		ID:          int64(1),
		Name:        "Chives",
		Description: "Floral green herbs",
		Price:       float64(0.5),
		Amount:      float64(500),
		Type:        models.Weight,
	}
	suite.leek = models.Ingredient{
		ID:           int64(2),
		Name:         "Leek",
		Description:  "Floral green herb with more cronch",
		Price:        float64(0.4),
		Amount:       float64(800),
		Alternatives: []models.Ingredient{suite.chives},
		Type:         models.Weight,
	}
	suite.chives.Alternatives = []models.Ingredient{suite.leek}

	suite.selectIngredients = "SELECT id, name, description, price, amount, type FROM ingredients WHERE id = ?"
	suite.selectAlternatives = "SELECT id, ingredient_id, alternative_id FROM alternatives WHERE id = ?"
	suite.insertIngredients = "INSERT INTO ingredients (name, description, price, amount, type) VALUES (?, ?, ?, ?, ?)"
	suite.insertAlternatives = "INSERT INTO alternatives (ingredient_id, alternative_id) VALUES (?, ?)"
	suite.deleteIngredients = "DELETE FROM ingredients WHERE id = ?"
	suite.deleteAlternatives = "DELETE FROM alternatives WHERE id = ?"
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

	leek1 := models.Ingredient{
		ID:           int64(2),
		Name:         "Leek",
		Description:  "Floral green herb with more cronch",
		Price:        float64(0.4),
		Amount:       float64(800),
		Alternatives: []models.Ingredient{suite.chives},
		Type:         models.Weight,
	}

	leek2 := models.Ingredient{
		ID:           int64(3),
		Name:         "NotLeek",
		Description:  "Floral green herb with more cronch",
		Price:        float64(0.4),
		Amount:       float64(800),
		Alternatives: []models.Ingredient{suite.chives},
		Type:         models.Weight,
	}

	chives1 := models.Ingredient{
		ID:          int64(1),
		Name:        "Chives",
		Description: "Floral green herbs",
		Price:       float64(0.5),
		Amount:      float64(500),
		Type:        models.Weight,
	}

	chives2 := models.Ingredient{
		ID:          int64(4),
		Name:        "NoTChives",
		Description: "Floral green herbs",
		Price:       float64(0.5),
		Amount:      float64(500),
		Type:        models.Weight,
	}

	assert.False(suite.T(), suite.chives.Equals(suite.leek))
	assert.False(suite.T(), chives2.Equals(suite.chives))
	assert.False(suite.T(), leek2.Equals(suite.leek))
	assert.True(suite.T(), leek1.Equals(suite.leek))
	assert.True(suite.T(), chives1.Equals(suite.chives))

}

func (suite *IngredientTestSuite) TestIngredientExists() {
	//insert the chives and leek into the ingredient table
	_, err := suite.app.DB.Exec(suite.insertIngredients, suite.chives.Name, suite.chives.Description, suite.chives.Price, suite.chives.Amount, suite.chives.Type)
	if err != nil {
		suite.T().Errorf("Error inserting chives: %s", err.Error())
	}
	_, err = suite.app.DB.Exec(suite.insertIngredients, suite.leek.Name, suite.leek.Description, suite.leek.Price, suite.leek.Amount, suite.leek.Type)
	if err != nil {
		suite.T().Errorf("Error inserting leek: %s", err.Error())
	}

	//check if chives and leek exist
	exists := suite.app.IngredientExists(suite.chives)
	assert.True(suite.T(), exists)
	exists = suite.app.IngredientExists(suite.leek)
	assert.True(suite.T(), exists)
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
