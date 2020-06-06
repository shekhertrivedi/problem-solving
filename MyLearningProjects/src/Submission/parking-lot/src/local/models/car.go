package models

type Car struct {
	Color  string `json:"color"`
	Number string `json:"number"`
}

func (c Car) GetNumber() string {
	return c.Number
}

func (c Car) GetColor() string {
	return c.Color
}
