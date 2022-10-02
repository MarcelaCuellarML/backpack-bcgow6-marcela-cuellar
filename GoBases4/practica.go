package main

import (
	"errors"
	"fmt"
	"os"
)

/*
	type MyError struct {
		statusCode int
		Message    string
	}

	func (err *MyError) Error() string {
		return fmt.Sprintf("%d %s", err.statusCode, err.Message)
	}

	func ObtenerError (status int)(code int err error){
		if status >= 400{
			return 500, &MyError{
				statusCode: code,
	            Message:    "Sucedio un error",
			}
		}
		return 200, nil
	}
*/
func division(num1, num2 float64) (result float64, err error) {
	if num2 == 0 {
		err = errors.New("no se divide por cero")
		return
	}
	result = num1 / num2
	return
}

func main() {
	result, err := division(10, 2)
	if err != nil {
		fmt.Printf("Ha ocurrido un error inesperado %v, \n", err)
		os.Exit(1)
	}
	fmt.Println("Resultado: ", result)
	/*
		status, err := ObtenerError(200)
		//Los errores se pueden contener unos a otros

			err2 := fmt.Errorf("Second error: %w", err1)
			fmt.Println(err2)
	*/

}
