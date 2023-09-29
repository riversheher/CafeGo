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
	//insert the chives into the ingredient table
	result, err := suite.app.DB.Exec(suite.insertIngredients, suite.chives.Name, suite.chives.Description, suite.chives.Price, suite.chives.Amount, suite.chives.Type)
	if err != nil {
		suite.T().Errorf("Error inserting chives: %s", err.Error())
	}
	chivesID, err := result.LastInsertId()
	if err != nil {
		suite.T().Errorf("Error getting chives ID: %s", err.Error())
	}

	//update the chives
	chives2 := models.Ingredient{
		ID:          chivesID,
		Name:        "Chives2",
		Description: "Floral green herbs2",
		Price:       float64(0.52),
		Amount:      float64(502),
		Type:        models.Weight,
	}

	err = suite.app.UpdateIngredient(chives2)

	//get chives from the ingredient table
	chives, err := suite.app.GetIngredient(chivesID)
	if err != nil {
		suite.T().Errorf("Error getting chives: %s", err.Error())
	}
	assert.Equal(suite.T(), chives2.Name, chives.Name)
	assert.Equal(suite.T(), chives2.Amount, chives.Amount)
	assert.Equal(suite.T(), chives2.Description, chives.Description)
	assert.Equal(suite.T(), chives2.Price, chives.Price)
	assert.Equal(suite.T(), chives2.Amount, chives.Amount)

	//delete chives from the ingredient table
	_, err = suite.app.DB.Exec(suite.deleteIngredients, chivesID)

}

func (suite *IngredientTestSuite) TestInsertIngredient() {
	suite.app.InsertIngredient(suite.chives)
	//get chives from the ingredient table
	chives, err := suite.app.GetIngredient(suite.chives.ID)
	if err != nil {
		suite.T().Errorf("Error getting chives: %s", err.Error())
	}
	assert.Equal(suite.T(), suite.chives.Name, chives.Name)
	assert.Equal(suite.T(), suite.chives.Amount, chives.Amount)
	assert.Equal(suite.T(), suite.chives.Description, chives.Description)
	assert.Equal(suite.T(), suite.chives.Price, chives.Price)
	assert.Equal(suite.T(), suite.chives.Amount, chives.Amount)

	//delete chives from the ingredient table
	_, err = suite.app.DB.Exec(suite.deleteIngredients, suite.chives.ID)
}

func (suite *IngredientTestSuite) TestDeleteIngredient() {
	//insert the chives into the ingredient table
	result, err := suite.app.DB.Exec(suite.insertIngredients, suite.chives.Name, suite.chives.Description, suite.chives.Price, suite.chives.Amount, suite.chives.Type)
	if err != nil {
		suite.T().Errorf("Error inserting chives: %s", err.Error())
	}
	chivesID, err := result.LastInsertId()
	if err != nil {
		suite.T().Errorf("Error getting chives ID: %s", err.Error())
	}

	//delete chives from the ingredient table
	err = suite.app.DeleteIngredient(suite.chives)
	if err != nil {
		suite.T().Errorf("Error deleting chives: %s", err.Error())
	}

	//check if chives exists
	exists := suite.app.IngredientExists(suite.chives)
	assert.False(suite.T(), exists)

	//delete chives from the ingredient table
	_, err = suite.app.DB.Exec(suite.deleteIngredients, chivesID)
	if err != nil {
		suite.T().Errorf("Error deleting chives: %s", err.Error())
	}

}

func (suite *IngredientTestSuite) TestGetAlternatives() {
	//insert the chives
	result, err := suite.app.DB.Exec(suite.insertIngredients, suite.chives.Name, suite.chives.Description, suite.chives.Price, suite.chives.Amount, suite.chives.Type)
	if err != nil {
		suite.T().Errorf("Error inserting chives: %s", err.Error())
	}
	chivesID, err := result.LastInsertId()
	if err != nil {
		suite.T().Errorf("Error getting chives ID: %s", err.Error())
	}

	//insert the leek
	_, err = suite.app.DB.Exec(suite.insertIngredients, suite.leek.Name, suite.leek.Description, suite.leek.Price, suite.leek.Amount, suite.leek.Type)
	if err != nil {
		suite.T().Errorf("Error inserting leek: %s", err.Error())
	}

	//insert the onion
	_, err = suite.app.DB.Exec(suite.insertIngredients, suite.onion.Name, suite.onion.Description, suite.onion.Price, suite.onion.Amount, suite.onion.Type)
	if err != nil {
		suite.T().Errorf("Error inserting onion: %s", err.Error())
	}

	//insert the alternatives
	for _, alternative := range suite.chives.Alternatives {
		_, err = suite.app.DB.Exec(suite.insertAlternatives, chivesID, alternative)
		if err != nil {
			suite.T().Errorf("Error inserting alternatives: %s", err.Error())
		}
	}

	//get alternatives
	alternatives, err := suite.app.GetAlternatives(chivesID)
	if err != nil {
		suite.T().Errorf("Error getting alternatives: %s", err.Error())
	}
	assert.True(suite.T(), alternatives[0] == suite.leek.ID || alternatives[0] == suite.onion.ID)
	assert.True(suite.T(), alternatives[1] == suite.leek.ID || alternatives[1] == suite.onion.ID)

	//delete chives, leek, onion and alternatives from the ingredient table
	_, err = suite.app.DB.Exec(suite.deleteIngredients, chivesID)
	if err != nil {
		suite.T().Errorf("Error deleting chives: %s", err.Error())
	}
	_, err = suite.app.DB.Exec(suite.deleteIngredients, suite.leek.ID)
	if err != nil {
		suite.T().Errorf("Error deleting leek: %s", err.Error())
	}
	_, err = suite.app.DB.Exec(suite.deleteIngredients, suite.onion.ID)
	if err != nil {
		suite.T().Errorf("Error deleting onion: %s", err.Error())
	}
}

func (suite *IngredientTestSuite) TestInsertAlternative() {
	//insert ingredients
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

	//insert alternatives
	_, err = suite.app.AddAlternative(chivesID, leekID)
	if err != nil {
		suite.T().Errorf("Error inserting alternatives: %s", err.Error())
	}

	//get alternatives
	alternatives, err := suite.app.GetAlternatives(chivesID)
	if err != nil {
		suite.T().Errorf("Error getting alternatives: %s", err.Error())
	}
	assert.True(suite.T(), alternatives[0] == suite.leek.ID)
}

func (suite *IngredientTestSuite) TestDeleteAlternative() {
	//insert ingredients
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

	//insert alternatives
	_, err = suite.app.DB.Exec(suite.insertAlternatives, chivesID, leekID)
	if err != nil {
		suite.T().Errorf("Error inserting alternatives: %s", err.Error())
	}

	//delete alternatives
	err = suite.app.DeleteAlternative(suite.chives.ID, suite.leek.ID)
	if err != nil {
		suite.T().Errorf("Error deleting alternatives: %s", err.Error())
	}

	//get alternatives
	alternatives, err := suite.app.GetAlternatives(suite.chives.ID)
	if err != nil {
		suite.T().Errorf("Error getting alternatives: %s", err.Error())
	}
	assert.True(suite.T(), len(alternatives) == 0)
}

func TestIngredient(t *testing.T) {
	suite.Run(t, new(IngredientTestSuite))
}
