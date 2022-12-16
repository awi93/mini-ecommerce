package employee

import (
	"github.com/awi93/mini-ecommerce/src/models/dtos"
	"github.com/awi93/mini-ecommerce/src/models/entities"
)

func (r *repo) composeEmployee(data *dtos.CsvEmployee) *entities.Employee {
	return &entities.Employee{
		Id:        data.Id,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Title:     data.Title,
		WorkPhone: data.WorkPhone,
	}
}

func (r *repo) CreateFromCsv(data *dtos.CsvEmployee) (*entities.Employee, error) {
	employee := r.composeEmployee(data)
	result := r.DB.Create(&employee)
	return employee, result.Error
}
