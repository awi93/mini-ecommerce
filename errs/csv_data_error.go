package errs

type CsvData struct {
	Index  int
	Errors []string
}

type CsvDataError struct {
	Errors []*CsvData
}

func (c *CsvDataError) Error() string {
	return "invalid csv data given"
}

func NewCsvDataError(Errors []*CsvData) *CsvDataError {
	return &CsvDataError{
		Errors: Errors,
	}
}
