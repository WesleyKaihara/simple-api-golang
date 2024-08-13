package main

import (
	"api-golang/controller"
	"api-golang/db"
	"api-golang/repository"
	"api-golang/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()

	if err != nil {
		panic(err)
	}

	// Repository Layer
	ProductRepository := repository.NewProductRepository(dbConnection)

	// UseCase Layer
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)

	// Controllers Layer
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.GET("/product/:productId", ProductController.GetProductById)
	server.POST("/product", ProductController.CreateProduct)

	server.Run(":8000")
}
