package https

import (
	"github.com/awi93/mini-ecommerce/bootstrap"
	orderRepo "github.com/awi93/mini-ecommerce/src/repos/order"
	orderSvc "github.com/awi93/mini-ecommerce/src/services/order"
	"github.com/spf13/viper"
)

var OrderService orderSvc.Service

func InitDependency(config *viper.Viper) error {
	db, err := bootstrap.InitDB(config)
	if err != nil {
		return err
	}
	repo := orderRepo.NewRepo(&orderRepo.Args{
		DB: db,
	})
	OrderService = orderSvc.NewService(&orderSvc.Args{
		Repo: repo,
	})
	return nil
}
