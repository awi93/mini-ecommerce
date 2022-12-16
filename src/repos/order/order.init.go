package order

import (
	"encoding/csv"

	"github.com/awi93/mini-ecommerce/src/models/dtos"
	"github.com/awi93/mini-ecommerce/src/models/entities"
	"gorm.io/gorm"
)

type Repo interface {
	FindOrderDetailById(id int64) (*dtos.Order, error)
	ReadOrderCsv(reader *csv.Reader) ([]*dtos.CsvOrder, error)
	CreateOrderFromCsv(data *dtos.CsvOrder) (*entities.Order, error)
	ReadOrderItemCsv(reader *csv.Reader) ([]*dtos.CsvOrderDetail, error)
	CreateOrderDetailFromCsv(data *dtos.CsvOrderDetail) (*entities.OrderDetail, error)
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
