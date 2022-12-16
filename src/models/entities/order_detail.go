package entities

type OrderDetail struct {
	Id        int64   `json:"id" gorm:"id"`
	OrderId   int64   `json:"order_id" gorm:"order_id"`
	ProductId int64   `json:"product_id" gorm:"product_id"`
	Quantity  int     `json:"quantity" gorm:"quantity"`
	UnitPrice float64 `json:"unit_price" gorm:"unit_price"`
	Discount  float64 `json:"discount" gorm:"discount"`
}
