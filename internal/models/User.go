package models

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type User struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Rewards int64  `json:"rewards"`
}

const (
	UserTable = "users"
)

// users are equal if their IDs or phone numbers are equal (phone numbers are unique)
// It is possible for the IDs to not be equal if we don't know it at the time of input
func (this User) Equals(other User) bool {
	return (this.ID == other.ID || this.Phone == other.Phone)
}

func CreateUserTables(db *sql.DB) {
	users := `CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		phone TEXT,
		rewards INTEGER,
		UNIQUE(phone)
	);`

	_, err := db.Exec(fmt.Sprintf(users, UserTable))
	if err != nil {
		panic(err)
	}
}

func (app *Application) UserExists(user User) bool {
	var exists bool
	var phone string
	query := fmt.Sprintf("SELECT phone FROM %s WHERE phone = ?", UserTable)
	err := app.DB.QueryRow(query, user.Phone).Scan(&phone)

	if err != nil {
		//checks if the error is not a "no rows" error, meaning the error isn't that the user doesn't exist
		if err != sql.ErrNoRows {
			app.ErrLog.Println(err)
		}
		exists = false
	} else {
		exists = true
	}

	return exists

}

func (app *Application) GetUserByPhone(phone string) User {
	var user User
	query := fmt.Sprintf("SELECT * FROM %s WHERE phone = ?", UserTable)
	err := app.DB.QueryRow(query, phone).Scan(&user.ID, &user.Name, &user.Phone, &user.Rewards)
	if err != nil {
		app.ErrLog.Println(err)
	}
	return user
}

func (app *Application) UpdateUserByPhone(user User) {
	query := fmt.Sprintf("UPDATE %s SET name = ? rewards = ? WHERE phone = ?", UserTable)
	_, err := app.DB.Exec(query, user.Name, user.Rewards, user.Phone)
	if err != nil {
		app.ErrLog.Println(err)
	}
}

func (app *Application) InsertUser(user User) {
	query := fmt.Sprintf("INSERT INTO %s (name, phone, rewards) VALUES (?, ?, ?)", UserTable)
	_, err := app.DB.Exec(query, user.Name, user.Phone, user.Rewards)
	if err != nil {
		app.ErrLog.Println(err)
	}
}
