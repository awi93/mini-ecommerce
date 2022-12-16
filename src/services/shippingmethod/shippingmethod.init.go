package shippingmethod

import (
	"encoding/csv"

	"github.com/awi93/mini-ecommerce/src/models/dtos"
	"github.com/awi93/mini-ecommerce/src/repos/shippingmethod"
)

type Service interface {
	Import(reader *csv.Reader) ([]*dtos.ImportResult, error)
}

type service struct {
	Repo shippingmethod.Repo
}

type Args struct {
	Repo shippingmethod.Repo
}

func NewService(args *Args) Service {
	return &service{
		Repo: args.Repo,
	}
}
