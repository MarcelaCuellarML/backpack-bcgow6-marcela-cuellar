package handler

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/MarcelaCuellarML/backpack-bcgow6-marcela-cuellar/GoWeb1/Clase2/internal/products"
	//"github.com/MarcelaCuellarML/backpack-bcgow6-marcela-cuellar/GoWeb1/Clase2/pkg/web"
)

type request struct {
	Id           int     `json:"id"`
	Name         string  `json:"name" binding:"required"`
	Color        string  `json:"color" binding:"required"`
	Price        float64 `json:"price" binding:"required"`
	Stock        int     `json:"stock" binding:"required"`
	Code         string  `json:"code" binding:"required"`
	Published    bool    `json:"published" binding:"required"`
	CreationDate string  `json:"creationDate" binding:"required"`
}

type Products struct {
	service products.ProductsService
}

func NewProduct(p products.ProductsService) *Products {
	return &Products{
		service: p,
	}
}

// /Funcion para validar el token
func validateToken(ctx *gin.Context) string {
	token := ctx.Request.Header.Get("token")
	if token != os.Getenv("TOKEN") || token == "" {
		return "no autorizado"
	}
	return "autorizado"
}

//*****************************************

// Store Product godoc
// @Summary  Store product
// @Tags     Products
// @Accept   json
// @Produce  json
// @Param    token    header    string          true  "token requerido"
// @Param    product  body      ProductRequest  true  "Product to Store"
// @Success  200      {object}  web.Response
// @Failure  401      {object}  web.Response
// @Failure  400      {object}  web.Response
// @Failure  409      {object}  web.Response
// @Router   /products [POST]
func (prod *Products) CreateProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		validToken := validateToken(ctx)
		if validToken == "no autorizado" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		} else {
			var productAdd request
			if err := ctx.ShouldBindJSON(&productAdd); err != nil {
				fmt.Println("Error al recibir json")
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}

			p, err := prod.service.AgregarProducto(productAdd.Stock, productAdd.Name, productAdd.Color, productAdd.Code, productAdd.CreationDate, productAdd.Price, productAdd.Published)
			fmt.Println("valor de p: ", p)
			if err != nil {
				ctx.JSON(http.StatusNotFound, gin.H{
					"error": err.Error(),
				})
				return
			}
			ctx.JSON(http.StatusOK, p)

		}

	}
}

// ListProducts godoc
// @Summary  Show list products
// @Tags     Products
// @Produce  json
// @Param    token  header    string        true  "token"
// @Success  200    {object}  web.Response  "List products"
// @Failure  401    {object}  web.Response  "Unauthorized"
// @Failure  500    {object}  web.Response  "Internal server error "
// @Failure  404    {object}  web.Response  "Not found products"
// @Router   /products [GET]
func (prod *Products) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		validToken := validateToken(ctx)
		if validToken == "no autorizado" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		} else {
			p, err := prod.service.GetAll()
			if err != nil {
				ctx.JSON(http.StatusNotFound, gin.H{
					"error": err.Error(),
				})
				return
			}
			ctx.JSON(200, p)
		}
	}
}

func (prod *Products) UpdateProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		validToken := validateToken(ctx)
		if validToken == "no autorizado" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		} else {
			var productUpdate request
			if err := ctx.ShouldBindJSON(&productUpdate); err != nil {
				fmt.Println("Error al recibir json")
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}

			id, err := strconv.Atoi(ctx.Param("id"))
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id invalido - " + err.Error()})
				return
			}

			p, err := prod.service.ActualizarProducto(id, productUpdate.Stock, productUpdate.Name, productUpdate.Color, productUpdate.Code, productUpdate.CreationDate, productUpdate.Price, productUpdate.Published)
			fmt.Println("valor de p: ", p)
			if err != nil {
				ctx.JSON(http.StatusNotFound, gin.H{
					"error": err.Error(),
				})
				return
			}
			ctx.JSON(http.StatusOK, p)

		}

	}
}
