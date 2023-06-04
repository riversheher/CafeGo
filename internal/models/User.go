package models

type User struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Rewards int64  `json:"rewards"`
}
