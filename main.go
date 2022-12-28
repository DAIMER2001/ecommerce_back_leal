package main

import (
	"ecommerce/config"
	"ecommerce/features"
	"ecommerce/repository"
	"ecommerce/server"
	"log"

	"github.com/spf13/viper"
)

func Run() error {
	var cfg config.Config
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath("./")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()

	if err != nil {
		return err
	}

	err = viper.Unmarshal(&cfg)

	if err != nil {
		return err
	}
	db, err := config.InitDB(cfg)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	clientRepository := repository.NewClientRepository(db)
	clientService := features.NewClientFeatures(clientRepository)
	productRepository := repository.NewProductRepository(db)
	productService := features.NewProductFeatures(productRepository)

	svr := server.New(cfg, clientService, productService)
	err = svr.Run()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatalln(err)
	}
}
