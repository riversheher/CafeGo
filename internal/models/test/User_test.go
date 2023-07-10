package models_test

import (
	"testing"

	"github.com/rainbowriverrr/CafeGo/internal/models"
	_ "modernc.org/sqlite"
)

func init() {

}

func TestCreateUserTables(t *testing.T) {
	models.CreateUserTables(app.DB)

	//check if table exists
	query := "SELECT name FROM sqlite_schema WHERE type='table' AND name = ?;"
	rows, err := app.DB.Query(query, models.UserTable)
	if err != nil {
		t.Errorf("Error querying database: %s", err.Error())
	}

	defer rows.Close()

	if rows.Next() == false {
		t.Errorf("Table %s does not exist", models.UserTable)
	}
}

func TestUserExists(t *testing.T) {
	//create table and insert user
	models.CreateUserTables(app.DB)
	user := models.User{
		Name:  "test",
		Phone: "1234567890",
	}
	app.InsertUser(user)

	//check if user exists
	exists := app.UserExists(user)
	if exists == false {
		t.Errorf("User %s does not exist", user.Name)
	}
}

func TestGetUserByPhone(t *testing.T) {

}
