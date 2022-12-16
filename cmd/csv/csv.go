package csv

import (
	"github.com/awi93/mini-ecommerce/cmd/csv/customer"
	"github.com/awi93/mini-ecommerce/cmd/csv/employee"
	"github.com/awi93/mini-ecommerce/cmd/csv/order"
	"github.com/awi93/mini-ecommerce/cmd/csv/orderdetail"
	"github.com/awi93/mini-ecommerce/cmd/csv/product"
	"github.com/awi93/mini-ecommerce/cmd/csv/shippingmethod"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "csv",
	Short: "Http Service",
}

func init() {
	Command.AddCommand(customer.Command)
	Command.AddCommand(employee.Command)
	Command.AddCommand(order.Command)
	Command.AddCommand(orderdetail.Command)
	Command.AddCommand(product.Command)
	Command.AddCommand(shippingmethod.Command)

	customer.Command.PersistentFlags().StringP("file", "f", "", "please input csv file to import")
	employee.Command.PersistentFlags().StringP("file", "f", "", "please input csv file to import")
	order.Command.PersistentFlags().StringP("file", "f", "", "please input csv file to import")
	orderdetail.Command.PersistentFlags().StringP("file", "f", "", "please input csv file to import")
	product.Command.PersistentFlags().StringP("file", "f", "", "please input csv file to import")
	shippingmethod.Command.PersistentFlags().StringP("file", "f", "", "please input csv file to import")

	customer.Command.MarkPersistentFlagRequired("file")
	employee.Command.MarkPersistentFlagRequired("file")
	order.Command.MarkPersistentFlagRequired("file")
	orderdetail.Command.MarkPersistentFlagRequired("file")
	product.Command.MarkPersistentFlagRequired("file")
	shippingmethod.Command.MarkPersistentFlagRequired("file")
}
