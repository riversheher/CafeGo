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
	selectQuery string //removes dependency on the methods in the models package
	insertQuery string //removes dependency on the methods in the models package
	deleteQuery string //removes dependency on the methods in the models package
}

func (suite *UserTestSuite) SetupTest() {
	suite.app = models.Application{
		DB: database.InitDB("testUser"),
	}
	models.CreateUserTables(suite.app.DB)
	suite.selectQuery = "SELECT id, name, phone, rewards FROM users WHERE phone = ?"
	suite.insertQuery = "INSERT INTO users (name, phone, rewards) VALUES (?, ?, ?)"
	suite.deleteQuery = "DELETE FROM users WHERE phone = ?"
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
	user1 := models.User{
		ID:      1,
		Name:    "test",
		Phone:   "1234567890",
		Rewards: 0,
	}
	eqUser1 := models.User{
		ID:      1,
		Name:    "dslkfjsdl",
		Phone:   "1234567890",
		Rewards: 0,
	}
	alsoEQUser1 := models.User{
		ID:      -1,
		Name:    "test",
		Phone:   "1234567890",
		Rewards: 0,
	}
	user2 := models.User{
		ID:      2,
		Name:    "test",
		Phone:   "0987654321",
		Rewards: 0,
	}

	assert.True(suite.T(), user1.Equals(eqUser1))
	assert.True(suite.T(), user1.Equals(alsoEQUser1))
	assert.False(suite.T(), user1.Equals(user2))
}

func (suite *UserTestSuite) TestUserExists() {
	user1 := models.User{
		ID:      1,
		Name:    "test",
		Phone:   "1234567890",
		Rewards: 0,
	}
	user2 := models.User{
		ID:      2,
		Name:    "test",
		Phone:   "0987654321",
		Rewards: 0,
	}
	anotherUser1 := models.User{
		ID:      3,
		Name:    "test",
		Phone:   "1234567890",
		Rewards: 0,
	}

	//check if user exists
	exists := suite.app.UserExists(user1)
	assert.False(suite.T(), exists)
	exists = suite.app.UserExists(user2)
	assert.False(suite.T(), exists)

	//insert user
	_, err := suite.app.DB.Exec(suite.insertQuery, user1.Name, user1.Phone, user1.Rewards)
	if err != nil {
		suite.T().Errorf("Error inserting user: %s", err.Error())
	}

	//check if user exists
	exists = suite.app.UserExists(user1)
	assert.True(suite.T(), exists)
	exists = suite.app.UserExists(user2)
	assert.False(suite.T(), exists)
	exists = suite.app.UserExists(anotherUser1)
	assert.True(suite.T(), exists)

	//delete user
	_, err = suite.app.DB.Exec(suite.deleteQuery, user1.Phone)
	if err != nil {
		suite.T().Errorf("Error deleting user: %s", err.Error())
	}

}

func (suite *UserTestSuite) TestGetUserByPhone() {
	user1 := models.User{
		ID:      1,
		Name:    "test",
		Phone:   "1234567890",
		Rewards: 0,
	}
	user2 := models.User{
		ID:      2,
		Name:    "test",
		Phone:   "0987654321",
		Rewards: 0,
	}

	//insert users
	_, err := suite.app.DB.Exec(suite.insertQuery, user1.Name, user1.Phone, user1.Rewards)
	if err != nil {
		suite.T().Errorf("Error inserting user: %s", err.Error())
	}

	_, err = suite.app.DB.Exec(suite.insertQuery, user2.Name, user2.Phone, user2.Rewards)
	if err != nil {
		suite.T().Errorf("Error inserting user: %s", err.Error())
	}

	//get user
	user := suite.app.GetUserByPhone(user1.Phone)
	assert.Equal(suite.T(), user1, user)
	assert.NotEqual(suite.T(), user2, user)

	//delete users
	_, err = suite.app.DB.Exec(suite.deleteQuery, user1.Phone)
	if err != nil {
		suite.T().Errorf("Error deleting user: %s", err.Error())
	}
}

func (suite *UserTestSuite) UpdateUserByPhone() {
	user1 := models.User{
		ID:      1,
		Name:    "test",
		Phone:   "1234567890",
		Rewards: 0,
	}
	updateUser1 := models.User{
		ID:      1,
		Name:    "testing",
		Phone:   "1234567890",
		Rewards: 0,
	}

	//insert user
	_, err := suite.app.DB.Exec(suite.insertQuery, user1.Name, user1.Phone, user1.Rewards)
	if err != nil {
		suite.T().Errorf("Error inserting user: %s", err.Error())
	}

	//update user
	suite.app.UpdateUserByPhone(updateUser1)

	//get user
	user := suite.app.DB.QueryRow(suite.selectQuery, user1.Phone)
	assert.Equal(suite.T(), updateUser1, user)
	assert.NotEqual(suite.T(), user1, user)

	//delete user
	_, err = suite.app.DB.Exec(suite.deleteQuery, user1.Phone)
	if err != nil {
		suite.T().Errorf("Error deleting user: %s", err.Error())
	}
}

func (suite *UserTestSuite) TestInsertUser() {
	user := models.User{
		Name:    "test",
		Phone:   "1111111111",
		Rewards: 0,
	}

	user2 := models.User{
		Name:    "test2",
		Phone:   "2222222222",
		Rewards: 0,
	}

	//insert user
	returningUser, err := suite.app.InsertUser(user)
	if err != nil {
		suite.T().Errorf("Error inserting user: %s", err.Error())
	}
	assert.Equal(suite.T(), int64(1), returningUser.ID)

	//get user
	row := suite.app.DB.QueryRow(suite.selectQuery, user.Phone)
	var id int64
	var name string
	var phone string
	var rewards int64
	err = row.Scan(&id, &name, &phone, &rewards)
	if err != nil {
		suite.T().Errorf("Error getting user: %s", err.Error())
	}

	//check if user is correct
	assert.Equal(suite.T(), int64(1), id)
	assert.Equal(suite.T(), user.Name, name)
	assert.Equal(suite.T(), user.Phone, phone)
	assert.Equal(suite.T(), user.Rewards, rewards)

	//insert new user with auto increment
	returningUser, err = suite.app.InsertUser(user2)
	if err != nil {
		suite.T().Errorf("Error inserting user: %s", err.Error())
	}
	assert.Equal(suite.T(), int64(2), returningUser.ID)
}

func TestUser(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
