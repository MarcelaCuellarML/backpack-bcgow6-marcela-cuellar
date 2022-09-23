package main

import "fmt"

var temperatura int
var humedad int
var presion string

func main() {
	humedad = 74
	temperatura = 16
	presion = "1029mbar"

	fmt.Println("Datos meteorologicos en Bogotá colombia")
	fmt.Println("Temperatura: ", temperatura, "°C")
	fmt.Println("Humedad: ", humedad, "%")
	fmt.Println("Presion atmosferica: ", presion)
}
