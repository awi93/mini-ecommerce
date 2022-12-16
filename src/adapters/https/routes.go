package https

import (
	"net/http"

	"github.com/awi93/mini-ecommerce/src/adapters/https/handlers/order"
	"github.com/awi93/mini-ecommerce/src/utils/httputil"
	"github.com/go-chi/chi"
)

func Router() *chi.Mux {
	r := chi.NewRouter()

	oh := order.NewHandler(&order.Args{
		Service: OrderService,
	})

	r.Get("/order-details/{id}", func(w http.ResponseWriter, r *http.Request) {
		httputil.HandleRequest(w, r, oh.Find)
	})

	return r
}
