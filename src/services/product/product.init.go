package product

import (
	"encoding/csv"

	"github.com/awi93/mini-ecommerce/src/models/dtos"
	"github.com/awi93/mini-ecommerce/src/repos/product"
)

type Service interface {
	Import(reader *csv.Reader) ([]*dtos.ImportResult, error)
}

type service struct {
	Repo product.Repo
}

type Args struct {
	Repo product.Repo
}

func NewService(args *Args) Service {
	return &service{
		Repo: args.Repo,
	}
}
