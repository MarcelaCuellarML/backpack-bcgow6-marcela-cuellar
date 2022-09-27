package main

import "fmt"

type Perro struct {
	Nombre string
	Edad   int
	Raza   string
	Peso   float64
	Color  []string
	owner  Dueño
	//Cuando un elemento empieza con minuscula sera privada para el resto del paquete, pero si esta en mayuscula es publica
}

func (p Perro) Ladrar() {
	fmt.Println("Guau")
}

func (p Perro) Saludar() {
	fmt.Println("Guau soy ", p.Nombre)
}

func (p Perro) Renombrar(newName string) {
	p.Nombre = newName
	fmt.Println("Guau ahora me llamo ", p.Nombre)
}

func (p Perro) Dormir() {
	fmt.Println("Estoy durmiendo")
}

func (g Gato) Dormir() {
	fmt.Println("Estoy durmiendo")
}

func (g Perro) Comer() {
	fmt.Println("Estoy comiendo")
}

type Dueño struct {
	Documento   int
	Nombre      string
	Nrocontacto int
}

type Gato struct {
	Nombre   string
	Edad     int
	Raza     string
	Peso     float64
	Color    []string
	Caracter string
	owner    Dueño
}

type Animal interface {
	Dormir()
	Comer()
}

func NewAnimalPerro() Animal {
	return &Perro{}
}

func main() {

	dog := Perro{
		Nombre: "Española", Edad: 5, Raza: "husky", Peso: 5.5, Color: []string{"blanco", "gris", "negro"},
		owner: Dueño{
			Documento:   123456,
			Nombre:      "Dani",
			Nrocontacto: 312542125,
		},
	}

	fmt.Println(dog)
	dog.Ladrar()
	dog.Saludar()
	dog.Renombrar("Pirulina")

	a := NewAnimalPerro()
	a.Comer()
	a.Dormir()
}
