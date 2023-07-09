package models

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

const (
	AdminTable = "admins"
)

type Admin struct {
	User
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateAdminTables(db *sql.DB) {
	admins := `CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		email TEXT,
		password TEXT,
		FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
	);`

	_, err := db.Exec(fmt.Sprintf(admins, AdminTable))

	if err != nil {
		panic(err)
	}
}
