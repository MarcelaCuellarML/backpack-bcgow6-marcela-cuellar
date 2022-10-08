package products

import (
	"fmt"

	"github.com/MarcelaCuellarML/backpack-bcgow6-marcela-cuellar/GoWeb1/Clase2/internal/domain"
)

type serviceProducts struct {
	repo ProductsRepository
}

func NewService(pr ProductsRepository) ProductsService {
	return &serviceProducts{
		repo: pr,
	}
}

type ProductsService interface {
	GetAll() ([]domain.Products, error)
	GetProductByID(id int) (domain.Products, error)
	AgregarProducto(Stock int, Name, Color, Code, CreationDate string, Price float64, Published bool) (domain.Products, error)
	ActualizarProducto(Id, Stock int, Name, Color, Code, CreationDate string, Price float64, Published bool) (domain.Products, error)
	DeleteElement(Id int) ([]domain.Products, error)
	UpdateQuantity(Id, cant int) (domain.Products, error)
}

// Funcion para obtener el slice de elementos
func (s *serviceProducts) GetAll() ([]domain.Products, error) {
	ps, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return ps, nil
}

// Funcion para obtener un elemento por su ID
func (s *serviceProducts) GetProductByID(id int) (domain.Products, error) {
	ps, err := s.repo.GetProductByID(id)
	if err != nil {
		return domain.Products{}, err
	}
	return ps, nil
}

// Funcion para agregar un producto al slice
func (s *serviceProducts) AgregarProducto(Stock int, Name, Color, Code, CreationDate string, Price float64, Published bool) (domain.Products, error) {
	fmt.Println("llegue al servicio de creacion de producto")
	newId, _ := s.repo.GenerateNewID()
	prod := domain.Products{Id: newId, Name: Name, Color: Color, Price: Price, Stock: Stock, Code: Code, Published: Published, CreationDate: CreationDate}
	ps, err := s.repo.AgregarProducto(prod)
	fmt.Println("Producto que llega al servicio: ", prod.Name)
	if err != nil {
		return domain.Products{}, err
	}
	return ps, nil
}

// Funcion para actuazar un producto
func (s *serviceProducts) ActualizarProducto(Id, Stock int, Name, Color, Code, CreationDate string, Price float64, Published bool) (domain.Products, error) {
	fmt.Println("llegue al servicio de actualizacion de producto")
	prod := domain.Products{Id: Id, Name: Name, Color: Color, Price: Price, Stock: Stock, Code: Code, Published: Published, CreationDate: CreationDate}
	ps, err := s.repo.ActualizarProducto(Id, prod)
	fmt.Println("Producto que llega al servicio: ", prod.Name)
	if err != nil {
		return domain.Products{}, err
	}
	return ps, nil
}

func (s *serviceProducts) DeleteElement(Id int) ([]domain.Products, error) {
	fmt.Println("llegue al servicio de eliminacion de producto")
	ps, err := s.repo.DeleteElement(Id)
	fmt.Println("Id del producto a eliminar que llega al servicio: ", Id)
	if err != nil {
		return []domain.Products{}, err
	}
	return ps, nil
}

func (s *serviceProducts) UpdateQuantity(Id, cant int) (domain.Products, error) {
	fmt.Println("llegue al servicio de actualizacion de cantidad de producto")
	ps, err := s.repo.UpdateQuantity(Id, cant)
	if err != nil {
		return domain.Products{}, err
	}
	return ps, nil
}
