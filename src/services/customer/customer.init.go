package customer

import (
	"encoding/csv"

	"github.com/awi93/mini-ecommerce/src/models/dtos"
	"github.com/awi93/mini-ecommerce/src/repos/customer"
)

type Service interface {
	Import(reader *csv.Reader) ([]*dtos.ImportResult, error)
}

type service struct {
	Repo customer.Repo
}

type Args struct {
	Repo customer.Repo
}

func NewService(args *Args) Service {
	return &service{
		Repo: args.Repo,
	}
}
