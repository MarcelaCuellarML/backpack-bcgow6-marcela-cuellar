package main

import (
	"backpack-bcgow6-marcela-cuellar/GoWeb1/Clase2/cmd/server/handler"
	"backpack-bcgow6-marcela-cuellar/GoWeb1/Clase2/internal/products"

	"github.com/gin-gonic/gin"
)

func main() {
	repo := products.NewRepository()
	service := products.NewService(repo)
	handler := handler.NewProduct(service)

	router := gin.Default()

	pr := router.Group("/products")
	pr.POST("/AddProduct", handler.CreateProduct())
	pr.GET("/", handler.GetAll())

	router.Run()
}
