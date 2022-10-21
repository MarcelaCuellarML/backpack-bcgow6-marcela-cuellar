package products

import (
	"fmt"

	"github.com/MarcelaCuellarML/backpack-bcgow6-marcela-cuellar/GoTesting/GoWebTest/GoWeb/internal/domain"
	"github.com/MarcelaCuellarML/backpack-bcgow6-marcela-cuellar/GoWeb1/Clase2/pkg/store"
)

var products []domain.Products

type repositoryProducts struct {
	db store.Store
}

func NewRepository(db store.Store) ProductsRepository {
	return &repositoryProducts{
		db: db,
	}
}

type ProductsRepository interface {
	GetAll() ([]domain.Products, error)
	GetProductByID(id int) (domain.Products, error)
	AgregarProducto(product domain.Products) (domain.Products, error)
	GenerateNewID() (int, error)
	UpdateProduct(id int, newProduct domain.Products) (domain.Products, error)
	DeleteElement(id int) ([]domain.Products, error)
	UpdateQuantity(id, cant int) (domain.Products, error)
}

// Funcion para obtener el listado de elementos
func (rp *repositoryProducts) GetAll() ([]domain.Products, error) {
	err := rp.db.Read(&products)
	if err != nil {
		return nil, err
	}
	return products, nil
}

// Funcion para obtener un elemento por su ID
func (rp *repositoryProducts) GetProductByID(id int) (domain.Products, error) {
	err := rp.db.Read(&products)
	if err != nil {
		return domain.Products{}, err
	}
	var productosFiltrados domain.Products
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
func (r *repositoryProducts) AgregarProducto(product domain.Products) (domain.Products, error) {

	var productsList []domain.Products
	err := r.db.Read(&productsList)
	if err != nil {
		return domain.Products{}, err
	}
	fmt.Println("nuevo producto a agregar: ", product)
	productsList = append(productsList, product)
	if err := r.db.Write(productsList); err != nil {
		return domain.Products{}, err
	}
	return product, nil
}

// Funcion para obtener un nuevo id
func (rp *repositoryProducts) GenerateNewID() (int, error) {
	err := rp.db.Read(&products)
	if err != nil {
		return 0, err
	}
	newId := len(products) + 1
	fmt.Println("Nuevo id generado en el servicio: ", newId)
	return newId, nil
}

// Funcion para acualizar un producto
func (rp *repositoryProducts) UpdateProduct(id int, newProduct domain.Products) (domain.Products, error) {
	err := rp.db.Read(&products)
	if err != nil {
		return domain.Products{}, err
	}
	for i := 0; i < len(products); i++ {
		if products[i].Id == id {
			products[i] = newProduct
		}
	}
	rp.db.Write(products)
	fmt.Println("producto filtrado: ", newProduct)
	return newProduct, nil
}

// Funcion para eliminar un elemento
func (rp *repositoryProducts) DeleteElement(id int) ([]domain.Products, error) {
	err := rp.db.Read(&products)
	if err != nil {
		return []domain.Products{}, err
	}
	for i := 0; i < len(products); i++ {
		if products[i].Id == id {
			products = append(products[:i], products[i+1:]...)
		}
	}
	rp.db.Write(products)
	// err = rp.db.Read(&products)
	// if err != nil {
	// 	return []domain.Products{}, err
	// }
	fmt.Println("producto eliminado")
	return products, nil
}

// funcion para actualizar la cantidad en stock de un producto
func (rp *repositoryProducts) UpdateQuantity(id, cant int) (domain.Products, error) {
	err := rp.db.Read(&products)
	var product domain.Products
	if err != nil {
		return domain.Products{}, err
	}
	for i := 0; i < len(products); i++ {
		if products[i].Id == id {
			//fmt.Println("cantidad del producto solicitado: ", products[i].Stock)
			products[i].Stock = cant
			product = products[i]
			//fmt.Println("producto cambiado: ", product)
		}

	}
	rp.db.Write(product)
	fmt.Println("producto cambiado: ", products)
	return product, nil
}
