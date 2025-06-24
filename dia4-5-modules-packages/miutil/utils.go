// miutil/utils.go
package miutil

import "fmt"

// SaludarFormateado recibe un nombre y devuelve un saludo completo.
func SaludarFormateado(nombre string) string {
	return fmt.Sprintf("¡Hola, %s! Bienvenido a Go Modules.", nombre)
}

// SumInt recibe dos enteros y devuelve su suma.
func SumInt(a, b int) int {
	return a + b
}
