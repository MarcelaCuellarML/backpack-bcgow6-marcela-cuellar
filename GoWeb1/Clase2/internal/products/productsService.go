package products

import "fmt"

type ProductsService interface {
	GetAll() ([]Products, error)
	//GetProductByID() (Products, error)
	AgregarProducto(product Products) (string, error)
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
func (s *serviceProducts) AgregarProducto(Id, Stock int, Name, Color, Code, CreationDate string, Price float64, Published bool) (string, error) {
	ps, err := s.repo.AgregarProducto(product)
	fmt.Println("llegue al servicio")
	fmt.Println("Producto que llega al servicio: ", product.Name)
	if err != nil {
		return "nil", err
	}
	return ps, nil
}
