package booking

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	kitlog "github.com/go-kit/log"
	"github.com/gorilla/mux"
)

func MakeHandler(bs Service, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	listCargosHandler := kithttp.NewServer(makeListCargosEndpoint())
	r := mux.NewRouter()

	r.Handle("/booking/v1/cargos", listCargosHandler).Methods("GET")

	return r
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {

}
