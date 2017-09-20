package Product

import (
	"context"

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

func MakeServerEnpoints(s Service) Endpoints {
	return Endpoints{
		GetProductEndpoint:  MakeGetProductEndpoint(s),
		PostProductEndpoint: MakePostProductEndpoint(s),
	}
}

func MakeGetProductEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getProductRequest)
		p, err := s.GetProduct(ctx, req.ID)
		return getProductResponse{Product: p, Err: err}, nil
	}
}
func MakeCreateProductEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request)
}

type getProductRequest struct {
	ID string
}
type getProductResponse struct {
	Product Product `json:"product,omitempty"`
	Err     error   `json:"err,omitempty"`
}

type createProductRequest struct {
	Product Product
}
type createProductResponse struct {
	Err error `json:"err,omitempty"`
}
