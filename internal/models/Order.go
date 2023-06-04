package models

import (
	"time"
)

type Order struct {
	ID       int64     `json:"id"`
	Customer User      `json:"customer"`
	Products []Product `json:"products"`
	Time     time.Time `json:"time"`
	Subtotal float64   `json:"subtotal"`
}
