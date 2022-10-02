/*
En tu función “main”, define una variable llamada “salary” y asignarle un valor de tipo “int”.
Crea un error personalizado con un struct que implemente “Error()” con el mensaje “error: el salario ingresado no alcanza el mínimo imponible" y lánzalo en caso de que “salary” sea menor a
150.000. Caso contrario, imprime por consola el mensaje “Debe pagar impuesto”.
*/
package main

import "fmt"

type MyError struct {
	Message string
}

func (err *MyError) Error() string {
	return err.Message
}

func GenerateError(salary int) (err error) {
	if salary < 150000 {
		return &MyError{
			Message: "error: el salario ingresado no alcanza el mínimo imponible",
		}
	}
	return nil
}

func main() {

	//salary := 160000
	salary := 50000
	validation := GenerateError(salary)
	if validation != nil {
		fmt.Println(validation.Error())
	} else {
		fmt.Println("Debe pagar impuestos")
	}

}
