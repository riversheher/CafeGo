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

Personal Information
Preffered Name: River Wang
First Name: Tong Zhang
Last Name: Wang
Date of Birth: 2002/08/30
Address: 36 Brampton Cres SW, Calgary, AB T2W 0Y4
Nationality: Canadian
Email: Friver980@outlook.com
Contact person in case of emergency in Thailand: Grace You
Relationship with contact person: Close friend
Telephone Number: +1 707-776-6228
Email: minakochan@gmail.co
Do you need an interpreter? No, but I speak English

Wire Transfer Information
Sender's First Name: Tongzhang-Wang
Sender's Address and Country: 36 Brampton Cres SW, Calgary, AB T2W 0Y4, Canada
Patients name: River Wang
Amount Transferred: 110500 THB
Transfer Date: June 10th 2023
Sender's Bank Name: Simplii Financial
Sender's Bank Address: Simplii Financial P.O. Box 603, Station Agin Court, Scarbourough, ON M1S 5K9, Canada
Expected Surgery Date: February 16th 2024
List of Surgical Procedures:
1. FOREHEAD/ BROW RIDGE BONE CONTOURING
2. JAW AND CHIN CONTOURING SURGERY
