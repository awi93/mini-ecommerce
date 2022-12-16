package employee

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/awi93/mini-ecommerce/bootstrap"
	repo "github.com/awi93/mini-ecommerce/src/repos/employee"
	svc "github.com/awi93/mini-ecommerce/src/services/employee"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var employeeService svc.Service

func Init(config *viper.Viper) error {
	db, err := bootstrap.InitDB(config)
	if err != nil {
		return err
	}
	employeeRepo := repo.NewRepo(&repo.Args{
		DB: db,
	})
	employeeService = svc.NewService(&svc.Args{
		Repo: employeeRepo,
	})

	return nil
}

var Command = &cobra.Command{
	Use:   "employee",
	Short: "Import Employee Data",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("Reading Configuration File")
		err := bootstrap.ReadConfig("config.yml")
		if err != nil {
			return err
		}
		log.Println("Configuration File Read")

		log.Println("Initializing Dependency")
		err = Init(bootstrap.GetConfig())
		if err != nil {
			return err
		}
		log.Println("Dependency Initialized")

		filename := cmd.Flag("file").Value.String()
		file, err := os.Open(filename)
		if err != nil {
			return err
		}

		r := csv.NewReader(file)
		r.Comma = ';'

		results, err := employeeService.Import(r)
		if err != nil {
			return err
		}

		for _, result := range results {
			fmt.Printf("-- Row : [%d], Status : %s\n", result.Index, result.Status)
			for _, err := range result.Errors {
				fmt.Printf("---- %s \n", err)
			}
		}

		return nil
	},
}
