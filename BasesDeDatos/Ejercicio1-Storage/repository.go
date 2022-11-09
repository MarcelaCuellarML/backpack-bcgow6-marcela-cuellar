package storage

import (
	"database/sql"
	"log"
)

var (
	StorageDB *sql.DB
)

func conn() {
	dataSource := "root:abcd1234@tcp(localhost:3306)/storage"
	// Open inicia un pool de conexiones. Sólo abrir una vez
	var err error
	StorageDB, err = sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	if err = StorageDB.Ping(); err != nil {
		panic(err)
	}
	log.Println("database Configured")
}

type Repository interface {
	Store(name, productType string, count int, price float64) (model.Product, error)
	GetOne(id int)
	Update(product model.Product) (model.Product, error)
	GetAll() ([]model.Product, error)
	Delete(id int) error
}
type repository struct{}

func NewRepo() Repository {
	return &repository{}
}

func (r *repository) Store(product model.Product) (models.Product, error) {
	db := db.StorageDB                                                                             // se inicializa la base
	stmt, err := db.Prepare("INSERT INTO products(name, type, count, price) VALUES( ?, ?, ?, ? )") // se prepara el SQL
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
	var result sql.Result
	result, err = stmt.Exec(product.Name, product.Type, product.Count, product.Price) // retorna un sql.Result y un error
	if err != nil {
		return models.Product{}, err
	}
	insertedId, _ := result.LastInsertId() // del sql.Resul devuelto en la ejecución obtenemos el Id insertado
	product.ID = int(insertedId)

	return product, nil
}
