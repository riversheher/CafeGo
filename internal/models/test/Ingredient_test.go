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
	onion              models.Ingredient
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
		ID:          int64(2),
		Name:        "Leek",
		Description: "Floral green herb with more cronch",
		Price:       float64(0.4),
		Amount:      float64(800),
		Type:        models.Weight,
	}
	suite.onion = models.Ingredient{
		ID:          int64(3),
		Name:        "Onion",
		Description: "Floral green herb with more cronch",
		Price:       float64(0.4),
		Amount:      float64(800),
		Type:        models.Count,
	}
	suite.chives.Alternatives = []int64{suite.leek.ID, suite.onion.ID}
	suite.leek.Alternatives = []int64{suite.chives.ID, suite.onion.ID}
	suite.onion.Alternatives = []int64{suite.chives.ID, suite.leek.ID}

	suite.selectIngredients = "SELECT id, name, description, price, amount, type FROM ingredients WHERE id = ?"
	suite.selectAlternatives = "SELECT id, ingredient_id, alternative_id FROM alternatives WHERE id = ?"
	suite.insertIngredients = "INSERT INTO ingredients (name, description, price, amount, type) VALUES (?, ?, ?, ?, ?)"
	suite.insertAlternatives = "INSERT INTO alternatives (ingredient_id, alternative_id) VALUES (?, ?)"
	suite.deleteIngredients = "DELETE FROM ingredients WHERE id = ?"
	suite.deleteAlternatives = "DELETE FROM alternatives WHERE id = ?"

	//build tables
	models.CreateIngredientTables(suite.app.DB)

	//insert ingredients
	for _, ingredient := range []models.Ingredient{suite.chives, suite.leek, suite.onion} {
		result, err := suite.app.DB.Exec(suite.insertIngredients, ingredient.Name, ingredient.Description, ingredient.Price, ingredient.Amount, ingredient.Type)
		if err != nil {
			suite.T().Errorf("Error inserting ingredient: %s", err.Error())
		}
		ingredient.ID, err = result.LastInsertId()
		if err != nil {
			suite.T().Errorf("Error getting last insert id: %s", err.Error())
		}
	}

	//insert alternatives
	for _, ingredient := range []models.Ingredient{suite.chives, suite.leek, suite.onion} {
		for _, alternative := range ingredient.Alternatives {
			_, err := suite.app.DB.Exec(suite.insertAlternatives, ingredient.ID, alternative)
			if err != nil {
				suite.T().Errorf("Error inserting alternative: %s", err.Error())
			}
		}
	}
}

func (suite *IngredientTestSuite) TearDownTest() {
	database.DeleteDB("testIngredient")
}

func (suite *IngredientTestSuite) TestIngredientEquals() {
	assert.True(suite.T(), suite.chives.Equals(suite.chives))
	assert.False(suite.T(), suite.chives.Equals(suite.leek))
}

func (suite *IngredientTestSuite) TestGetIngredient() {
	ingredient, err := suite.app.GetIngredient(suite.chives.ID)
	if err != nil {
		suite.T().Errorf("Error getting ingredient: %s", err.Error())
	}
	assert.True(suite.T(), ingredient.Equals(suite.chives))
}

func (suite *IngredientTestSuite) TestUpdateIngredient() {
	suite.chives.Description = "Floral green herb"
	err := suite.app.UpdateIngredient(suite.chives)
	if err != nil {
		suite.T().Errorf("Error updating ingredient: %s", err.Error())
	}
	ingredient := models.Ingredient{}
	err = suite.app.DB.QueryRow(suite.selectIngredients, suite.chives.ID).Scan(&ingredient.ID, &ingredient.Name, &ingredient.Description, &ingredient.Price, &ingredient.Amount, &ingredient.Type)
	if err != nil {
		suite.T().Errorf("Error getting ingredient: %s", err.Error())
	}
	assert.True(suite.T(), ingredient.Equals(suite.chives))
}

func (suite *IngredientTestSuite) TestInsertIngredient() {
	ingredient := models.Ingredient{
		Name:         "Garlic",
		Description:  "Floral green herb with more cronch",
		Price:        float64(0.4),
		Amount:       float64(800),
		Alternatives: []int64{},
		Type:         models.Count,
	}
	id, err := suite.app.InsertIngredient(ingredient)
	if err != nil {
		suite.T().Errorf("Error inserting ingredient: %s", err.Error())
	}
	ingredient.ID = id

	newIngredient := models.Ingredient{}
	err = suite.app.DB.QueryRow(suite.selectIngredients, ingredient.ID).Scan(&newIngredient.ID, &newIngredient.Name, &newIngredient.Description, &newIngredient.Price, &newIngredient.Amount, &newIngredient.Type)
	if err != nil {
		suite.T().Errorf("Error getting ingredient: %s", err.Error())
	}
	assert.True(suite.T(), ingredient.Equals(newIngredient))
}

func (suite *IngredientTestSuite) TestDeleteIngredient() {
	err := suite.app.DeleteIngredient(suite.onion)
	if err != nil {
		suite.T().Errorf("Error deleting ingredient: %s", err.Error())
	}
	rows := suite.app.DB.QueryRow(suite.selectIngredients, suite.onion.ID)
	err = rows.Scan()
	if err == nil {
		suite.T().Errorf("Ingredient not deleted")
	}
}

func (suite *IngredientTestSuite) TestGetAlternatives() {
	alternatives, err := suite.app.GetAlternatives(suite.chives.ID)
	if err != nil {
		suite.T().Errorf("Error getting alternatives: %s", err.Error())
	}
	assert.Equal(suite.T(), suite.chives.Alternatives, alternatives)

}

func (suite *IngredientTestSuite) TestDeleteAlternative() {
	err := suite.app.DeleteAlternative(suite.leek.ID, suite.chives.ID)
	if err != nil {
		suite.T().Errorf("Error deleting alternative: %s", err.Error())
	}
	rows := suite.app.DB.QueryRow(suite.selectAlternatives, suite.chives.ID, suite.leek.ID)
	err = rows.Scan()
	if err == nil {
		suite.T().Errorf("Alternative not deleted")
	}
}

func TestIngredient(t *testing.T) {
	suite.Run(t, new(IngredientTestSuite))
}
