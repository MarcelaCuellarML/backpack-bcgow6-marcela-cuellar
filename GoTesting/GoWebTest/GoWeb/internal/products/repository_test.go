package products

import "github.com/MarcelaCuellarML/backpack-bcgow6-marcela-cuellar/GoWeb1/Clase2/internal/domain"

type StubReadProducts struct{}
var prod domain.Products


func (rp StubReadProducts) GetAllProds() (prod domain.Products) {
	prod1 := prod{
		prod.Id = 1
		prod.Name = "leche"

	}
	return "12345678" // notar que en lugar de retornar "", retornamos un valor concreto
}
