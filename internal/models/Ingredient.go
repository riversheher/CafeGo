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
	ID           int64   `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	Alternatives []int64 `json:"alternatives"`
	Amount       float64 `json:"amount"`
	Type         Unit    `json:"type"`
}

const (
	IngredientTable   = "ingredients"
	AlternativesTable = "alternatives"
)

func (a Ingredient) Equals(b Ingredient) bool {
	return (a.ID == b.ID &&
		a.Name == b.Name &&
		a.Description == b.Description &&
		a.Price == b.Price &&
		a.Amount == b.Amount &&
		a.Type == b.Type)
}

func CreateIngredientTables(db *sql.DB) {
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

func (app *Application) IngredientExists(ingredient Ingredient) bool {
	var exists bool
	var ID string
	query := fmt.Sprintf("SELECT id FROM %s WHERE id = ?", IngredientTable)
	err := app.DB.QueryRow(query, ingredient.ID).Scan(&ID)

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

func (app *Application) GetIngredient(id int64) (Ingredient, error) {
	var ingredient Ingredient = Ingredient{}
	query := fmt.Sprintf("SELECT id, name, description, price, amount, type FROM %s WHERE id = ?", IngredientTable)
	err := app.DB.QueryRow(query, id).Scan(&ingredient.ID, &ingredient.Name, &ingredient.Description, &ingredient.Price, &ingredient.Amount, &ingredient.Type)
	if err != nil {
		app.ErrLog.Println(err)
	}

	var alternatives []int64
	query = fmt.Sprintf("SELECT alternative_id FROM %s WHERE ingredient_id = ?", AlternativesTable)
	rows, err := app.DB.Query(query, id)
	if err != nil {
		app.ErrLog.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var alternative int64
		err = rows.Scan(&alternative)
		if err != nil {
			app.ErrLog.Println(err)
		}
		alternatives = append(alternatives, alternative)
	}

	ingredient.Alternatives = alternatives

	return ingredient, nil
}

func (app *Application) UpdateIngredient(ingredient Ingredient) error {
	return nil
}

func (app *Application) InsertIngredient(ingredient Ingredient) (int64, error) {
	return 0, nil
}

func (app *Application) DeleteIngredient(ingredient Ingredient) error {
	return nil
}

func (app *Application) GetAlternatives(ingredientID int64) ([]int64, error) {
	return []int64{}, nil
}

func (app *Application) AddAlternative(ingredient int64, alternative int64) (int64, error) {
	return 0, nil
}

func (app *Application) DeleteAlternative(ingredient int64, alternative int64) error {
	return nil
}
