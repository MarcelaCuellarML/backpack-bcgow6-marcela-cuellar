/*
Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento.
Para ello requieren realizar un programa que se encargue de calcular el precio total de Productos, Servicios y Mantenimientos. Debido a la fuerte demanda y para
optimizar la velocidad requieren que el cálculo de la sumatoria se realice en paralelo mediante 3 go routines.
Se requieren 3 estructuras:
Productos: nombre, precio, cantidad.
Servicios: nombre, precio, minutos trabajados.
Mantenimiento: nombre, precio.
Se requieren 3 funciones:
Sumar Productos: recibe un array de producto y devuelve el precio total (precio * cantidad).
Sumar Servicios: recibe un array de servicio y devuelve el precio total (precio * media hora trabajada, si no llega a trabajar 30 minutos se le cobra como si hubiese
trabajado media hora).
Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.

Los 3 se deben ejecutar concurrentemente y al final se debe mostrar por pantalla el monto final (sumando el total de los 3).

*/

// La funcion go routine no retorna nada, se puede escribir en un canal

package main

import "fmt"

type Products struct {
	Name     string
	Price    float64
	Quantity int
}
type Services struct {
	Name    string
	Price   float64
	Minutes int
}
type Mantenaince struct {
	Name    string
	Price   float64
	Minutes int
}

func AddProduct(product []Products) {
	totalPriceProduct := 0.0
	for _, product := range product {
		totalPriceProduct += product.Price * float64(product.Quantity)
	}
	fmt.Println("totalPriceProduct: ", totalPriceProduct)
	//en lugar de hacer return se debe escribir en el canal
}

func AddService(service []Services) {
	totalPriceService := 0.0
	halfTime := 0
	for _, service := range service {
		halfTime += validateHour(service.Minutes)
	}
	for _, service := range service {
		totalPriceService += service.Price * float64(halfTime)
	}
	fmt.Println("totalPriceService: ", totalPriceService)
}

func validateHour(minutes int) (halfHour int) {
	if minutes <= 30 {
		return 1
	} else if minutes > 30 && minutes < 60 {
		return 2
	} else if minutes%60 != 0 {
		hours := minutes / 60
		halfHour := hours * 2
		halfHour += 1
		return halfHour
	} else {
		hours := minutes / 60
		halfHour := hours * 2
		return halfHour
	}
}

func AddMantenaince(mantenaince []Mantenaince) {
	totalPriceMantenaince := 0.0
	for _, mantenaince := range mantenaince {
		totalPriceMantenaince += mantenaince.Price
	}
	fmt.Println("totalPriceMantenaince: ", totalPriceMantenaince)
}
func main() {
	productArroz := Products{
		Name:     "Arroz",
		Price:    100.0,
		Quantity: 2,
	}
	productAzucar := Products{
		Name:     "Azucar",
		Price:    20.0,
		Quantity: 5,
	}
	servicioArreglo := Services{
		Name:    "Arreglo",
		Price:   10.0,
		Minutes: 120,
	}
	mantenimientoPC := Mantenaince{
		Name:  "Mantenimiento PC",
		Price: 10.0,
	}

	go AddProduct([]Products{productArroz, productAzucar})
	go AddService([]Services{servicioArreglo})
	go AddMantenaince([]Mantenaince{mantenimientoPC})

}
