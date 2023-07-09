package database

import (
	"os"

	"database/sql"

	_ "modernc.org/sqlite"
)

// create a sqlite database file
func createDB(name string) {
	file, err := os.Create(name + ".db")
	if err != nil {
		panic(err)
	}
	file.Close()
}

// check if a sqlite database file exists, returns true if it does, false otherwise
func dbExists(name string) bool {
	_, err := os.Stat(name + ".db")
	return !os.IsNotExist(err)
}

// delete a sqlite database file
func deleteDB(name string) {
	err := os.Remove(name + ".db")
	if err != nil {
		panic(err)
	}
}

func TableExists(db *sql.DB, tableName string) (bool, error) {
	var exists bool
	err := db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name=?", tableName).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

// Initializes the application database, returns db connection
func InitDB(appName string) *sql.DB {
	if !dbExists(appName) {
		createDB(appName)
	}

	db, err := sql.Open("sqlite", appName+".db")
	if err != nil {
		panic(err)
	}

	return db
}
