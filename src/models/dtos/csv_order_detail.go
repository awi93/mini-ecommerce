package dtos

type CsvOrderDetail struct {
	Id        int64
	OrderId   int64
	ProductId int64
	Quantity  int
	UnitPrice float64
	Discount  float64
}
