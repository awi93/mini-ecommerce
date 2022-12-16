package order

import "github.com/awi93/mini-ecommerce/src/models/dtos"

func (o *service) FindOrderDetailById(id int64) (*dtos.Order, error) {
	return o.Repo.FindOrderDetailById(id)
}
