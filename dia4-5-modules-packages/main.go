// main.go
package main

import (
	"complex-structures/dia4-5-modules-packages/miutil"
	"fmt"
)

func main() {
	// 1) Saludo
	saludo := miutil.SaludarFormateado("Priscila")
	fmt.Println(saludo)

	// 2) Suma de enteros
	resultado := miutil.SumInt(7, 8)
	fmt.Printf("7 + 8 = %d\n", resultado)
}
