package main

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"main/internal/core/services"
	"main/internal/driven"
	"main/internal/handlers"
)

func main() {
	transferRepository := driven.NewTransferRepository(&mongo.Database{})

	customerClient := driven.NewCustomerClient(driven.NewDefaultHttpAdapter())
	customerService := services.NewCustomerService(customerClient)

	transferAuthorizationClient := driven.NewTransferClient(driven.NewDefaultHttpAdapter())
	transferAuthorizationService := services.NewTransferAuthorizationService(transferAuthorizationClient)

	transferService := services.NewTransferService(transferRepository, customerService, transferAuthorizationService)

	handler := handlers.NewHTTPHandler(transferService)

	router := gin.New()
	router.POST("/transfer/", handler.Create)

	router.Run(":8080")
}
