package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"online-shop/apps/auth"
	"online-shop/apps/products"
	"online-shop/apps/transactions"
	"online-shop/external/database"
	"online-shop/internal/config"

)

func main() {
	filename := "cmd/api/config.yaml"
	if err := config.LoadConfig(filename); err != nil {
		panic(err)
	}

	db, err := database.ConnectionPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	if db != nil {
		log.Println("DB Connected")
	}

	router := fiber.New(fiber.Config{
		Prefork: true,
		AppName: config.Cfg.App.Name,
	})

	auth.Init(router, db)
	products.Init(router, db)
	transactions.Init(router, db)


	if err := router.Listen(config.Cfg.App.Port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
