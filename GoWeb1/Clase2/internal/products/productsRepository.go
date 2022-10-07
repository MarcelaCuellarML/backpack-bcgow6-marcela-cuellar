package products

import (
	"fmt"

	"github.com/MarcelaCuellarML/backpack-bcgow6-marcela-cuellar/GoWeb1/Clase2/pkg/store"
	//"github.com/MarcelaCuellarML/backpack-bcgow6-marcela-cuellar/GoWeb1/Clase2/internal/domain"
)

var products []domains.Products

type repositoryProducts struct {
	db store.Store
}

func NewRepository(db store.Store) ProductsRepository {
	return &repositoryProducts{
		db: db,
	}
}

type ProductsRepository interface {
	GetAll() ([]products, error)
	GetProductByID(id int) (products, error)
	AgregarProducto(product products) (string, error)
	NewID() (NewId int)
	ActualizarProducto(id int, product products) (string, error)
}

// Funcion para obtener el slice de elementos
func (rp *repositoryProducts) GetAll() ([]products, error) {
	err := rp.db.Read(&products)
	if err != nil {
		return nil, err
	}
	return products, nil
}

// Funcion para obtener un elemento por su ID
func (rp *repositoryProducts) GetProductByID(id int) (products, error) {
	err := rp.db.Read(&products)
	if err != nil {
		return nil, err
	}
	var productosFiltrados products
	for _, prod := range products {
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
func (r *repositoryProducts) AgregarProducto(stock int, name, color, code, creationDate string, price float64, published bool) (products, error) {

	var productsList []products
	err := r.db.Read(&productsList)
	if err != nil {
		return products{}, err
	}
	newId, _ := GenerateNewID()
	newProduct := products{newId, name, color, price, stock, code, published, creationDate}
	fmt.Println("nuevo producto a agregar: ", newProduct)
	productsList = append(productsList, newProduct)
	if err := r.db.Write(productsList); err != nil {
		return products{}, err
	}

	return newProduct, nil
}

// Funcion para obtener un nuevo id
func GenerateNewID() (int, error) {
	rp := repositoryProducts{}
	err := rp.db.Read(&products)
	if err != nil {
		return 0, err
	}
	newId := len(products) + 1
	fmt.Println("Nuevo id generado en el servicio: ", newId)
	return newId, nil
}

func (rp *repositoryProducts) ActualizarProducto(id int, product products) (string, error) {
	err := rp.db.Read(&products)
	if err != nil {
		return "", err
	}
	for i := 0; i < len(products); i++ {
		if products[i].Id == id {
			products[i] = product
		}
	}
	fmt.Println("producto filtrado: ", product)
	return product, nil
}
