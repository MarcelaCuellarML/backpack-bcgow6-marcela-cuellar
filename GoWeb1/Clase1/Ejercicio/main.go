/*
Crea dentro de la carpeta go-web un archivo llamado main.go
Crea un servidor web con Gin que te responda un JSON que tenga una clave “message” y diga Hola seguido por tu nombre.
Pegale al endpoint para corroborar que la respuesta sea la correcta.
PUNTO 3
Ya habiendo creado y probado nuestra API que nos saluda, generamos una ruta que devuelve un listado de la temática elegida.
Dentro del “main.go”, crea una estructura según la temática con los campos correspondientes.
Genera un endpoint cuya ruta sea /temática (en plural). Ejemplo: “/productos”
Genera un handler para el endpoint llamado “GetAll”.
Crea una slice de la estructura, luego devuelvela a través de nuestro endpoint.
PUNTO 4
Según la temática elegida, necesitamos agregarles filtros a nuestro endpoint, el mismo se tiene que poder filtrar por todos los campos.
Dentro del handler del endpoint, recibí del contexto los valores a filtrar.
Luego genera la lógica de filtrado de nuestro array.
Devolver por el endpoint el array filtrado.
PUNTO 5
Generar un nuevo endpoint que nos permita traer un solo resultado del array de la temática. Utilizando path parameters el endpoint debería ser /temática/:id (recuerda
que siempre tiene que ser en plural la temática). Una vez recibido el id devuelve la posición correspondiente.
1. Genera una nueva ruta.
2. Genera un handler para la ruta creada.
3. Dentro del handler busca el item que necesitas.
4. Devuelve el item según el id.
Si no encontraste ningún elemento con ese id devolver como código de respuesta 404.
*/

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

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

var ProductList []Products

func main() {
	router := gin.Default()

	router.GET("/saludo", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola Marce",
		})
	})

	//listado de endpoints de acceso
	agrupador := router.Group("/listadoProds")
	{
		//agrupador.GET("/")
		agrupador.GET("/:id", GetProductByID) //con los dos puntos se marca que va a recibir un parametro
		agrupador.GET("", GetProductByFilter) //el filtro no va asi, va con el request,
		agrupador.POST("/NewProd", AgregarProducto)
	}

	router.GET("/listado", func(c *gin.Context) {
		respuesta := leerJson()
		listadoProds := GetAll(respuesta)
		c.JSON(200, gin.H{
			"listado de productos": listadoProds,
		})
	})
	router.Run()

}

func leerJson() []byte {
	readFile, err := os.ReadFile("products.json")
	if err != nil {
		panic(err)
	}
	//listado := fmt.Sprint(string(readFile))
	return readFile
}

func GetAll(listado []byte) (product []Products) {
	err := json.Unmarshal(listado, &product)
	if err != nil {
		panic(err)
	}
	//fmt.Println("archivo prueba: ", product)
	return product
}

//************CLASE DE LA TARDE***************//

func GetProductByID(ctx *gin.Context) {
	respuesta := leerJson()
	listadoProds := GetAll(respuesta)
	var productosFiltrados []Products
	idRecibido, _ := strconv.Atoi(ctx.Param("id"))
	fmt.Println("id recibido: ", idRecibido)
	for _, prod := range listadoProds {
		if prod.Id == idRecibido {
			productosFiltrados = append(productosFiltrados, prod)
		}
	}
	ctx.JSON(http.StatusOK, productosFiltrados)
	//ctx.JSON(http.StatusOK, product) //implementar los datos que se deben retornar en el ctx.JSON
}

func GetProductByFilter(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != "12345" || token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		return
	} else {
		respuesta := leerJson()
		listadoProds := GetAll(respuesta)
		var productosFiltrados []Products
		//PRUEBA 1
		prue, _ := ctx.GetQuery("code")
		fmt.Println("Query: ", prue)
		for _, e := range listadoProds {
			if e.Code == ctx.Query("code") {
				productosFiltrados = append(productosFiltrados, e)
			}
		}
		ctx.JSON(http.StatusOK, productosFiltrados)
	}
}

//**************************+GO WEB DIA 2*****************************************//

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
		respuesta := leerJson()
		listadoProds := GetAll(respuesta)
		ProductList = listadoProds
		prod.Id = len(ProductList) + 1
		ProductList = append(ProductList, prod)
		ctx.JSON(http.StatusOK, prod)
	}

}
