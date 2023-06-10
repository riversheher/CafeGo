package models

type Menu struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Products    []Product `json:"products"`
}
