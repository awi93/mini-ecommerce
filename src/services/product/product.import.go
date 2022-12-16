package product

import (
	"encoding/csv"
	"errors"
	"log"

	"github.com/awi93/mini-ecommerce/src/models/dtos"
)

func (s *service) Import(reader *csv.Reader) ([]*dtos.ImportResult, error) {
	csvRecords, err := s.Repo.ReadCsv(reader)
	if err != nil {
		log.Println(err)
		return nil, errors.New("fail to read csv file")
	}

	var results = make([]*dtos.ImportResult, len(csvRecords))
	for i, csvRecord := range csvRecords {
		errors := s.validateCsv(csvRecord)
		if len(errors) > 0 {
			results[i] = &dtos.ImportResult{
				Index:  i,
				Status: "FAILED",
				Errors: errors,
			}
			continue
		}
		_, err := s.Repo.CreateFromCsv(csvRecord)
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

func (s *service) validateCsv(record *dtos.CsvProduct) []string {
	var errors []string
	if record.Id == 0 {
		errors = append(errors, "id can't be empty")
	}

	if record.ProductName == "" {
		errors = append(errors, "shipping method can't be empty")
	}

	return errors
}
