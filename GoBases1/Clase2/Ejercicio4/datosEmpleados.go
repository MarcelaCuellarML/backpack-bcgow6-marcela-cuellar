package main

import "fmt"

func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	var mayores int
	//Primer punto: Saber cuántos de sus empleados son mayores de 21 años.
	for key, element := range employees {

		//var empleadosMayores []string

		if element > 21 {
			println("nombre: ", key)
			mayores++
		}
	}
	fmt.Println("El numero de empleados mayores de 21 años es: ", mayores)
	//Segundo punto: Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.

	//employees["Federico"] = 25

	//Tercer punto: Eliminar a Pedro del mapa.

	//delete(employees, "Pedro")
}
