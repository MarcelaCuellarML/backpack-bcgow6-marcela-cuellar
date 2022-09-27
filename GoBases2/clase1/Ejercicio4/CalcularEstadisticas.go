// Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones de los alumnos de un curso, requiriendo calcular los valores mínimo, máximo y promedio
// de sus calificaciones.
// Se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio) y que devuelva otra función ( y un mensaje en caso que el cálculo no esté definido)
// que se le puede pasar una cantidad N de enteros y devuelva el cálculo que se indicó en la función anterior
package main

import (
	"fmt"
)

func selectOperation(funcion string, notas ...float64) {
	switch funcion {
	case "minimum":
		fmt.Println(minimum(notas...))
	case "average":
		fmt.Println(average(notas...))
	case "maximum":
		fmt.Println(maximun(notas...))
	default:
		fmt.Println("La operacion ingresada no existe")
	}
}

func average(notas ...float64) (result float64) {
	var totalNotas float64
	for i := 0; i < len(notas); i++ {
		totalNotas += notas[i]
	}
	return totalNotas / float64(len(notas))
}
func minimum(notas ...float64) (result float64) {
	var notaMinima float64
	notaMinima = notas[0]
	for i := 0; i < len(notas); i++ {
		if notaMinima > notas[i] {
			notaMinima = notas[i]
		}
	}
	return notaMinima
}
func maximun(notas ...float64) (result float64) {
	var notaMaxima float64
	notaMaxima = 0.0
	for i := 0; i < len(notas); i++ {
		if notaMaxima < notas[i] {
			notaMaxima = notas[i]
		}
	}
	return notaMaxima
}
func main() {
	selectOperation("average", 5, 4, 3.5)
	selectOperation("minimum", 5, 2.5, 3)
	selectOperation("maximum", 5, 1, 4)
}
