package models_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/rainbowriverrr/CafeGo/internal/models"
)

type ProductTestSuite struct {
	suite.Suite
	app                      models.Application
	selectProducts           string
	selectProductIngredients string
}

func (suite *IngredientTestSuite) TestProductEquals() {
}

func (suite *IngredientTestSuite) TestProductExists() {

}

func (suite *IngredientTestSuite) TestGetProduct() {

}

func (suite *IngredientTestSuite) TestUpdateProduct() {

}

func (suite *IngredientTestSuite) TestInsertProduct() {

}

func (suite *IngredientTestSuite) TestDeleteProduct() {

}

func (suite *IngredientTestSuite) TestGetProductIngredients() {

}

func (suite *IngredientTestSuite) TestInsertProductIngredients() {

}

func (suite *IngredientTestSuite) TestDeleteProductIngredients() {

}

func TestProduct(t *testing.T) {
	suite.Run(t, new(ProductTestSuite))
}
