/*
Una empresa de redes sociales requiere implementar una estructura usuario con funciones que vayan agregando información a la estructura. Para optimizar y ahorrar memoria requieren que la
estructura de usuarios ocupe el mismo lugar en memoria para el main del programa y para las funciones.
La estructura debe tener los campos: Nombre, Apellido, Edad, Correo y Contraseña
Y deben implementarse las funciones:
Cambiar nombre: me permite cambiar el nombre y apellido.
Cambiar edad: me permite cambiar la edad.
Cambiar correo: me permite cambiar el correo.
Cambiar contraseña: me permite cambiar la contraseña.
*/
package main

import "fmt"

type User struct {
	name     string
	lastName string
	age      string
	email    string
	password string
}

func updateEmail(user *User, newEmail string) {
	user.email = newEmail
	fmt.Println("Dato ", user.email, "ingresado correctamente")
}

func updateLastName(user *User, newLastName string) {
	user.lastName = newLastName
	fmt.Println("Dato ", user.lastName, "ingresado correctamente")
}

func updateName(user *User, newName string) {
	user.name = newName
	fmt.Println("Dato ", user.name, "ingresado correctamente")
}

func updateAge(user *User, newAge string) {
	user.age = newAge
	fmt.Println("Dato ", user.age, "ingresado correctamente")
}

func updatePass(user *User, newPass string) {
	user.password = newPass
	fmt.Println("Dato ", user.password, "ingresado correctamente")
}

func main() {
	user1 := &User{
		name:     "John",
		lastName: "Doe",
		age:      "30",
		email:    "john@doe.com",
		password: "654321",
	}

	fmt.Println("Usuario inicial: ", user1)

	updateLastName(user1, "Galvira")
	updateName(user1, "Marcela")
	updateAge(user1, "25")
	updateEmail(user1, "marcela@gmail.com")
	updatePass(user1, "1234")

	fmt.Println("Usuario actualizado: ", user1)
}
