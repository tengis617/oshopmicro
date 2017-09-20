package Product

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

var (
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)

func MakeHTTPHandler(s Service) http.Handler {
	r := mux.NewRouter()
	e := MakeServerEnpoints(s)

	r.Methods("GET").Path("/products/{id}").Handler(httptransport.NewServer(
		e.GetProductEndpoint,
		decodeGetProductRequest,
		encodeResponse,
	))

	r.Methods("POST").Path("/products/").Handler(httptransport.NewServer(
		e.CreateProductEndpoint,
		decodeCreateProductRequest,
		encodeResponse,
	))
	return r
}

func decodeGetProductRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}

	return getProductRequest{ID: id}, nil
}

func decodeCreateProductRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req createProductRequest
	if e := json.NewDecoder(r.Body).Decode(&req.Product); e != nil {
		return nil, e
	}
	return req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
