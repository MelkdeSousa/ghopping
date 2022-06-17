package database

import (
	"path/filepath"

	"github.com/MelkdeSousa/ghopping/database/models"
	"github.com/MelkdeSousa/ghopping/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository() *ProductRepository {
	db, err := gorm.Open(sqlite.Open(filepath.Join("database.db")), &gorm.Config{})

	db.AutoMigrate(&models.Product{})

	if err != nil {
		panic(err)
	}

	return &ProductRepository{db: db}
}

func (p *ProductRepository) GetAll() ([]domain.Product, error) {
	var products []domain.Product
	p.db.Table("products").Select("id, name, description, price, quantity").Where("deleted_at IS NULL").Order("created_at desc").Scan(&products)
	return products, nil
}

func (p *ProductRepository) GetById(id string) (domain.Product, error) {
	var product domain.Product
	err := p.db.Where("id = ?", id).First(&product).Error

	return product, err
}

func (p *ProductRepository) Insert(id, name, description string, price int, quantity int) error {
	product := models.Product{
		ID:          id,
		Name:        name,
		Description: description,
		Price:       price,
		Quantity:    quantity,
	}

	return p.db.Create(&product).Error
}

func (p *ProductRepository) DeleteById(id string) error {
	return p.db.Table("products").Where("id = ?", id).Delete(&models.Product{}).Error
}

func (p *ProductRepository) UpdateById(id, name, description string, price, quantity int) error {
	product, err := p.GetById(id)

	if err != nil {
		return err
	}

	product.Name = name
	product.Description = description
	product.Price = price
	product.Quantity = quantity

	return p.db.Save(&product).Error
}
