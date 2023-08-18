package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/rainbowriverrr/CafeGo/internal/models"
	"github.com/rainbowriverrr/CafeGo/pkg/database"
	_ "modernc.org/sqlite"
)

type AdminTestSuite struct {
	suite.Suite
	app         models.Application
	selectQuery string
	insertQuery string
	deleteQuery string
}

func (suite *AdminTestSuite) SetupTest() {
	suite.app = models.Application{
		DB: database.InitDB("adminDB"),
	}
	models.CreateUserTables(suite.app.DB)
	models.CreateAdminTables(suite.app.DB)
	suite.selectQuery = "SELECT id, email, password FROM admins WHERE email = ?"
	suite.insertQuery = "INSERT INTO admins (user_id, email, password) VALUES (?, ?, ?)"
	suite.deleteQuery = "DELETE FROM admins WHERE email = ?"
}

func (suite *AdminTestSuite) TearDownTest() {
	database.DeleteDB("adminDB")
}

func (suite *AdminTestSuite) TestCreateAdminTable() {
	//check if table exists
	exists, err := database.TableExists(suite.app.DB, models.UserTable)
	if err != nil {
		suite.T().Errorf("Error checking if table exists: %s", err.Error())
	}
	assert.True(suite.T(), exists)
}

func (suite *AdminTestSuite) TestAdminEquals() {

}

func (suite *AdminTestSuite) TestGetAdminByEmail() {

}

func (suite *AdminTestSuite) TestUpdateAdminByEmail() {

}

func (suite *AdminTestSuite) TestInsertAdmin() {

}

func TestAdmin(t *testing.T) {
	suite.Run(t, new(AdminTestSuite))
}
