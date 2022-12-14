/*
La misma empresa necesita leer el archivo almacenado, para ello requiere que: se imprima por pantalla mostrando los valores tabulados, con un título (tabulado a la
izquierda para el ID y a la derecha para el Precio y Cantidad), el precio, la cantidad y abajo del precio se debe visualizar el total (Sumando precio por cantidad)
Ejemplo:
ID                            Precio  Cantidad
111223                      30012.00         1
444321                    1000000.00         4
434321                         50.50         1
                          4030062.50

*/

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	readFile, err := os.ReadFile("../Ejercicio1/prueba.txt")
	fmt.Println("readfile: ", readFile)
	prue := strings.SplitAfter(string(readFile), "\n")
	fmt.Println("largo: ", len(prue))
	fmt.Println("prueba: ", prue)
	for i := 0; i < len(prue); i++ {
		fmt.Println(prue[i])
	}
	if err != nil {
		println(err)
	}
	//fmt.Printf("ID \t Precio \t Cantidad \n %s", string(readFile[:]))

}
