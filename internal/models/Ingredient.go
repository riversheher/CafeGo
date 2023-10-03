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
		return ingredient, err
	}

	ingredient.Alternatives, err = app.GetAlternatives(id)
	if err != nil {
		app.ErrLog.Println(err)
		return ingredient, err
	}

	return ingredient, nil
}

func (app *Application) UpdateIngredient(ingredient Ingredient) error {
	query := fmt.Sprintf("UPDATE %s SET name = ?, description = ?, price = ?, amount = ?, type = ? WHERE id = ?", IngredientTable)
	_, err := app.DB.Exec(query, ingredient.Name, ingredient.Description, ingredient.Price, ingredient.Amount, ingredient.Type, ingredient.ID)
	if err != nil {
		app.logError(err)
		return err
	}
	return nil
}

func (app *Application) InsertIngredient(ingredient Ingredient) (int64, error) {
	query := fmt.Sprintf("INSERT INTO %s (name, description, price, amount, type) VALUES (?, ?, ?, ?, ?)", IngredientTable)
	res, err := app.DB.Exec(query, ingredient.Name, ingredient.Description, ingredient.Price, ingredient.Amount, ingredient.Type)
	if err != nil {
		return 0, err
	} else {
		ID, err := res.LastInsertId()
		if err != nil {
			app.logError(err)
			return 0, err
		}

		for _, alternative := range ingredient.Alternatives {
			_, err := app.AddAlternative(ID, alternative)
			if err != nil {
				app.logError(err)
				return ID, err
			}
		}

		return ID, err
	}
}

func (app *Application) DeleteIngredient(ingredient Ingredient) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", IngredientTable)
	_, err := app.DB.Exec(query, ingredient.ID)
	if err != nil {
		app.logError(err)
		return err
	}
	return nil
}

func (app *Application) GetAlternatives(ingredientID int64) ([]int64, error) {
	var alternatives []int64

	query := fmt.Sprintf("SELECT alternative_id FROM %s WHERE ingredient_id = ?", AlternativesTable)
	rows, err := app.DB.Query(query, ingredientID)
	if err != nil {
		app.logError(err)
		return alternatives, err
	}
	defer rows.Close()

	for rows.Next() {
		var alternative int64
		err = rows.Scan(&alternative)
		if err != nil {
			app.logError(err)
			return alternatives, err
		}
		alternatives = append(alternatives, alternative)
	}

	return alternatives, nil
}

func (app *Application) AddAlternative(ingredient int64, alternative int64) (int64, error) {
	query := fmt.Sprintf("INSERT INTO %s (ingredient_id, alternative_id) VALUES (?, ?)", AlternativesTable)
	res, err := app.DB.Exec(query, ingredient, alternative)
	if err != nil {
		app.logError(err)
		return 0, err
	}
	ID, err := res.LastInsertId()
	if err != nil {
		app.logError(err)
		return 0, err
	}
	return ID, nil
}

func (app *Application) DeleteAlternative(ingredient int64, alternative int64) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE ingredient_id = ? AND alternative_id = ?", AlternativesTable)
	_, err := app.DB.Exec(query, ingredient, alternative)
	if err != nil {
		app.logError(err)
		return err
	}
	return nil
}
