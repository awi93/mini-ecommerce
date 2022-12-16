package product

import (
	"github.com/awi93/mini-ecommerce/src/models/dtos"
	"github.com/awi93/mini-ecommerce/src/models/entities"
)

func (r *repo) composeProduct(data *dtos.CsvProduct) *entities.Product {
	return &entities.Product{
		Id:          data.Id,
		ProductName: data.ProductName,
		UnitPrice:   data.UnitPrice,
		InStock:     data.InStock,
	}
}

func (r *repo) CreateFromCsv(data *dtos.CsvProduct) (*entities.Product, error) {
	product := r.composeProduct(data)
	result := r.DB.Create(&product)
	return product, result.Error
}
