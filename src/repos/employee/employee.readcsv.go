package employee

import (
	"encoding/csv"
	"errors"
	"strconv"

	"github.com/awi93/mini-ecommerce/src/models/dtos"
)

func (r *repo) ReadCsv(reader *csv.Reader) ([]*dtos.CsvEmployee, error) {
	var results = make([]*dtos.CsvEmployee, 0)
	var i = 0
	for {
		row, err := reader.Read()
		if err != nil {
			break
		}
		if len(row) < 5 {
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
		var csvRow = dtos.CsvEmployee{
			Id:        id,
			FirstName: row[1],
			Title:     row[3],
		}
		if row[2] != "" {
			csvRow.LastName = &row[2]
		}
		if row[4] != "" {
			csvRow.WorkPhone = &row[4]
		}

		results = append(results, &csvRow)
	}
	return results, nil
}
