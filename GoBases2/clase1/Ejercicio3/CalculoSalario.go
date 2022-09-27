// Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.
// Si es categoría C, su salario es de $1.000 por hora
// Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
// Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual
// Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría, y que devuelva su salario.
package main

import "fmt"

func CalculateSalary(minutes int, category string) float64 {
	salaryCatA := 3000
	salaryCatB := 1500
	salaryCatC := 1000
	var salary float64

	hours := minutes / 60

	switch category {
	case "A":
		salary = float64(hours) * float64(salaryCatA)
		salary += salary * 0.5
		return salary
	case "B":
		salary = float64(hours) * float64(salaryCatB)
		salary += salary * 0.2
		return salary
	case "C":
		salary = float64(hours) * float64(salaryCatC)
		return salary
	}

	return 0
}
func main() {
	fmt.Println(CalculateSalary(60, "A"))
	fmt.Println(CalculateSalary(60, "B"))
	fmt.Println(CalculateSalary(60, "C"))
}
