package employee

import (
	"encoding/csv"

	"github.com/awi93/mini-ecommerce/src/models/dtos"
	"github.com/awi93/mini-ecommerce/src/models/entities"
	"gorm.io/gorm"
)

type Repo interface {
	ReadCsv(reader *csv.Reader) ([]*dtos.CsvEmployee, error)
	CreateFromCsv(data *dtos.CsvEmployee) (*entities.Employee, error)
	GetDB() *gorm.DB
}

type repo struct {
	DB *gorm.DB
}

type Args struct {
	DB *gorm.DB
}

func NewRepo(args *Args) Repo {
	return &repo{
		DB: args.DB,
	}
}

func (r *repo) GetDB() *gorm.DB {
	return r.DB
}
