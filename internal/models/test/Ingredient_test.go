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
		Alternatives: []int64{suite.chives.ID},
		Type:         models.Weight,
	}
	suite.chives.Alternatives = []int64{suite.leek.ID}

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
		Alternatives: []int64{int64(3)},
		Type:         models.Weight,
	}

	leek2 := models.Ingredient{
		ID:           int64(3),
		Name:         "NotLeek",
		Description:  "Floral green herb with more cronch",
		Price:        float64(0.4),
		Amount:       float64(800),
		Alternatives: []int64{int64(2)},
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
	result, err := suite.app.DB.Exec(suite.insertIngredients, suite.chives.Name, suite.chives.Description, suite.chives.Price, suite.chives.Amount, suite.chives.Type)
	if err != nil {
		suite.T().Errorf("Error inserting chives: %s", err.Error())
	}
	chivesID, err := result.LastInsertId()
	if err != nil {
		suite.T().Errorf("Error getting chives ID: %s", err.Error())
	}

	result, err = suite.app.DB.Exec(suite.insertIngredients, suite.leek.Name, suite.leek.Description, suite.leek.Price, suite.leek.Amount, suite.leek.Type)
	if err != nil {
		suite.T().Errorf("Error inserting leek: %s", err.Error())
	}
	leekID, err := result.LastInsertId()
	if err != nil {
		suite.T().Errorf("Error getting leek ID: %s", err.Error())
	}

	//check if chives and leek exist
	exists := suite.app.IngredientExists(suite.chives)
	assert.True(suite.T(), exists)
	exists = suite.app.IngredientExists(suite.leek)
	assert.True(suite.T(), exists)

	//delete chives and leek from the ingredient table
	_, err = suite.app.DB.Exec(suite.deleteIngredients, chivesID)
	if err != nil {
		suite.T().Errorf("Error deleting chives: %s", err.Error())
	}

	_, err = suite.app.DB.Exec(suite.deleteIngredients, leekID)
	if err != nil {
		suite.T().Errorf("Error deleting leek: %s", err.Error())
	}
}

func (suite *IngredientTestSuite) TestGetIngredient() {
	//insert the chives and leek into the ingredient table
	result, err := suite.app.DB.Exec(suite.insertIngredients, suite.chives.Name, suite.chives.Description, suite.chives.Price, suite.chives.Amount, suite.chives.Type)
	if err != nil {
		suite.T().Errorf("Error inserting chives: %s", err.Error())
	}
	chivesID, err := result.LastInsertId()
	if err != nil {
		suite.T().Errorf("Error getting chives ID: %s", err.Error())
	}

	result, err = suite.app.DB.Exec(suite.insertIngredients, suite.leek.Name, suite.leek.Description, suite.leek.Price, suite.leek.Amount, suite.leek.Type)
	if err != nil {
		suite.T().Errorf("Error inserting leek: %s", err.Error())
	}
	leekID, err := result.LastInsertId()
	if err != nil {
		suite.T().Errorf("Error getting leek ID: %s", err.Error())
	}

	//get chives and leek from the ingredient table
	chives, err := suite.app.GetIngredient(chivesID)
	if err != nil {
		suite.T().Errorf("Error getting chives: %s", err.Error())
	}
	assert.Equal(suite.T(), suite.chives.Name, chives.Name)
	assert.Equal(suite.T(), suite.chives.Amount, chives.Amount)
	assert.Equal(suite.T(), suite.chives.Description, chives.Description)
	assert.Equal(suite.T(), suite.chives.Price, chives.Price)

	leek, err := suite.app.GetIngredient(leekID)
	if err != nil {
		suite.T().Errorf("Error getting leek: %s", err.Error())
	}
	assert.Equal(suite.T(), suite.leek.Name, leek.Name)
	assert.Equal(suite.T(), suite.leek.Amount, leek.Amount)
	assert.Equal(suite.T(), suite.leek.Description, leek.Description)
	assert.Equal(suite.T(), suite.leek.Price, leek.Price)

	//delete chives and leek from the ingredient table
	_, err = suite.app.DB.Exec(suite.deleteIngredients, chivesID)
	if err != nil {
		suite.T().Errorf("Error deleting chives: %s", err.Error())
	}

	_, err = suite.app.DB.Exec(suite.deleteIngredients, leekID)
	if err != nil {
		suite.T().Errorf("Error deleting leek: %s", err.Error())
	}
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
