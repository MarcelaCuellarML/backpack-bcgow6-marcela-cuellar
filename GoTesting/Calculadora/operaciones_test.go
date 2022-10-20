package calculadora

// Se importa el package testing
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 5
	num2 := 3
	resultadoEsperado := 2

	// Se ejecuta el test
	resultado := Restar(num1, num2)

	// Se validan los resultados
	if resultado != resultadoEsperado {
		t.Errorf("Funcion restar() arrojo el resultado = %v, pero el esperado es  %v", resultado, resultadoEsperado)
	}

	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
}

func TestRestarBad(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 5
	num2 := 2
	resultadoEsperado := 2

	// Se ejecuta el test
	resultado := Restar(num1, num2)

	// Se validan los resultados

	assert.NotEqual(t, resultadoEsperado, resultado, "no deben ser iguales")
}

func TestRestarNegative(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := -5
	num2 := 2
	resultadoEsperado := -7

	// Se ejecuta el test
	resultado := Restar(num1, num2)

	// Se validan los resultados

	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
}

func RestarNegativeBad(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := -5
	num2 := 2
	resultadoEsperado := -7

	// Se ejecuta el test
	resultado := Restar(num1, num2)

	// Se validan los resultados

	if resultado != resultadoEsperado {
		t.Errorf("Funcion restar() arrojo el resultado = %v, pero el esperado es  %v", resultado, resultadoEsperado)
	}
	assert.NotEqual(t, resultadoEsperado, resultado, "no deben ser iguales")
}

func TestDividir(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 10
	num2 := 2
	resultadoEsperado := 5

	// Se ejecuta el test
	resultado, err := Dividir(num1, num2)
	if err != nil {
		t.Errorf(err.Error())
	}
	// Se validan los resultados
	if resultado != resultadoEsperado {
		t.Errorf("Funcion restar() arrojo el resultado = %v, pero el esperado es  %v", resultado, resultadoEsperado)
	}

	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
}

func TestDividirBad(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 10
	num2 := 0
	resultadoEsperado := 4

	// Se ejecuta el test
	resultado, err := Dividir(num1, num2)
	if err != nil {
		t.Errorf(err.Error())
	}
	// Se validan los resultados
	assert.NotEqual(t, resultadoEsperado, resultado, "no deben ser iguales")
}
