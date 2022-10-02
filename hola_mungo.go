package main

import "fmt"

func main() {
	var saludo string

	saludo = "Hola Bootcanmp!!!"
	fmt.Print(saludo)
	fmt.Println(validateHour(35))

}

func validateHour(minutes int) (halfHour int) {
	if minutes <= 30 {
		return 1
	} else if minutes > 30 && minutes < 60 {
		return 2
	} else if minutes%60 != 0 {
		hours := minutes / 60
		halfHour := hours * 2
		halfHour += 1
		return halfHour
	} else {
		hours := minutes / 60
		halfHour := hours * 2
		return halfHour
	}
}
