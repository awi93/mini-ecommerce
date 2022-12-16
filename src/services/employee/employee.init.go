package employee

import (
	"encoding/csv"

	"github.com/awi93/mini-ecommerce/src/models/dtos"
	"github.com/awi93/mini-ecommerce/src/repos/employee"
)

type Service interface {
	Import(reader *csv.Reader) ([]*dtos.ImportResult, error)
}

type service struct {
	Repo employee.Repo
}

type Args struct {
	Repo employee.Repo
}

func NewService(args *Args) Service {
	return &service{
		Repo: args.Repo,
	}
}
