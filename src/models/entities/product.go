package entities

type Product struct {
	Id          int64   `json:"id" gorm:"id"`
	ProductName string  `json:"product_name" gorm:"product_name"`
	UnitPrice   float64 `json:"unit_price" gorm:"unit_price"`
	InStock     bool    `json:"in_stock" gorm:"in_stock"`
}
