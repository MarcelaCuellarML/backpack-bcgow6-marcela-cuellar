package main

import (
	"fmt"
	"os"
)

//prints del paquete fmt

func main() {
	message := "Hello World :)"
	fmt.Print(message)

	//readfile
	readFile, err := os.ReadFile("nameFile")
	if err != nil {
		println(err)
	}
	fmt.Println("%s", readFile)

	//WriteFile
	//os.WriteFile("nombreFile(se le envia una variable de mensaje)")
}

//la funcion printf permite utilizar lso verbos de accion
