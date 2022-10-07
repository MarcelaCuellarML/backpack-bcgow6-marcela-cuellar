package main

import (
	"github.com/gin-gonic/gin"

	"github.com/MarcelaCuellarML/backpack-bcgow6-marcela-cuellar/GoWeb1/Clase2/cmd/server/handler"
	"github.com/MarcelaCuellarML/backpack-bcgow6-marcela-cuellar/GoWeb1/Clase2/internal/products"
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
