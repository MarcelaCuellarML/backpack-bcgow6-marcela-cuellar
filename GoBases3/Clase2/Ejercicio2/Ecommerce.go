/*
Una importante empresa de ventas web necesita agregar una funcionalidad para agregar productos a los usuarios. Para ello requieren que tanto los usuarios como los productos tengan la misma
dirección de memoria en el main del programa como en las funciones.
Se necesitan las estructuras:
Usuario: Nombre, Apellido, Correo, Productos (array de productos).
Producto: Nombre, precio, cantidad.
Se requieren las funciones:
Nuevo producto: recibe nombre y precio, y retorna un producto.
Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el producto al usuario.
Borrar productos: recibe un usuario, borra los productos del usuario.

*/

package main

import "fmt"

type Product struct {
	name  string
	price float64
	cant  int
}

type User struct {
	name     string
	lastName string
	email    string
	products []Product
}

func NewProduct(name string, price float64, cant int) (productoNuevo []Product) {
	var product Product
	product.name = name
	product.price = price
	product.cant = cant
	fmt.Println(product)
	return
}

func AddProduct(user *User, products []Product) {
	user.products = append(user.products, products...)
}

// Esto es un metodo a la estructura (lleva parentesis con la estructura referenciada)
func (user *User) DeleteProduct() {
	user.products = []Product{}
}

func main() {

	userDemo := &User{
		name:     "Marcela",
		lastName: "Cuellar",
		email:    "nombre.candidato@example.com",
	}

	produtoGuardado := NewProduct("Arroz", 10500, 10)
	NewProduct("Azucar", 8000, 30)
	AddProduct(userDemo, produtoGuardado)
	userDemo.DeleteProduct()
}
