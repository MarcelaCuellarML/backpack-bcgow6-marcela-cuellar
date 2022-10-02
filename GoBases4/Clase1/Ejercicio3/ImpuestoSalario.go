/*
Repite el proceso anterior, pero ahora implementando “fmt.Errorf()”, para que el mensaje de error reciba por parámetro el valor de “salary” indicando que no alcanza el mínimo imponible (el
	mensaje mostrado por consola deberá decir: “error: el mínimo imponible es de 150.000 y el salario ingresado es de: [salary]”, siendo [salary] el valor de tipo int pasado por parámetro).
*/

package main

import (
	"fmt"
)

func ObtenerError(salary int) (err error) {
	if salary < 150000 {
		err = fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %v", salary)
	}
	return
}

func main() {

	//salary := 160000
	salary := 149999
	validation := ObtenerError(salary)
	if validation != nil {
		fmt.Println(validation.Error())
	} else {
		fmt.Println("Debe pagar impuestos")
	}

}
