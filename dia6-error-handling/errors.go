// errors.go
package main

import (
	"errors"
	"fmt"
)

// ------------------------------------------------------------
// 1) Errores sentinel y variables de error
// ------------------------------------------------------------
var ErrNegativeInput = errors.New("input no puede ser negativo")

// ------------------------------------------------------------
// 2) Función que devuelve un error si la entrada es inválida
// ------------------------------------------------------------
func RaizCuadrada(n float64) (float64, error) {
	if n < 0 {
		// Retornamos el error sentinel
		return 0, ErrNegativeInput
	}
	return n * n, nil
}

// ------------------------------------------------------------
// 3) Error personalizado mediante struct
// ------------------------------------------------------------
type ErrDivZero struct {
	Dividendo, Divisor int
}

func (e ErrDivZero) Error() string {
	return fmt.Sprintf("no se puede dividir %d entre %d", e.Dividendo, e.Divisor)
}

func Dividir(a, b int) (int, error) {
	if b == 0 {
		// Retornamos un ErrDivZero con contexto
		return 0, ErrDivZero{a, b}
	}
	return a / b, nil
}

// ------------------------------------------------------------
// 4) main: demo de manejo de errores
// ------------------------------------------------------------
func main() {
	// Caso válido
	if r, err := RaizCuadrada(4); err != nil {
		fmt.Println("Error en RaizCuadrada:", err)
	} else {
		fmt.Println("Raíz cuadrada de 4 =", r)
	}

	// Caso inválido (negativo)
	if _, err := RaizCuadrada(-1); err != nil {
		// Podemos comparar con ErrNegativeInput
		if errors.Is(err, ErrNegativeInput) {
			fmt.Println("¡Recibimos un número negativo!")
		} else {
			fmt.Println("Otro error:", err)
		}
	}

	// División válida
	if q, err := Dividir(10, 2); err != nil {
		fmt.Println("Error en Dividir:", err)
	} else {
		fmt.Println("10 / 2 =", q)
	}

	// División por cero
	if _, err := Dividir(5, 0); err != nil {
		// Tipo assertion para obtener información adicional
		if de, ok := err.(ErrDivZero); ok {
			fmt.Printf("División inválida: %s\n", de.Error())
		} else {
			fmt.Println("Otro error:", err)
		}
	}
}
