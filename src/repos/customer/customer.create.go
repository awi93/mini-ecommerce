package customer

import (
	"github.com/awi93/mini-ecommerce/src/models/dtos"
	"github.com/awi93/mini-ecommerce/src/models/entities"
)

func (s *repo) composeCustomer(data *dtos.CsvCustomer) *entities.Customer {
	return &entities.Customer{
		Id:              data.Id,
		FirstName:       data.FirstName,
		LastName:        data.LastName,
		Email:           data.Email,
		PhoneNumber:     data.PhoneNumber,
		FaxNumber:       data.FaxNumber,
		BillingAddress:  data.BillingAddress,
		BillingCity:     data.BillingCity,
		BillingState:    data.BillingState,
		BillingZipCode:  data.BillingZipCode,
		CompanyName:     data.CompanyName,
		CompanyWebsite:  data.CompanyWebsite,
		ShipAddress:     data.ShipAddress,
		ShipCity:        data.ShipAddress,
		ShipState:       data.ShipState,
		ShipZipCode:     data.ShipZipCode,
		ShipPhoneNumber: data.ShipPhoneNumber,
	}
}

func (s *repo) CreateFromCsv(data *dtos.CsvCustomer) (*entities.Customer, error) {
	customer := s.composeCustomer(data)
	result := s.DB.Create(&customer)
	return customer, result.Error
}
