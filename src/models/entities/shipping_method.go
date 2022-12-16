package entities

type ShippingMethod struct {
	Id             int64  `json:"id" gorm:"id"`
	ShippingMethod string `json:"shipping_method" gorm:"shipping_method"`
}
