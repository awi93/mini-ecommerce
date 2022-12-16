package dtos

type CsvEmployee struct {
	Id        int64
	FirstName string
	LastName  *string
	Title     string
	WorkPhone *string
}
