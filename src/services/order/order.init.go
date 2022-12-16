package order

import (
	"encoding/csv"

	"github.com/awi93/mini-ecommerce/src/models/dtos"
	"github.com/awi93/mini-ecommerce/src/repos/order"
)

type Service interface {
	ImportOrder(reader *csv.Reader) ([]*dtos.ImportResult, error)
	ImportOrderDetail(reader *csv.Reader) ([]*dtos.ImportResult, error)
	FindOrderDetailById(id int64) (*dtos.Order, error)
}

type service struct {
	Repo order.Repo
}

type Args struct {
	Repo order.Repo
}

func NewService(args *Args) Service {
	return &service{
		Repo: args.Repo,
	}
}
