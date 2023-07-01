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

// create Tables for the database
func createTables(db *sql.DB) {

}

// Initializes the application database, returns db connection
func InitDB(appName string) *sql.DB {
	if !dbExists(appName) {
		createDB(appName)
	}

	db, err := sql.Open("sqlite3", appName+".db")
	if err != nil {
		panic(err)
	}

	return db
}
