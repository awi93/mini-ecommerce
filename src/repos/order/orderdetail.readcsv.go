package order

import (
	"encoding/csv"
	"errors"
	"strconv"

	"github.com/awi93/mini-ecommerce/src/models/dtos"
)

func (r *repo) ReadOrderItemCsv(reader *csv.Reader) ([]*dtos.CsvOrderDetail, error) {
	var results = make([]*dtos.CsvOrderDetail, 0)
	var i = 0
	for {
		row, err := reader.Read()
		if err != nil {
			break
		}
		if len(row) < 4 {
			return nil, errors.New("invalid csv file")
		}
		if i == 0 {
			i++
			continue
		}
		id, err := strconv.ParseInt(row[0], 10, 64)
		if err != nil {
			return nil, errors.New("invalid id value")
		}
		orderId, err := strconv.ParseInt(row[1], 10, 64)
		if err != nil {
			return nil, errors.New("invalid order id value")
		}
		productId, err := strconv.ParseInt(row[2], 10, 64)
		if err != nil {
			return nil, errors.New("invalid product id value")
		}
		quantity, err := strconv.ParseInt(row[3], 10, 16)
		if err != nil {
			return nil, errors.New("invalid quantity value")
		}
		unitPrice, err := strconv.ParseFloat(row[4], 64)
		if err != nil {
			return nil, errors.New("invalid unit price value")
		}
		discount, err := strconv.ParseFloat(row[5], 64)
		if err != nil {
			return nil, errors.New("invalid discount value")
		}
		var csvRow = dtos.CsvOrderDetail{
			Id:        id,
			OrderId:   orderId,
			ProductId: productId,
			Quantity:  int(quantity),
			UnitPrice: unitPrice,
			Discount:  discount,
		}

		results = append(results, &csvRow)
	}
	return results, nil
}
