package models

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type Admin struct {
	User
	Email    string `json:"email"`
	Password string `json:"password"`
}

func createAdminTable(db *sql.DB) {

}
