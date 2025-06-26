// panic_recover.go
package main

import "fmt"

// ------------------------------------------------------------
// 1) Función que puede panicar
// ------------------------------------------------------------
func puedePanicar(divisor int) {
	if divisor == 0 {
		panic("¡división por cero detectada!")
	}
	fmt.Println("Resultado seguro:", 100/divisor)
}

// ------------------------------------------------------------
// 2) Wrapper que recupera el panic
// ------------------------------------------------------------
func safeEjecutar(divisor int) {
	// defer se ejecuta al salir de la función, incluso tras un panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("RECUPERADO del panic:", r)
		}
	}()

	// Aquí puede ocurrir un panic
	puedePanicar(divisor)

	// Si no panic, continúa normalmente
	fmt.Println("safeEjecutar completado sin panic.")
}

// ------------------------------------------------------------
// 3) main: demostración de defer, panic y recover
// ------------------------------------------------------------
func main() {
	fmt.Println("=== Llamada segura con divisor=5 ===")
	safeEjecutar(5)

	fmt.Println("\n=== Llamada segura con divisor=0 ===")
	safeEjecutar(0)

	fmt.Println("\n=== main continúa normalmente ===")
}
