package order

import (
	"github.com/awi93/mini-ecommerce/src/models/dtos"
	"github.com/awi93/mini-ecommerce/src/models/entities"
)

func (r *repo) composeOrder(data *dtos.CsvOrder) *entities.Order {
	return &entities.Order{
		Id:                  data.Id,
		CustomerId:          data.CustomerId,
		EmployeeId:          data.EmployeeId,
		OrderDate:           data.OrderDate,
		PurchaseOrderNumber: data.PurchaseOrderNumber,
		ShipDate:            data.ShipDate,
		ShippingMethodId:    data.ShippingMethodId,
		FreightCharge:       data.FreightCharge,
		Taxes:               data.Taxes,
		PaymentReceived:     data.PaymentReceived,
		Comment:             data.Comment,
	}
}

func (r *repo) CreateOrderFromCsv(data *dtos.CsvOrder) (*entities.Order, error) {
	order := r.composeOrder(data)
	result := r.DB.Create(&order)
	return order, result.Error
}
