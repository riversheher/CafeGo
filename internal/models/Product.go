package models

type Product struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Price       float64      `json:"price"`
	Ingredients []Ingredient `json:"ingredients"`
	Discount    float64      `json:"discount"`
	Type        string       `json:"type"`
}
