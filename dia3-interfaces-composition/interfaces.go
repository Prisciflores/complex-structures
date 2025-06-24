// interfaces.go
package main

import "fmt"

// ------------------------------------------------------------------------------------------------
// 1. Definición de la interfaz Pagable
// ------------------------------------------------------------------------------------------------
// Cualquier tipo que implemente Pagar(monto float64) error la satisface.
type Pagable interface {
	Pagar(monto float64) error
}

// ------------------------------------------------------------------------------------------------
// 2. Implementación: Tarjeta
// ------------------------------------------------------------------------------------------------
type Tarjeta struct {
	Titular string // Nombre del dueño
	Numero  string // Número completo de la tarjeta
}

// Pagar con receptor por puntero para poder mutar estado si fuera necesario.
func (t *Tarjeta) Pagar(monto float64) error {
	// Simulamos la operación mostrando solo los últimos 4 dígitos
	ultimos := t.Numero[len(t.Numero)-4:]
	fmt.Printf("Pagando $%.2f con tarjeta de %s (****%s)\n", monto, t.Titular, ultimos)
	return nil // nil = sin error
}

// ------------------------------------------------------------------------------------------------
// 3. Implementación: Transferencia
// ------------------------------------------------------------------------------------------------
type Transferencia struct {
	Banco  string // Nombre del banco
	Cuenta string // Número de cuenta
}

// Pagar con transferencia; devolvemos error si la cuenta no está definida.
func (tr *Transferencia) Pagar(monto float64) error {
	if tr.Cuenta == "" {
		return fmt.Errorf("cuenta inválida")
	}
	fmt.Printf("Transferencia de $%.2f a la cuenta %s en %s\n", monto, tr.Cuenta, tr.Banco)
	return nil
}

// ------------------------------------------------------------------------------------------------
// 4. Función genérica que usa Pagable
// ------------------------------------------------------------------------------------------------
func procesarPago(p Pagable, monto float64) {
	if err := p.Pagar(monto); err != nil {
		fmt.Println("🔴 Error al pagar:", err)
	} else {
		fmt.Println("✅ Pago exitoso")
	}
}

// ------------------------------------------------------------------------------------------------
// 5. main: demostración de Polimorfismo
// ------------------------------------------------------------------------------------------------
func main() {
	// 5.1 Pago con Tarjeta
	t := &Tarjeta{Titular: "Priscila", Numero: "1234567812345678"}
	procesarPago(t, 120.50)

	fmt.Println()

	// 5.2 Pago con Transferencia inválida
	tr := &Transferencia{Banco: "Banco XYZ", Cuenta: ""}
	procesarPago(tr, 200.00)

	// 5.3 Arreglamos la cuenta y probamos de nuevo
	tr.Cuenta = "000123456789"
	procesarPago(tr, 200.00)
}
