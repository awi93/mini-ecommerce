package product

import (
	"encoding/csv"
	"errors"
	"strconv"

	"github.com/awi93/mini-ecommerce/src/models/dtos"
)

func (r *repo) ReadCsv(reader *csv.Reader) ([]*dtos.CsvProduct, error) {
	var results = make([]*dtos.CsvProduct, 0)
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

		unitPrice, err := strconv.ParseFloat(row[2], 64)
		if err != nil {
			return nil, errors.New("invalid unit price value")
		}

		inStock, err := strconv.ParseBool(row[3])
		if err != nil {
			return nil, errors.New("invalid in stock value")
		}

		var csvRow = dtos.CsvProduct{
			Id:          id,
			ProductName: row[1],
			UnitPrice:   unitPrice,
			InStock:     inStock,
		}

		results = append(results, &csvRow)
	}
	return results, nil
}
