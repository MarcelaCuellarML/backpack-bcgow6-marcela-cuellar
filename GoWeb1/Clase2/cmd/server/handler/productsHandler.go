package handler

import (
	"fmt"
	"net/http"

	"backpack-bcgow6-marcela-cuellar/GoWeb1/Clase2/internal/products"

	"github.com/gin-gonic/gin"
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

func (prod *Products) CreateProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "12345" || token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			return
		} else {
			var productAdd request
			if err := ctx.ShouldBindJSON(&productAdd); err != nil {
				fmt.Println("Error al recibir json")
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}

			productoCrear := products.Products{
				Name:  productAdd.Name,
				Color: productAdd.Color,
				Price: productAdd.Price,
				Stock: productAdd.Stock,
				Code:  productAdd.Code,
			}
			p, err := prod.service.AgregarProducto(productAdd.Name, productAdd.Color, productAdd.Price, productAdd.Stock, productAdd.Code)
			fmt.Println("valor de p: ", p)
			if err != nil {
				ctx.JSON(http.StatusNotFound, gin.H{
					"error": err.Error(),
				})
				return
			}
			ctx.JSON(http.StatusOK, p)

			// if err := ctx.ShouldBindJSON(&productAdd); err != nil {
			// 	prod.service.AgregarProducto(products.Products(productAdd))
			// 	fmt.Println("producto creado: ", products.Products(productAdd))
			// 	ctx.JSON(http.StatusOK, products.Products(productAdd))
			// } else {
			// 	ctx.JSON(http.StatusBadRequest, "error de guardado y llamado a guardar")
			// }

		}

	}
}

func (prod *Products) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "12345" || token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			return
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
