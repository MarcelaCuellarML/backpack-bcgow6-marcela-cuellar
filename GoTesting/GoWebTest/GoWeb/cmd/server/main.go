package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/MarcelaCuellarML/backpack-bcgow6-marcela-cuellar/GoTesting/GoWebTest/GoWeb/cmd/server/handler"
	"github.com/MarcelaCuellarML/backpack-bcgow6-marcela-cuellar/GoTesting/GoWebTest/GoWeb/docs"
	"github.com/MarcelaCuellarML/backpack-bcgow6-marcela-cuellar/GoTesting/GoWebTest/GoWeb/internal/products"
	"github.com/MarcelaCuellarML/backpack-bcgow6-marcela-cuellar/GoTesting/GoWebTest/GoWeb/pkg/store"
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
// @BasePath  /

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("no ha sido posible cargar wl archivo .env")
	}

	db := store.New(store.FileType, "/Users/marcuellar/backpack-bcgow6-marcela-cuellar/GoWeb1/Clase2/cmd/server/productos.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	handler := handler.NewProduct(service)

	router := gin.Default()
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	pr := router.Group("/products")
	pr.POST("/AddProduct", handler.CreateProduct())
	pr.GET("/", handler.GetAll())
	pr.PUT("/UpdateProduct/:id", handler.UpdateProduct())
	pr.GET("/:id", handler.GetProductByID())
	pr.DELETE("/:id", handler.DeleteElement())
	pr.PATCH("/:id", handler.UpdateQuantity())

	err = router.Run()
	if err != nil {
		return
	}
}
