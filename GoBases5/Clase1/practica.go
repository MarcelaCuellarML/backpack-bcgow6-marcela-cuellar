//Funciones diferidas, las lee aun cuando existe un panic, las ejecuta despues de terminar la ejecucion del main. Para ejecutar la funcion diferida para que se ejecute despues de un panic
//debe llevar el recover

//Context: se usa para el traslado de informacion
//deadline: limite de ejecucion
//done: define cuando se cumple o no una propiedad
//error: mensaje de error
//Value: almacenar y obtener un dato

package main

import (
	"context"
	"fmt"
)

func main() {
	/*
	   	defer func(){
	   		if err := recover(); err!= nil{
	               fmt.Println("Esta linea la muestra si se genera un paniic y se termina la ejecucion", err)
	               os.Exit(1)
	           }
	   	}
	*/

	ctx := context.Background()

	ctx2 := context.WithValue(ctx, "Clave", "Valor")

	fmt.Println(ctx, ctx2)

}
