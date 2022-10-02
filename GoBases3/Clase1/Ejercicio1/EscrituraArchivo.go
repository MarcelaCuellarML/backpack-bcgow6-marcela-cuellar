//Una empresa que se encarga de vender productos de limpieza necesita:
//Implementar una funcionalidad para guardar un archivo de texto, con la información de productos comprados, separados por punto y coma (csv).
//Debe tener el id del producto, precio y la cantidad.
//Estos valores pueden ser hardcodeados o escritos en duro en una variable.

package main

import (
	"fmt"
	"os"
)

type Producto struct {
	id       int
	price    float64
	quantity int
}

func main() {
	//id := 1
	//price := 10.0
	//cant := 2
	//insertProduct := fmt.Sprintf("%v ; %f ; %v", id, price, cant)
	//products := []byte(insertProduct)
	//os.WriteFile("./productos.txt", products, 777)
	producto := Producto{
		id:       1,
		price:    10.0,
		quantity: 2,
	}
	//strings :=
	prueba := fmt.Sprintf("%v, %f, %v", producto.id, producto.price, producto.quantity)
	products := []byte(prueba)
	os.WriteFile("prueba.txt", products, 777)
}
