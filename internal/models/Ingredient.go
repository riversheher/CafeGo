package models

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

// Enum for units of measurement does not specify units, just type of unit (for better compatibility from imperial to metric)
type Unit int

const (
	Weight Unit = iota
	Volume
	Count
)

func (u Unit) String() string {
	switch u {
	case Weight:
		return "Weight"
	case Volume:
		return "Volume"
	case Count:
		return "Count"
	}
	return "Unknown"
}

type Ingredient struct {
	ID           int64        `json:"id"`
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	Price        float64      `json:"price"`
	Alternatives []Ingredient `json:"alternatives"`
	Amount       float64      `json:"amount"`
	Type         Unit         `json:"type"`
}

const (
	IngredientTable   = "ingredients"
	AlternativesTable = "alternatives"
)

func createIngredientTables(db *sql.DB) {
	ingredients := `CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		description TEXT,
		price REAL,
		amount REAL,
		type INTEGER
	);`

	alternatives := `CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		ingredient_id INTEGER,
		alternative_id INTEGER,
		FOREIGN KEY(ingredient_id) REFERENCES ingredients(id) ON DELETE CASCADE ON UPDATE CASCADE,
		FOREIGN KEY(alternative_id) REFERENCES ingredients(id) ON DELETE CASCADE ON UPDATE CASCADE
	);`

	_, err := db.Exec(fmt.Sprintf(ingredients, IngredientTable))
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(fmt.Sprintf(alternatives, AlternativesTable))
	if err != nil {
		panic(err)
	}
}
