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

func createUserTables(db *sql.DB) {
	users := `CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		phone TEXT,
		rewards INTEGER
	);`

	_, err := db.Exec(fmt.Sprintf(users, UserTable))
	if err != nil {
		panic(err)
	}
}
