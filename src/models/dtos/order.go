package dtos

import "time"

type OrderQuery struct {
	Id                  int64     `gorm:"id"`
	CustomerId          int64     `gorm:"customer_id"`
	CustomerFirstName   string    `gorm:"customer_first_name"`
	CustomerLastName    *string   `json:"customer_last_name"`
	EmployeeId          int64     `gorm:"employee_id"`
	EmployeeFirstName   string    `json:"employee_first_name"`
	EmployeeLastName    *string   `json:"employee_last_name"`
	OrderDate           time.Time `gorm:"order_date"`
	PurchaseOrderNumber string    `gorm:"purchase_order_number"`
	ShipDate            time.Time `gorm:"ship_date"`
	ShippingMethodId    int64     `gorm:"shipping_method_id"`
	ShippingMethod      string    `gorm:"shipping_method"`
	FreightCharge       float64   `gorm:"freight_charge"`
	Taxes               float64   `gorm:"taxes"`
	PaymentReceived     bool      `gorm:"payment_received"`
	Comment             *string   `gorm:"comment"`
}

type Order struct {
	Id                  int64          `json:"id"`
	CustomerId          int64          `json:"customer_id"`
	Customer            string         `json:"customer"`
	EmployeeId          int64          `json:"employee_id"`
	Employee            string         `json:"employee"`
	OrderDate           time.Time      `json:"order_date"`
	PurchaseOrderNumber string         `json:"purchase_order_number"`
	ShipDate            time.Time      `json:"ship_date"`
	ShippingMethodId    int64          `json:"shipping_method_id"`
	ShippingMethod      string         `json:"shipping_method"`
	FreightCharge       float64        `json:"freight_charge"`
	Taxes               float64        `json:"taxes"`
	Subtotal            float64        `json:"subtotal"`
	GrandTotal          float64        `json:"grand_total"`
	PaymentReceived     bool           `json:"payment_received"`
	Comment             *string        `json:"comment"`
	OrderDetails        []*OrderDetail `json:"order_details"`
}

type OrderDetail struct {
	Id          int64   `json:"id" gorm:"id"`
	OrderId     int64   `json:"order_id" gorm:"order_id"`
	ProductId   int64   `json:"product_id" gorm:"product_id"`
	ProductName string  `json:"product_name" gorm:"product_name"`
	Quantity    int     `json:"quantity" gorm:"quantity"`
	UnitPrice   float64 `json:"unit_price" gorm:"unit_price"`
	Discount    float64 `json:"discount" gorm:"discount"`
	Subtotal    float64 `json:"subtoal" gorm:"subtotal"`
}
