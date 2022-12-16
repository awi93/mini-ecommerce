package dtos

type CsvCustomer struct {
	Id              int64
	FirstName       string
	LastName        *string
	Email           *string
	PhoneNumber     *string
	FaxNumber       *string
	BillingAddress  *string
	BillingCity     *string
	BillingState    *string
	BillingZipCode  *string
	CompanyName     *string
	CompanyWebsite  *string
	ShipAddress     *string
	ShipCity        *string
	ShipState       *string
	ShipZipCode     *string
	ShipPhoneNumber *string
}
