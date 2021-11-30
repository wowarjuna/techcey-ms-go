package booking

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type listCargosRequest struct {}

type listCargosResponse struct {
	Cargos [Cargo]
}

func makeListCargosEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(listCargoRequest)
		return list
	}
}