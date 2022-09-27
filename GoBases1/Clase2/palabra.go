package main

import (
	"fmt"
	"strings"
)

func main() {

	palabra := "Marcela"
	fmt.Println("Extension de la palabra: ", len(palabra))

	characters := strings.Split(palabra, "")
	fmt.Print(characters, " , ")

	/*
	   var slice = []rune(palabra)

	   	for i := 0; i <= len(palabra); i++ {
	   		fmt.Print(slice[i], " , ")
	   	}
	*/
}
