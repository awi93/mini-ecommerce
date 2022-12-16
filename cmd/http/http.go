package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/awi93/mini-ecommerce/bootstrap"
	"github.com/awi93/mini-ecommerce/src/adapters/https"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "http",
	Short: "Http Service",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("Reading Configuration File")
		bootstrap.ReadConfig("config.yml")
		log.Println("Configuration File Read")
		log.Println("Initializing Dependency")
		https.InitDependency(bootstrap.GetConfig())
		log.Println("Dependency Initialized")

		log.Println("Registering route to router")
		router := https.Router()
		log.Println("Route registered")
		log.Printf("Starting Http Service at :%v\n", bootstrap.GetConfig().GetString("http.port"))
		return http.ListenAndServe(fmt.Sprintf(":%s", bootstrap.GetConfig().GetString("http.port")), router)
	},
}
