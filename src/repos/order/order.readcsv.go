package order

import (
	"encoding/csv"
	"errors"
	"strconv"

	"github.com/awi93/mini-ecommerce/src/models/dtos"
	"github.com/golang-module/carbon/v2"
)

func (r *repo) ReadOrderCsv(reader *csv.Reader) ([]*dtos.CsvOrder, error) {
	var results = make([]*dtos.CsvOrder, 0)
	var i = 0
	for {
		row, err := reader.Read()
		if err != nil {
			break
		}
		if len(row) < 11 {
			return nil, errors.New("invalid csv file")
		}
		if i == 0 {
			i++
			continue
		}
		id, err := strconv.ParseInt(row[0], 10, 64)
		if err != nil {
			return nil, errors.New("invalid id value")
		}
		customerId, err := strconv.ParseInt(row[1], 10, 64)
		if err != nil {
			return nil, errors.New("invalid customer id value")
		}
		employeeId, err := strconv.ParseInt(row[2], 10, 64)
		if err != nil {
			return nil, errors.New("invalid employee id value")
		}
		orderDate := carbon.ParseByFormat(row[3], "m-d-Y").Carbon2Time()
		shipDate := carbon.ParseByFormat(row[5], "m-d-Y").Carbon2Time()
		shippingMethodId, err := strconv.ParseInt(row[6], 10, 64)
		if err != nil {
			return nil, errors.New("invalid shipping method id value")
		}
		freightCharge, err := strconv.ParseFloat(row[7], 64)
		if err != nil {
			return nil, errors.New("invalid freight charge value")
		}
		taxes, err := strconv.ParseFloat(row[8], 64)
		if err != nil {
			return nil, errors.New("invalid taxes value")
		}
		paymentReceived, err := strconv.ParseBool(row[9])
		if err != nil {
			return nil, errors.New("invalid payment received value")
		}
		var csvRow = dtos.CsvOrder{
			Id:                  id,
			CustomerId:          customerId,
			EmployeeId:          employeeId,
			OrderDate:           orderDate,
			PurchaseOrderNumber: row[4],
			ShipDate:            shipDate,
			ShippingMethodId:    shippingMethodId,
			FreightCharge:       freightCharge,
			Taxes:               taxes,
			PaymentReceived:     paymentReceived,
		}

		if row[10] != "" {
			csvRow.Comment = &row[10]
		}

		results = append(results, &csvRow)
	}
	return results, nil
}
