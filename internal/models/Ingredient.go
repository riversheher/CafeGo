package models

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
