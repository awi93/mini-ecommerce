package shippingmethod

import (
	"encoding/csv"
	"errors"
	"strconv"

	"github.com/awi93/mini-ecommerce/src/models/dtos"
)

func (r *repo) ReadCsv(reader *csv.Reader) ([]*dtos.CsvShippingMethod, error) {
	var results = make([]*dtos.CsvShippingMethod, 0)
	var i = 0
	for {
		row, err := reader.Read()
		if err != nil {
			break
		}
		if len(row) < 2 {
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
		var csvRow = dtos.CsvShippingMethod{
			Id:             id,
			ShippingMethod: row[1],
		}
		results = append(results, &csvRow)
	}
	return results, nil
}
