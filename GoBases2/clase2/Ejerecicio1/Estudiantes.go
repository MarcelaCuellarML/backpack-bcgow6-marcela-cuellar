/*
Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el detalle de los datos de cada uno de ellos/as, de la siguiente
manera:
Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]
Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y que tenga un método detalle

*/

package main

import "fmt"

type Estudent struct {
	Nombre   string
	Apellido string
	DNI      string
	Fecha    string
}

func imprimirEstudiante(e []Estudent) (estudiante Estudent) {
	estudiantes := append(e)
	for i := 0; i < len(estudiantes); i++ {
		fmt.Println("Nombre: ", estudiantes[i].Nombre)
		fmt.Println("Apellido: ", estudiantes[i].Apellido)
		fmt.Println("DNI: ", estudiantes[i].DNI)
		fmt.Println("Fecha: ", estudiantes[i].Fecha)
	}
	return
}

func main() {
	estudiante1 := Estudent{
		Nombre:   "Andres",
		Apellido: "Arias",
		DNI:      "1234",
		Fecha:    "12/22",
	}
	estudiante2 := Estudent{
		Nombre:   "Maria",
		Apellido: "Perez",
		DNI:      "4321",
		Fecha:    "10/22",
	}
	imprimirEstudiante([]Estudent{estudiante1, estudiante2})
	//imprimirEstudiante(estudiante2[]	Estudent)
}
