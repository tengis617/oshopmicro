package Product

import (
	"context"
	"time"

	"github.com/jinzhu/gorm"
)

// Service is a CRUD interface for products
type Service interface {
	GetProducts(ctx context.Context) ([]Product, error)
	GetProduct(ctx context.Context, id string) (Product, error)
	CreateProduct(ctx context.Context, p Product) error
	UpdateProduct(ctx context.Context, id string, p Product) error
	DeleteProduct(ctx context.Context, id string) error
}

// Product is a simple struct
type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	// i could create a seperate struct for this
	Manufacturer string    `json:"manufacturer"`
	DateAdded    time.Time `json:"dateAdded"`
}

type service struct {
	db *gorm.DB
}

func (s *service) GetProducts(ctx context.Context) ([]Product, error) {
	var products []Product

	if err := s.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (s *service) GetProduct(ctx context.Context, id string) (Product, error) {
	var p Product
	if err := s.db.First(&p, id).Error; err != nil {
		return Product{}, err
	}
	return p, nil
}

func (s *service) CreateProduct(ctx context.Context, p Product) error {
	if err := s.db.Create(&p).Error; err != nil {
		return err
	}
	return nil
}

func (s *service) UpdateProduct(ctx context.Context, id string, p Product) error {
	if err := s.db.Model(&p).Where("ID = ?", id).Update(p).Error; err != nil {
		return err
	}
	return nil
}
func (s *service) DeleteProduct(ctx context.Context, id string) error {
	if err := s.db.Delete(Product{}, "ID = ?", id).Error; err != nil {
		return err
	}

	return nil
}
