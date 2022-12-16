package entities

type Customer struct {
	Id              int64   `json:"id" gorm:"id"`
	FirstName       string  `json:"first_name" gorm:"first_name"`
	LastName        *string `json:"last_name" gorm:"last_name"`
	Email           *string `json:"email" gorm:"email"`
	PhoneNumber     *string `json:"phone_number" gorm:"phone_number"`
	FaxNumber       *string `json:"fax_number" gorm:"fax_number"`
	BillingAddress  *string `json:"billing_address" gorm:"billing_address"`
	BillingCity     *string `json:"billing_city" gorm:"billing_city"`
	BillingState    *string `json:"billing_state" gorm:"billing_state"`
	BillingZipCode  *string `json:"billing_zip_code" gorm:"billing_zip_code"`
	CompanyName     *string `json:"company_name" gorm:"company_name"`
	CompanyWebsite  *string `json:"company_website" gorm:"company_website"`
	ShipAddress     *string `json:"ship_address" gorm:"ship_address"`
	ShipCity        *string `json:"ship_city" gorm:"ship_city"`
	ShipState       *string `json:"ship_state" gorm:"ship_state"`
	ShipZipCode     *string `json:"ship_zip_code" gorm:"ship_zip_code"`
	ShipPhoneNumber *string `json:"ship_phone_number" gorm:"ship_phone_number"`
}
