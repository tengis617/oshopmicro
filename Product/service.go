package Product

import (
	"context"
	"errors"
	"time"
)

// Service is a CRUD interface for products
type Service interface {
	GetProducts(ctx context.Context) ([]Product, error)
	GetProduct(ctx context.Context, id string) (Product, error)
	PostProduct(ctx context.Context, p Product) error
	PutProduct(ctx context.Context, id string, p Product) error
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

var (
	ErrAlreadyExists  = errors.New("product already exists")
	ErrInconsistentID = errors.New("inconstistent product IDs")
	ErrNotFound       = errors.New("Product not found")
)
