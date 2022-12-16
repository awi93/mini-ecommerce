package order

import (
	"github.com/awi93/mini-ecommerce/src/models/dtos"
	"github.com/awi93/mini-ecommerce/src/models/entities"
)

func (r *repo) composeOrderDetail(data *dtos.CsvOrderDetail) *entities.OrderDetail {
	return &entities.OrderDetail{
		Id:        data.Id,
		OrderId:   data.OrderId,
		ProductId: data.ProductId,
		Quantity:  data.Quantity,
		UnitPrice: data.UnitPrice,
		Discount:  data.Discount,
	}
}

func (r *repo) CreateOrderDetailFromCsv(data *dtos.CsvOrderDetail) (*entities.OrderDetail, error) {
	orderDetail := r.composeOrderDetail(data)
	result := r.DB.Create(&orderDetail)
	return orderDetail, result.Error
}
