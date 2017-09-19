package Product

import (
	"github.com/go-kit/kit/endpoint"
)

// Endpoints is a collection of endpoints for product service
type Endpoints struct {
	GetProductsEndpoint   endpoint.Endpoint
	GetProductEndpoint    endpoint.Endpoint
	CreateProductEndpoint endpoint.Endpoint
	UpdateProductEndpoint endpoint.Endpoint
	DeleteProductEndpoint endpoint.Endpoint
}
