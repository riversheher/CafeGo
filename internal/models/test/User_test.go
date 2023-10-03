package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/rainbowriverrr/CafeGo/internal/models"
	"github.com/rainbowriverrr/CafeGo/pkg/database"
	_ "modernc.org/sqlite"
)

type UserTestSuite struct {
	suite.Suite
	app         models.Application
	user1       models.User
	user2       models.User
	user3       models.User
	selectQuery string //removes dependency on the methods in the models package
	insertQuery string //removes dependency on the methods in the models package
	deleteQuery string //removes dependency on the methods in the models package
}

func (suite *UserTestSuite) throwError(err error) {
	if err != nil {
		suite.T().Errorf("Error: %s", err.Error())
	}
}

func (suite *UserTestSuite) SetupTest() {
	suite.app = models.Application{
		DB: database.InitDB("testUser"),
	}
	models.CreateUserTables(suite.app.DB)
	suite.selectQuery = "SELECT id, name, phone, rewards FROM users WHERE phone = ?"
	suite.insertQuery = "INSERT INTO users (name, phone, rewards) VALUES (?, ?, ?)"
	suite.deleteQuery = "DELETE FROM users WHERE phone = ?"

	//insert test data
	suite.user1 = models.User{
		Name:    "User1",
		Phone:   "1111111111",
		Rewards: 0,
	}
	suite.user2 = models.User{
		Name:    "User2",
		Phone:   "2222222222",
		Rewards: 0,
	}
	suite.user3 = models.User{
		Name:    "User3",
		Phone:   "3333333333",
		Rewards: 0,
	}

	//initialize tables
	models.CreateUserTables(suite.app.DB)

	//insert test data
	result, err := suite.app.DB.Exec(suite.insertQuery, suite.user1.Name, suite.user1.Phone, suite.user1.Rewards)
	suite.throwError(err)
	suite.user1.ID, err = result.LastInsertId()
	suite.throwError(err)

	result, err = suite.app.DB.Exec(suite.insertQuery, suite.user2.Name, suite.user2.Phone, suite.user2.Rewards)
	suite.throwError(err)
	suite.user2.ID, err = result.LastInsertId()
	suite.throwError(err)

	result, err = suite.app.DB.Exec(suite.insertQuery, suite.user3.Name, suite.user3.Phone, suite.user3.Rewards)
	suite.throwError(err)
	suite.user3.ID, err = result.LastInsertId()
	suite.throwError(err)
}

func (suite *UserTestSuite) TearDownTest() {
	database.DeleteDB("testUser")
}

func (suite *UserTestSuite) TestCreateUserTables() {
	//check if table exists
	exists, err := database.TableExists(suite.app.DB, models.UserTable)
	if err != nil {
		suite.T().Errorf("Error checking if table exists: %s", err.Error())
	}
	assert.True(suite.T(), exists)
}

func (suite *UserTestSuite) TestUserEquals() {
	assert.True(suite.T(), suite.user1.Equals(suite.user1))
	assert.False(suite.T(), suite.user1.Equals(suite.user2))
}

func (suite *UserTestSuite) TestUserExists() {
	assert.True(suite.T(), suite.app.UserExists(suite.user1))
	assert.True(suite.T(), suite.app.UserExists(suite.user2))
	assert.True(suite.T(), suite.app.UserExists(suite.user3))
}

func (suite *UserTestSuite) TestUpdateUser() {
	suite.user3.Name = "User3Updated"
	suite.app.UpdateUser(suite.user3)
	var user models.User
	err := suite.app.DB.QueryRow(suite.selectQuery, suite.user3.Phone).Scan(&user.ID, &user.Name, &user.Phone, &user.Rewards)
	suite.throwError(err)
	assert.Equal(suite.T(), suite.user3.Name, user.Name)
}

func (suite *UserTestSuite) TestInsertUser() {
	user := models.User{
		Name:    "User4",
		Phone:   "4444444444",
		Rewards: 0,
	}
	result, err := suite.app.DB.Exec(suite.insertQuery, user.Name, user.Phone, user.Rewards)
	suite.throwError(err)
	user.ID, err = result.LastInsertId()
	suite.throwError(err)
	assert.True(suite.T(), suite.app.UserExists(user))
}

func TestUser(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
