package ordersvc

import (
	"context"
	"time"

	"github.com/jinzhu/gorm"
)

// OrderService is a CRUD interface for handling orders
type OrderService interface {
	CreateOrder(ctx context.Context, item string) error
	GetOrder(ctx context.Context, id string) (Order, error)
	GetOrders(ctx context.Context) ([]Order, error)
	DeleteOrder(ctx context.Context, id string) error
}

// Order represents an order made by the client
// for now Item will just be a string
type Order struct {
	ID        string    `json:"id" gorm:"primary_key"`
	Item      string    `json:"item"`
	CreatedOn time.Time `json:"createdOn"`
	Status    string    `json:"status"`
}

type orderService struct {
	db *gorm.DB
}

// NewOrderService is a order service constructor
func NewOrderService(db *gorm.DB) OrderService {
	return &orderService{db: db}
}

func (s *orderService) CreateOrder(ctx context.Context, item string) error {
	o := Order{
		Item:      item,
		CreatedOn: time.Now(),
		Status:    "Pending",
	}
	if err := s.db.Create(&o).Error; err != nil {
		return err
	}
	return nil
}

func (s *orderService) GetOrder(ctx context.Context, id string) (Order, error) {
	var order Order
	if err := s.db.First(&order, id).Error; err != nil {
		return Order{}, err
	}
	return order, nil
}
func (s *orderService) GetOrders(ctx context.Context) ([]Order, error) {
	var orders []Order
	if err := s.db.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}
func (s *orderService) DeleteOrder(ctx context.Context, id string) error {
	if err := s.db.Delete(Order{}, "ID = ?", id).Error; err != nil {
		return err
	}
	return nil
}
