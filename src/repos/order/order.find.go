package order

import (
	"errors"
	"fmt"

	"github.com/awi93/mini-ecommerce/src/models/dtos"
)

var orderQuery = `
	select
		o.id,
		o.customer_id,
		c.first_name as customer_first_name,
		c.last_name as customer_last_name,
		o.employee_id,
		e.first_name as employee_first_name,
		e.last_name as employee_last_name,
		o.order_date,
		o.purchase_order_number,
		o.ship_date,
		o.shipping_method_id,
		sm.shipping_method,
		o.freight_charge,
		o.taxes,
		o.payment_received,
		o."comment" 
	from
		orders o
		inner join
		customers c
		on o.customer_id = c.id
		inner join
		employees e
		on o.employee_id = e.id
		inner join
		shipping_methods sm
		on o.shipping_method_id = sm.id
	where
		o.id = ?
`
var orderDetailQuery = `
	select
		od.id,
		od.order_id,
		od.product_id,
		p.product_name,
		od.quantity,
		od.unit_price,
		od.discount
	from
		order_details od
		inner join 
		products p
		on od.product_id = p.id
	where
		od.order_id = ?
`

func (r *repo) FindOrderDetailById(id int64) (*dtos.Order, error) {
	var order *dtos.OrderQuery
	result := r.DB.Raw(orderQuery, id).Scan(&order)
	if result.Error != nil {
		return nil, result.Error
	}
	if order == nil {
		return nil, errors.New("record not found")
	}
	var orderDetails []*dtos.OrderDetail
	result = r.DB.Raw(orderDetailQuery, id).Scan(&orderDetails)
	if result.Error != nil {
		return nil, result.Error
	}

	var subtotal = 0.0
	for _, orderDetail := range orderDetails {
		sub := float64(orderDetail.Quantity) * (orderDetail.UnitPrice * (100.0 - orderDetail.Discount) / 100)
		orderDetail.Subtotal = sub
		subtotal += sub
	}

	var customerName = order.CustomerFirstName
	if order.CustomerLastName != nil {
		customerName = fmt.Sprintf("%s %s", customerName, *order.CustomerLastName)
	}

	var employeeName = order.EmployeeFirstName
	if order.EmployeeLastName != nil {
		employeeName = fmt.Sprintf("%s %s", employeeName, *order.EmployeeLastName)
	}

	return &dtos.Order{
		Id:                  order.Id,
		CustomerId:          order.CustomerId,
		Customer:            customerName,
		EmployeeId:          order.EmployeeId,
		Employee:            employeeName,
		OrderDate:           order.OrderDate,
		PurchaseOrderNumber: order.PurchaseOrderNumber,
		ShipDate:            order.ShipDate,
		ShippingMethodId:    order.ShippingMethodId,
		ShippingMethod:      order.ShippingMethod,
		FreightCharge:       order.FreightCharge,
		Taxes:               order.Taxes,
		Subtotal:            subtotal,
		GrandTotal:          subtotal + order.FreightCharge + order.Taxes,
		Comment:             order.Comment,
		OrderDetails:        orderDetails,
	}, nil
}
