package shippingmethod

import (
	"github.com/awi93/mini-ecommerce/src/models/dtos"
	"github.com/awi93/mini-ecommerce/src/models/entities"
)

func (r *repo) composeShippingMethod(data *dtos.CsvShippingMethod) *entities.ShippingMethod {
	return &entities.ShippingMethod{
		Id:             data.Id,
		ShippingMethod: data.ShippingMethod,
	}
}

func (r *repo) CreateFromCsv(data *dtos.CsvShippingMethod) (*entities.ShippingMethod, error) {
	shippingmethod := r.composeShippingMethod(data)
	result := r.DB.Create(&shippingmethod)
	return shippingmethod, result.Error
}
