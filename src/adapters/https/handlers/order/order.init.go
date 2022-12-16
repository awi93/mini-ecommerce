package order

import orderSvc "github.com/awi93/mini-ecommerce/src/services/order"

type Handler struct {
	service orderSvc.Service
}

type Args struct {
	Service orderSvc.Service
}

func NewHandler(Args *Args) *Handler {
	return &Handler{
		service: Args.Service,
	}
}
