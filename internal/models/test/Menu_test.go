package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/rainbowriverrr/CafeGo/internal/models"
	"github.com/rainbowriverrr/CafeGo/pkg/database"
)

type MenuTestSuite struct {
	suite.Suite
	app        models.Application
	selectMenu string
}

func (suite *MenuTestSuite) SetupTest() {
	suite.app = models.Application{
		DB: database.InitDB("testMenu"),
	}
	models.CreateMenuTables(suite.app.DB)
}

func (suite *MenuTestSuite) TearDownTest() {
	database.DeleteDB("testMenu")
}

func (suite *MenuTestSuite) TestCreateMenuTables() {
	exists, err := database.TableExists(suite.app.DB, models.MenuTable)
	if err != nil {
		suite.T().Errorf("Error checking if table exists: %s", err.Error())
	}
	assert.True(suite.T(), exists)

	exists, err = database.TableExists(suite.app.DB, models.ProductToMenuTable)
	if err != nil {
		suite.T().Errorf("Error checking if table exists: %s", err.Error())
	}
	assert.True(suite.T(), exists)
}

func (suite *MenuTestSuite) TestMenuExists() {

}

func (suite *MenuTestSuite) TestGetMenu() {

}

func (suite *MenuTestSuite) TestGetMenus() {

}

func (suite *MenuTestSuite) TestUpdateMenu() {

}

func (suite *MenuTestSuite) TestDeleteMenu() {

}

func (suite *MenuTestSuite) TestInsertMenu() {

}

func (suite *MenuTestSuite) TestAddProductToMenu() {

}

func (suite *MenuTestSuite) TestRemoveProductFromMenu() {

}

func (suite *MenuTestSuite) TestGetProductsFromMenu() {

}

func TestMenu(t *testing.T) {
	suite.Run(t, new(MenuTestSuite))
}
