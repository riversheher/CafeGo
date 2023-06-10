package models

import (
	"time"
)

type Order struct {
	ID       string    `json:"id"`
	Customer User      `json:"customer"`
	Products []Product `json:"products"`
	Time     time.Time `json:"time"`
	Subtotal float64   `json:"subtotal"`
}
