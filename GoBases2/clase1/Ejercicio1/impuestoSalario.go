package main

import "fmt"

const (
	porcentaje1 = 0.17
	porcentaje2 = 0.10
)

func impuestoSalarial(salario float64) float64 {
	var resultado, salarioMayor float64
	if salario > 50000 && salario < 150000 {
		resultado = salario * porcentaje1
		//fmt.Println("resultado: ", resultado)
		return resultado
	} else if salario > 150000 {
		resultado = salario * porcentaje1
		salarioMayor = salario - resultado
		resultado = salarioMayor * porcentaje2
		return resultado
	}
	return 0

}

func main() {
	fmt.Println(impuestoSalarial(170000))
}
