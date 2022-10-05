package products

import (
	"encoding/json"
	"fmt"
	"os"
)

type Products struct {
	Id           int
	Name         string
	Color        string
	Price        float64
	Stock        int
	Code         string
	Published    bool
	CreationDate string
}

type repositoryProducts struct{}

func NewRepository() ProductsRepository {
	return &repositoryProducts{}
}

type ProductsRepository interface {
	GetAll() ([]Products, error)
	GetProductByID(id int) (Products, error)
	AgregarProducto(product Products) (string, error)
}

// ***********FUNCIONES DE INTERACCION CON EL ARCHIVO*******************
func leerJson() []byte {
	readFile, err := os.ReadFile("/Users/marcuellar/backpack-bcgow6-marcela-cuellar/GoWeb1/Clase2/internal/products/productos.json")
	if err != nil {
		panic(err)
	}
	//listado := fmt.Sprint(string(readFile))
	return readFile
}

func TransformData(listado []byte) (product []Products) {
	err := json.Unmarshal(listado, &product)
	if err != nil {
		panic(err)
	}
	//fmt.Println("archivo prueba: ", product)
	return product
}

func GetList() (prod []Products) {
	respuesta := leerJson()
	listadoProds := TransformData(respuesta)
	return listadoProds
}

//*********************FIN****************************

// Funcion para obtener el slice de elementos
func (rp *repositoryProducts) GetAll() ([]Products, error) {
	listadoProds := GetList()
	return listadoProds, nil
}

// Funcion para obtener un elemento por su ID
func (rp *repositoryProducts) GetProductByID(id int) (Products, error) {
	listadoProds := GetList()
	var productosFiltrados Products
	for _, prod := range listadoProds {
		if prod.Id == id {
			productosFiltrados.Id = prod.Id
			productosFiltrados.Name = prod.Name
			productosFiltrados.Color = prod.Color
			productosFiltrados.Price = prod.Price
			productosFiltrados.Stock = prod.Stock
			productosFiltrados.Code = prod.Code
			productosFiltrados.Published = prod.Published
			productosFiltrados.CreationDate = prod.CreationDate
		}
	}
	return productosFiltrados, nil
}

// Funcion para agregar un producto al slice
func (rp *repositoryProducts) AgregarProducto(product Products) (string, error) {
	fmt.Println("llegue al repositorio")
	listadoProds := GetList()
	listadoProds = append(listadoProds, product)
	fmt.Println("listado productos: ", listadoProds)
	resp := "200"
	return resp, nil
}
