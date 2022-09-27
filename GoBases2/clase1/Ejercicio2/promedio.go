// Un colegio necesita calcular el promedio (por alumno) de sus calificaciones. Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio
// y un error en caso que uno de los números ingresados sea negativo
package main

import (
	"errors"
	"fmt"
)

func average(calificaciones ...float64) (result float64, err error) {
	var longitud int
	longitudFloat := 0.0
	longitud = len(calificaciones)
	for i := 0; i < longitud; i++ {
		if calificaciones[i] < 0 {
			return 0, errors.New("Se ha ingresado un valor negativo, no es posible generar el promedio")
		}
	}
	for _, value := range calificaciones {
		result += value
	}
	longitudFloat = float64(longitud)
	result = result / longitudFloat
	fmt.Println("resultado antes de salir del for: ", result)
	return result, nil
}
func main() {

	fmt.Println(average(5, 3.5, 2, 4, 2.3, 4.8))
}
