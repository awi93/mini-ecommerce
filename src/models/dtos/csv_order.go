package dtos

import "time"

type CsvOrder struct {
	Id                  int64     `json:"id" gorm:"id"`
	CustomerId          int64     `json:"customer_id" gorm:"customer_id"`
	EmployeeId          int64     `json:"employee_id" gorm:"employee_id"`
	OrderDate           time.Time `json:"order_date" gorm:"order_date"`
	PurchaseOrderNumber string    `json:"purchase_order_number" gorm:"purchase_order_number"`
	ShipDate            time.Time `json:"ship_date" gorm:"ship_date"`
	ShippingMethodId    int64     `json:"shipping_method_id" gorm:"shipping_method_id"`
	FreightCharge       float64   `json:"freight_charge" gorm:"freight_charge"`
	Taxes               float64   `json:"taxes" gorm:"taxes"`
	PaymentReceived     bool      `json:"payment_received" gorm:"payment_received"`
	Comment             *string   `json:"comment" gorm:"comment"`
}
