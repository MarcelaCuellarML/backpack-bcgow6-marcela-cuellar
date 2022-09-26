package main

import (
	"fmt"
)

func main() {
	var edad, edadMinima int
	var empleado bool
	var sueldo, sueldoIntereses, antiguedadEmpleoMinima, antiguedadEmpleo float64

	edad = 24
	edadMinima = 22
	antiguedadEmpleo = 1.5
	antiguedadEmpleoMinima = 1
	empleado = true
	sueldoIntereses = 100.000
	sueldo = 80.000

	if edad < edadMinima {
		fmt.Println("No cumple con la edad estipulada para acceder al prestamo")
	} else if !empleado {
		fmt.Println("Debe contar con un trabajo estable para poder acceder a un prestamo")
	} else if antiguedadEmpleo < antiguedadEmpleoMinima {
		fmt.Println("Debe tener como minimo un año de antiguedad en su empleo para acceder al prestamo")
	} else {
		if sueldo >= sueldoIntereses {
			fmt.Println("Su prestamo ha sido aprovado, y por su rango salarial no se le cobraran intereses")
		} else {
			fmt.Println("Su prestamo ha sido aprovado")
		}
	}
}
