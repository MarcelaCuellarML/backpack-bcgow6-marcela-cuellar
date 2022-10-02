package main

import (
	"fmt"
	"time"
)

func routines(i int, ch chan int) {
	fmt.Println(i, " - inicia")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(i, " - termina")
	ch <- i
}

func main() {
	//definicion de canales
	channel := make(chan int)
	routines(1, channel)
}
