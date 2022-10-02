/*
Varias tiendas de ecommerce necesitan realizar una funcionalidad en Go para administrar productos y retornar el valor del precio total.
Las empresas tienen 3 tipos de productos:
Pequeño, Mediano y Grande. (Se espera que sean muchos más)
Existen costos adicionales por mantener el producto en el almacén de la tienda, y costos de envío.

Sus costos adicionales son:
Pequeño: El costo del producto (sin costo adicional)
Mediano: El costo del producto + un 3% por mantenerlo en existencia en el almacén de la tienda. 0,03
Grande: El costo del producto + un 6%  por mantenimiento, y un costo adicional  por envío de $2500. 0,06

Requerimientos:
Crear una estructura “tienda” que guarde una lista de productos.
Crear una estructura “producto” que guarde el tipo de producto, nombre y precio
Crear una interface “Producto” que tenga el método “CalcularCosto”
Crear una interface “Ecommerce” que tenga los métodos “Total” y “Agregar”.
Se requiere una función “nuevoProducto” que reciba el tipo de producto, su nombre y precio y devuelva un Producto.
Se requiere una función “nuevaTienda” que devuelva un Ecommerce.
Interface Producto:
El método “CalcularCosto” debe calcular el costo adicional según el tipo de producto.
Interface Ecommerce:
  - El método “Total” debe retornar el precio total en base al costo total de los productos y los adicionales si los hubiera.
  - El método “Agregar” debe recibir un producto y añadirlo a la lista de la tienda
*/
package main

import (
	"fmt"
	"strings"
)

type Tienda struct {
	productos    []Producto
	NombreTienda string
}

type Producto struct {
	TipoProducto string
	Nombre       string
	Precio       float64
}

type ProductoInterface interface {
	CalcularCosto()
}

type TiendaInterface interface {
	Total()
	Agregar(prod Producto)
}

func (p Producto) CalcularCosto() float64 {
	tipo := p.TipoProducto
	tipoMayus := strings.ToUpper(tipo)
	switch tipoMayus {
	case "PEQUEÑO":
		return p.Precio
	case "MEDIANO":
		s := p.Precio * 0.03
		return p.Precio + s
	case "GRANDE":
		t := p.Precio * 0.06
		return t + 2500 + p.Precio
	}
	return 0
}

func (t Tienda) Total() float64 {
	costoTotal := 0.0
	for i := 0; i < len(t.productos); i++ {
		costoProducto := t.productos[i].CalcularCosto()
		//fmt.Println(t.productos[i].Nombre)
		costoTotal += costoProducto
		//fmt.Println("El costo del producto es: ", costoProducto)
	}
	fmt.Println("El costo total de la tienda es: ", costoTotal)
	return costoTotal
}

func (t Tienda) Agregar(producto Producto) (tiendaActualizada Tienda) {
	//fmt.Println("********************")
	//fmt.Println("Ingresa Agregar")
	t.productos = append(t.productos, producto)
	//fmt.Println("tienda post append -> : ", t.productos)
	return t
}

func nuevoProducto(tienda Tienda, tipoProducto, nombre string, precio float64) (tiendaActualizada Tienda) {

	producto := Producto{
		Nombre:       nombre,
		TipoProducto: tipoProducto,
		Precio:       precio,
	}
	//fmt.Println("********************")
	//fmt.Println("Ingresa metodo Nuevo Producto")
	//fmt.Println("tienda Actual -> : ", tienda.productos)
	//fmt.Println()
	return tienda.Agregar(producto)
}

func nuevaTienda(nombreTienda string) (tiendaCreada Tienda) {
	tienda := Tienda{
		productos:    []Producto{},
		NombreTienda: nombreTienda,
	}
	return tienda
}

func main() {
	// prod1 := Producto{
	// 	Nombre:       "Aventura",
	// 	TipoProducto: "PEQUEÑO",
	// 	Precio:       10.0,
	// }
	// prod2 := Producto{
	// 	Nombre:       "prod2",
	// 	TipoProducto: "mediano",
	// 	Precio:       20.0,
	// }
	// prod3 := Producto{
	// 	Nombre:       "prod3",
	// 	TipoProducto: "granDE",
	// 	Precio:       30.0,
	// }

	tienda1 := Tienda{NombreTienda: "tienda 1"}
	//productos: []Producto{prod1, prod2, prod3},}

	//tienda2 := Tienda{}

	tienda1 = nuevoProducto(tienda1, "pequeño", "Arroz", 25)
	tienda1 = nuevoProducto(tienda1, "mediano", "gaseosa", 50)
	tienda1 = nuevoProducto(tienda1, "grande", "garrafon de agua", 15)
	fmt.Println("")
	fmt.Println("---------------------")
	fmt.Println("TIENDA FINAL: ", tienda1.productos)
	fmt.Println("---------------------")

	// for i := 0; i < len(tienda1.productos); i++ {
	// 	fmt.Println(tienda1.productos[i].CalcularCosto())
	// }
	var tienda2 = nuevaTienda("tiendaNueva")
	tiendas := []Tienda{tienda1, tienda2}
	//fmt.Println("Tiendas existentes: ", tiendas)

	for i := 0; i < len(tiendas); i++ {
		//fmt.Println(tiendas[i].NombreTienda)
		fmt.Printf("productos de la tienda %s: %v", tiendas[i].NombreTienda, tiendas[i].productos)
		fmt.Println("")
	}

}
