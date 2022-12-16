package order

import (
	"encoding/csv"
	"errors"
	"log"

	"github.com/awi93/mini-ecommerce/src/models/dtos"
)

func (s *service) ImportOrder(reader *csv.Reader) ([]*dtos.ImportResult, error) {
	csvRecords, err := s.Repo.ReadOrderCsv(reader)
	if err != nil {
		log.Println(err)
		return nil, errors.New("fail to read csv file")
	}

	var results = make([]*dtos.ImportResult, len(csvRecords))
	for i, csvRecord := range csvRecords {
		errors := s.validateOrderCsv(csvRecord)
		if len(errors) > 0 {
			results[i] = &dtos.ImportResult{
				Index:  i,
				Status: "FAILED",
				Errors: errors,
			}
			continue
		}
		_, err := s.Repo.CreateOrderFromCsv(csvRecord)
		if err != nil {
			results[i] = &dtos.ImportResult{
				Index:  i,
				Status: "FAILED",
				Errors: []string{
					err.Error(),
				},
			}
			continue
		}
		results[i] = &dtos.ImportResult{
			Index:  i,
			Status: "SUCCEED",
		}
	}

	return results, nil
}

func (s *service) validateOrderCsv(record *dtos.CsvOrder) []string {
	var errors []string

	if record.Id == 0 {
		errors = append(errors, "id can't be empty")
	}
	if record.CustomerId == 0 {
		errors = append(errors, "customer id can't be empty")
	}
	if record.EmployeeId == 0 {
		errors = append(errors, "employee id can't be empty")
	}
	if record.ShippingMethodId == 0 {
		errors = append(errors, "shipping method id can't be empty")
	}

	return errors
}
