package main

import "fmt"

var nombre string
var direccion string
var temperatura int
var humedad int
var presion string

func main() {
	//clima()
	saludar()
}

func clima() {
	humedad = 74
	temperatura = 16
	presion = "1029mbar"

	fmt.Println("Datos meteorologicos en Bogotá colombia")
	fmt.Println("Temperatura: ", temperatura, "°C")
	fmt.Println("Humedad: ", humedad, "%")
	fmt.Println("Presion atmosferica: ", presion)
}

func saludar() {
	nombre = "Marcela Cuellar"
	direccion = "We Work"

	fmt.Println(nombre)
	fmt.Println(direccion)
}
