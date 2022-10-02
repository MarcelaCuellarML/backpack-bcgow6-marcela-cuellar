/*
Haz lo mismo que en el ejercicio anterior pero reformulando el código para que, en reemplazo de “Error()”,  se implemente “errors.New()”.
*/
package main

import (
	"errors"
	"fmt"
)

func ObtenerError(salary int) (err error) {
	if salary < 150000 {
		err = errors.New("error: el salario ingresado no alcanza el mínimo imponible")
	}
	return
}

func main() {

	salary := 160000
	//salary := 50000
	validation := ObtenerError(salary)
	if validation != nil {
		fmt.Println(validation.Error())
	} else {
		fmt.Println("Debe pagar impuestos")
	}
}
