package main

import (
	"log"

	"github.com/Teepee/internal/config"
	database "github.com/Teepee/internal/database/impl"
	"github.com/Teepee/internal/database/seeder"
	products "github.com/Teepee/internal/product/controller"
	products_repository "github.com/Teepee/internal/product/repository"
	products_service "github.com/Teepee/internal/product/service"
	shops_controller "github.com/Teepee/internal/shop/controller"
	shops_repository "github.com/Teepee/internal/shop/repository"
	shops_service "github.com/Teepee/internal/shop/service"
	"github.com/Teepee/internal/web"
	"github.com/labstack/echo/v5"
)

func main() {

	server := echo.New()

	config := config.CreateConfig()

	log.Print("Connecting to database")
	db, err := database.ConnectAndMigrateDatabase(config)
	if err != nil {
		// fmt.Print(err)
		log.Fatalf(err.Error())
		return
	}
	log.Print("Database connected")

	if config.Seed != "false" {
		log.Print("Starting to seed database")
		err = seeder.SeedDatabase(db)
		if err != nil {
			log.Fatalf(err.Error())
			return
		}
		log.Print("Finished seeding database")
	}

	server.Static(
		"/static",
		"static",
	)

	productsRepo := products_repository.InitializedProductRepository(db)
	productsService := products_service.InitializedProductsService(productsRepo)
	productsController := products.InitializedProductsController(productsService)
	productsController.InitializedRoute(server)

	shopsRepo := shops_repository.InitializedShopRepository(db)
	shopsService := shops_service.InitializedShopService(shopsRepo, productsRepo)
	shopsController := shops_controller.InititalizedShopController(shopsService)
	shopsController.InitializedRoute(server)

	webController := web.InitializedWebController()
	webController.InitializedRoute(server)

	if err := server.Start("0.0.0.0:1323"); err != nil {
		server.Logger.Error("failed to start server", "error", err)
	}
}
