package models_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/rainbowriverrr/CafeGo/internal/models"
	"github.com/rainbowriverrr/CafeGo/pkg/database"
)

type OrderTestSuite struct {
	suite.Suite
	app models.Application
}

func (suite *OrderTestSuite) SetupTest() {
	suite.app = models.Application{
		DB: database.InitDB("testOrder"),
	}
}

func (suite *OrderTestSuite) TearDownTest() {
	database.DeleteDB("testOrder")
}

func (suite *OrderTestSuite) TestCreateOrderTables() {
}

func (suite *OrderTestSuite) TestOrderEquals() {

}

func (suite *OrderTestSuite) TestOrderExists() {

}

func (suite *OrderTestSuite) TestGetOrder() {

}

func (suite *OrderTestSuite) TestGetAllOrders() {

}

func (suite *OrderTestSuite) TestGetInProgressOrders() {

}

func (suite *OrderTestSuite) TestGetCompletedOrders() {

}

func (suite *OrderTestSuite) TestUpdateOrder() {

}

func (suite *OrderTestSuite) TestInsertOrder() {

}

func (suite *OrderTestSuite) TestAddProductToOrder() {

}

func (suite *OrderTestSuite) TestRemoveProductFromOrder() {

}

func (suite *OrderTestSuite) TestGetProductsFromOrder() {

}

func TestOrder(t *testing.T) {
	suite.Run(t, new(OrderTestSuite))
}
