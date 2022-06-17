package repository

import "github.com/MelkdeSousa/ghopping/domain"

type ProductRepository interface {
	GetAll() ([]domain.Product, error)
	GetByID(id string) (domain.Product, error)
	Insert(
		id string,
		name, description string,
		price int,
		quantity int,
	) error
	DeleteById(id string) error
	UpdateById(
		id, name, description string,
		price, quantity int,
	) error
}
