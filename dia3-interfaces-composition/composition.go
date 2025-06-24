// composition.go
package main

import "fmt"

// ------------------------------------------------------------------------------------------------
// 1. Structs base
// ------------------------------------------------------------------------------------------------

// PersonaSimple tiene datos básicos y un método Saludar.
type PersonaSimple struct {
	Nombre string
	Edad   int
}

// Saludar es método por valor; no muta el struct.
func (p PersonaSimple) Saludar() {
	fmt.Printf("Hola, soy %s y tengo %d años.\n", p.Nombre, p.Edad)
}

// Direccion guarda datos de ubicación y un método Info.
type Direccion struct {
	Calle  string
	Ciudad string
	CP     string
}

// Info de Direccion devuelve una cadena formateada.
func (d Direccion) Info() string {
	return fmt.Sprintf("Dirección: %s, %s, CP %s", d.Calle, d.Ciudad, d.CP)
}

// ------------------------------------------------------------------------------------------------
// 2. Composición: embebiendo structs
// ------------------------------------------------------------------------------------------------

// PersonaCompleta embebe PersonaSimple y Direccion, y añade Email.
// Gracias al embedding, hereda Saludar() e Info() de ambos.
type PersonaCompleta struct {
	PersonaSimple
	Direccion
	Email string
}

// Contactar es un método propio que usa campos de los embebidos.
func (pc PersonaCompleta) Contactar() {
	fmt.Printf("Enviando email a %s: Hola %s, vives en %s. ¡Saludos!\n",
		pc.Email, pc.Nombre, pc.Ciudad)
}

// ------------------------------------------------------------------------------------------------
// 3. Sobreescribimos Info para evitar ambigüedad
// ------------------------------------------------------------------------------------------------
func (pc PersonaCompleta) Info() string {
	// Llamamos a Info de Direccion y combinamos con Email
	return fmt.Sprintf("%s | Email: %s", pc.Direccion.Info(), pc.Email)
}

// ------------------------------------------------------------------------------------------------
// 4. main: demostración de campos y métodos heredados
// ------------------------------------------------------------------------------------------------
func main() {
	// Instanciamos PersonaCompleta con datos completos
	pc := PersonaCompleta{
		PersonaSimple: PersonaSimple{Nombre: "Carlos", Edad: 30},
		Direccion:     Direccion{Calle: "C. Falsa 123", Ciudad: "Springfield", CP: "12345"},
		Email:         "carlos@ejemplo.com",
	}

	// Podemos usar Saludar() directamente
	pc.Saludar() // viene de PersonaSimple

	// Y Contactar() propio
	pc.Contactar()

	// Accedemos a campos embebidos sin calificar
	fmt.Printf("%s vive en %s, CP %s\n", pc.Nombre, pc.Ciudad, pc.CP)

	// Y usamos la versión personalizada de Info()
	fmt.Println(pc.Info())
}
