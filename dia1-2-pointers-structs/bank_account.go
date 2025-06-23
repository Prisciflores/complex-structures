// bank_account.go
// Ejemplo de CuentaBancaria que muestra el uso de receptores por valor y por puntero.

package main

import "fmt"

// ------------------------------------------------------------
// 1. Definición del struct CuentaBancaria
// ------------------------------------------------------------
// Representa una cuenta con un titular y su saldo actual.
type CuentaBancaria struct {
	Titular string  // Nombre del dueño de la cuenta
	Saldo   float64 // Saldo disponible en la cuenta
}

// ------------------------------------------------------------
// 2. Método receptor por valor: MostrarSaldo
// ------------------------------------------------------------
// Este método imprime el saldo de la cuenta.
// Al ser receptor por valor, no puede modificar el saldo original.
func (c CuentaBancaria) MostrarSaldo() {
	fmt.Printf("Cuenta de %s → Saldo: $%.2f\n", c.Titular, c.Saldo)
}

// ------------------------------------------------------------
// 3. Método receptor por puntero: Depositar
// ------------------------------------------------------------
// Este método recibe un puntero a CuentaBancaria, lo que
// le permite modificar el campo Saldo de la instancia original.
func (c *CuentaBancaria) Depositar(monto float64) {
	// Validación: el monto debe ser positivo
	if monto <= 0 {
		fmt.Println("🔴 El monto a depositar debe ser positivo")
		return
	}
	// Sumar el monto al saldo existente
	c.Saldo += monto
	fmt.Printf("🟢 Depositaste $%.2f. Nuevo saldo: $%.2f\n", monto, c.Saldo)
}

// ------------------------------------------------------------
// 4. Método receptor por puntero: Retirar
// ------------------------------------------------------------
// Similar a Depositar, este método puede modificar el Saldo
// original porque recibe un puntero.
func (c *CuentaBancaria) Retirar(monto float64) {
	// Validación: el monto debe ser positivo
	if monto <= 0 {
		fmt.Println("🔴 El monto a retirar debe ser positivo")
		return
	}
	// Validación: no retirar más de lo que hay
	if monto > c.Saldo {
		fmt.Println("🔴 Fondos insuficientes")
		return
	}
	// Restar el monto del saldo
	c.Saldo -= monto
	fmt.Printf("🟠 Retiraste $%.2f. Nuevo saldo: $%.2f\n", monto, c.Saldo)
}

// ------------------------------------------------------------
// 5. Función main: demostración de los métodos
// ------------------------------------------------------------
func main() {
	// 5.1 Crear una cuenta con saldo inicial
	cuenta := CuentaBancaria{Titular: "Priscila", Saldo: 100.00}

	// 5.2 Mostrar saldo (receptor por valor, no muta nada)
	cuenta.MostrarSaldo()

	// 5.3 Intentar retirar más de lo permitido (debe fallar)
	cuenta.Retirar(150.00)
	cuenta.MostrarSaldo() // Saldo sigue siendo 100.00

	// 5.4 Depositar un monto válido
	cuenta.Depositar(50.00)
	cuenta.MostrarSaldo() // Saldo ahora es 150.00

	// 5.5 Retirar un monto válido
	cuenta.Retirar(25.00)
	cuenta.MostrarSaldo() // Saldo ahora es 125.00

	// 5.6 Usar un puntero a la cuenta para depositar
	ptr := &cuenta
	// Go permite llamar métodos por valor o puntero indistintamente:
	ptr.Depositar(10.00) // receptor por puntero, muta saldo
	ptr.MostrarSaldo()   // receptor por valor, solo imprime
}
