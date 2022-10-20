package calculadora

import "errors"

// Función que recibe dos enteros y retorna la suma resultante
func Sumar(num1, num2 int) int {
	return num1 + num2
}

// Función que recibe dos enteros y retorna la resta o diferencia resultante
func Restar(num1, num2 int) int {
	return num1 - num2
}

// Función que recibe dos enteros y retorna la resta o diferencia resultante
func Dividir(num1, den int) (int, error) {
	err := errors.New("El denominador no puede ser 0")
	if den == 0 {
		return 0, err
	}
	return num1 / den, nil
}
