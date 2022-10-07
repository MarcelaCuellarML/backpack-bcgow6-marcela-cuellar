package products

import "fmt"

type ProductsService interface {
	GetAll() ([]Products, error)
	//GetProductByID() (Products, error)
	AgregarProducto(Stock int, Name, Color, Code, CreationDate string, Price float64, Published bool) (string, error)
	ActualizarProducto(Id, Stock int, Name, Color, Code, CreationDate string, Price float64, Published bool) (string, error)
}

type serviceProducts struct {
	repo ProductsRepository
}

func NewService(pr ProductsRepository) ProductsService {
	return &serviceProducts{
		repo: pr,
	}
}

// Funcion para obtener el slice de elementos
func (s *serviceProducts) GetAll() ([]Products, error) {
	ps, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return ps, nil
}

// Funcion para obtener un elemento por su ID
// func (s *serviceProducts) GetProductByID(id int) (Products, error) {
// 	ps, err := s.repo.GetProductByID(id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return ps, nil
// }

// Funcion para agregar un producto al slice
func (s *serviceProducts) AgregarProducto(Stock int, Name, Color, Code, CreationDate string, Price float64, Published bool) (string, error) {
	fmt.Println("llegue al servicio de creacion de producto")
	newId := s.repo.NewID()
	prod := Products{newId, Name, Color, Price, Stock, Code, Published, CreationDate}
	ps, err := s.repo.AgregarProducto(prod)
	fmt.Println("Producto que llega al servicio: ", prod.Name)
	if err != nil {
		return "nil", err
	}
	return ps, nil
}

func (s *serviceProducts) ActualizarProducto(Id, Stock int, Name, Color, Code, CreationDate string, Price float64, Published bool) (string, error) {
	fmt.Println("llegue al servicio de actualizacion de producto")
	prod := Products{Id, Name, Color, Price, Stock, Code, Published, CreationDate}
	ps, err := s.repo.ActualizarProducto(Id, prod)
	fmt.Println("Producto que llega al servicio: ", prod.Name)
	if err != nil {
		return "nil", err
	}
	return ps, nil
}
