package main

import (
	"github.com/gin-gonic/gin"

	"github.com/MarcelaCuellarML/backpack-bcgow6-marcela-cuellar/GoWeb1/Clase2/cmd/server/handler"
	"github.com/MarcelaCuellarML/backpack-bcgow6-marcela-cuellar/GoWeb1/Clase2/internal/products"
	"github.com/MarcelaCuellarML/backpack-bcgow6-marcela-cuellar/GoWeb1/Clase2/pkg/store"
)

// @title           Bootcamp Go Wave 6 - API - Marcela Cuellar
// @version         1.0
// @description     This is a simple API development by Marcela Cuellar to the Bootcamp Go Backend
// @termsOfService  https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name   API Support Digital House
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/productos

func main() {
	db := store.New(store.FileType, "./products.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	handler := handler.NewProduct(service)

	router := gin.Default()

	pr := router.Group("/products")
	pr.POST("/AddProduct", handler.CreateProduct())
	pr.GET("/", handler.GetAll())
	pr.PUT("/UpdateProduct/:id", handler.UpdateProduct())
	pr.GET("/:id", handler.GetProductByID())
	pr.DELETE("/:id", handler.DeleteElement())
	pr.PATCH("/:id", handler.UpdateQuantity())

	router.Run()
}
