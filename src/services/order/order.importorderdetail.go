package order

import (
	"encoding/csv"
	"errors"
	"log"

	"github.com/awi93/mini-ecommerce/src/models/dtos"
)

func (s *service) ImportOrderDetail(reader *csv.Reader) ([]*dtos.ImportResult, error) {
	csvRecords, err := s.Repo.ReadOrderItemCsv(reader)
	if err != nil {
		log.Println(err)
		return nil, errors.New("fail to read csv file")
	}

	var results = make([]*dtos.ImportResult, len(csvRecords))
	for i, csvRecord := range csvRecords {
		errors := s.validateOrderDetailCsv(csvRecord)
		if len(errors) > 0 {
			results[i] = &dtos.ImportResult{
				Index:  i,
				Status: "FAILED",
				Errors: errors,
			}
			continue
		}
		_, err := s.Repo.CreateOrderDetailFromCsv(csvRecord)
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

func (s *service) validateOrderDetailCsv(record *dtos.CsvOrderDetail) []string {
	var errors []string

	if record.Id == 0 {
		errors = append(errors, "id can't be empty")
	}
	if record.OrderId == 0 {
		errors = append(errors, "order id can't be empty")
	}
	if record.ProductId == 0 {
		errors = append(errors, "product id can't be empty")
	}

	return errors
}
