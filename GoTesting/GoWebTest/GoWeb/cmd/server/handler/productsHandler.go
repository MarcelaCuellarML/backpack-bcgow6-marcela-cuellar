package handler

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/MarcelaCuellarML/backpack-bcgow6-marcela-cuellar/GoWeb1/Clase2/internal/products"
	"github.com/MarcelaCuellarML/backpack-bcgow6-marcela-cuellar/GoWeb1/Clase2/pkg/web"
)

type Request struct {
	Id           int     `json:"id"`
	Name         string  `json:"name" binding:"required"`
	Color        string  `json:"color" binding:"required"`
	Price        float64 `json:"price" binding:"required"`
	Stock        int     `json:"stock" binding:"required"`
	Code         string  `json:"code" binding:"required"`
	Published    bool    `json:"published" binding:"required"`
	CreationDate string  `json:"creationDate" binding:"required"`
}
type RequestPatch struct {
	Nombre int `json:"cantidad" binding:"required"` 
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
// @Summary  create new product
// @Tags     Products
// @Accept   json
// @Produce  json
// @Param    token    header    string          true  "token requerido"
// @Param    product  body      Request  		true  "Product to Store"
// @Success  200      {object}  web.Response
// @Failure  401      {object}  web.Response
// @Failure  400      {object}  web.Response
// @Failure  409      {object}  web.Response
// @Router   /products/AddProduct [POST]
func (prod *Products) CreateProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		validToken := validateToken(ctx)
		if validToken == "no autorizado" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
		} else {
			var productAdd Request
			if err := ctx.ShouldBindJSON(&productAdd); err != nil {
				fmt.Println("Error al recibir json")
				ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "el elemento JSON  enviado presenta errores, por favor verifiquelo e intente nuevamente"))
			}

			p, err := prod.service.AgregarProducto(productAdd.Stock, productAdd.Name, productAdd.Color, productAdd.Code, productAdd.CreationDate, productAdd.Price, productAdd.Published)
			fmt.Println("valor de p: ", p)
			if err != nil {
				ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "no ha sido posible agregar el producto "))
				return
			}
			ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, p, ""))

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
			ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, p, ""))
		}
	}
}

// ListProducts godoc
// @Summary  Update one product for products
// @Tags     Products
// @Produce  json
// @Param    id     	path 	  int 	        	true   	"Id product"
// @Param    token  	header    string        	true  	"token"
// @Param    product  	body      Request		  	true   "Product to update"
// @Success  200    	{object}  web.Response  			"List products"
// @Failure  401    	{object}  web.Response  			"Unauthorized"
// @Failure  500    	{object}  web.Response  			"Internal server error "
// @Failure  404    	{object}  web.Response  			"Not found products"
// @Router   /products/UpdateProduct/{id} [PUT]

// Funcion para actualizar un producto
func (prod *Products) UpdateProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		validToken := validateToken(ctx)
		if validToken == "no autorizado" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
		} else {
			var productUpdate Request
			if err := ctx.ShouldBindJSON(&productUpdate); err != nil {
				fmt.Println("Error al recibir json")
				ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "el elemento JSON  enviado presenta errores, por favor verifiquelo e intente nuevamente"))
			}

			id, err := strconv.Atoi(ctx.Param("id"))
			if err != nil {
				ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "el id ingresado no es valido"))
				return
			}

			p, err := prod.service.UpdateProduct(id, productUpdate.Stock, productUpdate.Name, productUpdate.Color, productUpdate.Code, productUpdate.CreationDate, productUpdate.Price, productUpdate.Published)
			fmt.Println("valor de p: ", p)
			if err != nil {
				ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "no ha sido posible actualizar elemento solicitado, por favor intente mas tarde"))
				return
			}
			ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, p, ""))

		}

	}
}

// ListProducts godoc
// @Summary  Show product by ID for products
// @Tags     Products
// @Produce  json
// @Param    id     	path 	  int 	        	true   	"Id product"
// @Param    token  	header    string        	true  	"token"
// @Success  200    	{object}  web.Response  			"List products"
// @Failure  401    	{object}  web.Response  			"Unauthorized"
// @Failure  500    	{object}  web.Response  			"Internal server error "
// @Failure  404    	{object}  web.Response  			"Not found products"
// @Router   /products/{id} [GET]

// Funcion para obtener un producto por su id
func (prod *Products) GetProductByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		validToken := validateToken(ctx)
		if validToken == "no autorizado" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		} else {
			id, err := strconv.Atoi(ctx.Param("id"))
			if err != nil {
				ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "el id ingresado no es valido"))
				return
			}
			p, err := prod.service.GetProductByID(id)
			if err != nil {
				ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "no ha sido posible encontrar el elemento solicitado, por favor intente mas tarde"))
				return
			}
			ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, p, ""))
		}
	}
}

// ListProducts godoc
// @Summary  Delete element for products
// @Tags     Products
// @Produce  json
// @Param    id     	path 	  int 	        	true   	"Id product"
// @Param    token  	header    string        	true  	"token"
// @Success  200    	{object}  web.Response  			"List products"
// @Failure  401    	{object}  web.Response  			"Unauthorized"
// @Failure  500    	{object}  web.Response  			"Internal server error "
// @Failure  404    	{object}  web.Response  			"Not found products"
// @Router   /products/{id} [DELETE]

// Funcion para eliminar un producto
func (prod *Products) DeleteElement() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		validToken := validateToken(ctx)
		if validToken == "no autorizado" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		} else {
			id, err := strconv.Atoi(ctx.Param("id"))
			if err != nil {
				ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "el id ingresado no es valido"))
				return
			}
			p, err := prod.service.DeleteElement(id)
			if err != nil {
				ctx.JSON(http.StatusNotFound, gin.H{
					"error": err.Error(),
				})
				return
			}
			ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, p, ""))
		}
	}
}

// Update quantity Product godoc
// @Summary      Update quantity product
// @Tags         Products
// @Accept       json
// @Produce      json
// @Description  This endpoint update field quantity from an product
// @Param        token  	header    string            true  "Token header"
// @Param        id     	path      int               true  "Product Id"
// @Param        quantity   body      RequestPatch		true  "Product quantity"
// @Success      200   	 	{object}  web.Response
// @Failure      401    	{object}  web.Response
// @Failure      400   		{object}  web.Response
// @Failure      404   		{object}  web.Response
// @Router       /products/{id} [PATCH]
// Funcion para actualizar la cantidad de un producto
func (prod *Products) UpdateQuantity() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		validToken := validateToken(ctx)
		if validToken == "no autorizado" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		} else {
			id, err := strconv.Atoi(ctx.Param("id"))
			if err != nil {
				ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "el id ingresado no es valido"))
				return
			}
			cant, err := strconv.Atoi(ctx.Query("Stock"))
			if err != nil {
				ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "la cantidad ingresada no es valida"))
				return
			}
			p, err := prod.service.UpdateQuantity(id, cant)
			if err != nil {
				ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusInternalServerError, nil, "no ha sido posible actualizar la cantidad ingresada, por favor intente mas tarde"))
				return
			}
			ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, p, ""))
		}

	}
}
