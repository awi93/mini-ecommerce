package cmd

import (
	"log"
	"os"

	"github.com/awi93/mini-ecommerce/cmd/csv"
	"github.com/awi93/mini-ecommerce/cmd/http"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "mini-ecommerce",
	Short: "Mini Ecommerce",
}

func Execute() {
	RootCmd.AddCommand(http.Command)
	RootCmd.AddCommand(csv.Command)

	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err.Error())
		os.Exit(-1)
	}
}
