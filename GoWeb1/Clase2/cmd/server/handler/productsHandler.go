package handler

import (
	"net/http"

	"github.com/MarcelaCuellarML/backpack-bcgow6-marcela-cuellar/GoWeb1/Clase2/internal/products"
	"github.com/gin-gonic/gin"
)

type Products struct {
	Id           int     `json:"id"`
	Name         string  `json:"name" binding:"required"`
	Color        string  `json:"color" binding:"required"`
	Price        float64 `json:"price" binding:"required"`
	Stock        int     `json:"stock" binding:"required"`
	Code         string  `json:"code" binding:"required"`
	Published    bool    `json:"published" binding:"required"`
	CreationDate string  `json:"creationDate" binding:"required"`
}

type ProductService struct {
	service products.Service
}

func NewProduct(p products.Service) *Products {
	return &Products{
		service: p,
	}
}

func AgregarProducto(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != "12345" || token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		return
	} else {
		var prod Products
		if err := ctx.ShouldBindJSON(&prod); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		ctx.JSON(http.StatusOK, prod)
	}

}

func GetAll(ctx *gin.Context) {
	if token != "12345" || token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		return
	} else {
		p, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, p)
	}
}
