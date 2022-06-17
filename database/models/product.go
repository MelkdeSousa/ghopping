package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	ID          string
	Name        string
	Description string
	Price       int
	Quantity    int
}
