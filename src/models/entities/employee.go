package entities

type Employee struct {
	Id        int64   `json:"id" gorm:"id"`
	FirstName string  `json:"first_name" gorm:"first_name"`
	LastName  *string `json:"last_name" gorm:"last_name"`
	Title     string  `json:"title" gorm:"title"`
	WorkPhone *string `json:"work_phone" gorm:"work_phone"`
}
