package customer

import (
	"encoding/csv"
	"errors"
	"strconv"

	"github.com/awi93/mini-ecommerce/src/models/dtos"
)

func (s *repo) ReadCsv(reader *csv.Reader) ([]*dtos.CsvCustomer, error) {
	var results = make([]*dtos.CsvCustomer, 0)
	var i = 0
	for {
		row, err := reader.Read()
		if err != nil {
			break
		}
		if len(row) < 17 {
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
		var csvRow = dtos.CsvCustomer{
			Id:        id,
			FirstName: row[1],
		}
		if row[2] != "" {
			csvRow.LastName = &row[2]
		}
		if row[3] != "" {
			csvRow.Email = &row[3]
		}
		if row[4] != "" {
			csvRow.PhoneNumber = &row[4]
		}
		if row[5] != "" {
			csvRow.FaxNumber = &row[5]
		}
		if row[6] != "" {
			csvRow.BillingAddress = &row[6]
		}
		if row[7] != "" {
			csvRow.BillingCity = &row[7]
		}
		if row[8] != "" {
			csvRow.BillingState = &row[8]
		}
		if row[9] != "" {
			csvRow.BillingZipCode = &row[9]
		}
		if row[10] != "" {
			csvRow.CompanyName = &row[10]
		}
		if row[11] != "" {
			csvRow.CompanyWebsite = &row[11]
		}
		if row[12] != "" {
			csvRow.ShipAddress = &row[12]
		}
		if row[13] != "" {
			csvRow.ShipCity = &row[13]
		}
		if row[14] != "" {
			csvRow.ShipState = &row[14]
		}
		if row[15] != "" {
			csvRow.ShipZipCode = &row[15]
		}
		if row[16] != "" {
			csvRow.ShipPhoneNumber = &row[16]
		}
		results = append(results, &csvRow)
	}
	return results, nil
}
