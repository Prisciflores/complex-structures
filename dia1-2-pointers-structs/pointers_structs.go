// pointers_structs.go
package main

import "fmt"

// ------------------------------------------------------------
// 1. Definición de un struct básico: Persona
// ------------------------------------------------------------

// Persona representa a una persona con Nombre y Edad.
type Persona struct {
	Nombre string
	Edad   int
}

// ------------------------------------------------------------
// 2. Método con receptor por valor
// ------------------------------------------------------------

// Saludar imprime un saludo usando los datos de la Persona.
// Como el receptor es por valor (p Persona), trabaja sobre una COPIA
// de la instancia original. No puede mutar los campos de la Persona
// que llamó al método.
func (p Persona) Saludar() {
	fmt.Printf("Hola, soy %s y tengo %d años. (receptor por valor)\n", p.Nombre, p.Edad)
}

// ------------------------------------------------------------
// 3. Método con receptor por puntero
// ------------------------------------------------------------

// CumplirAnios incrementa la Edad de la Persona.
// Aquí el receptor es *Persona, es decir, un puntero a la instancia original.
// Al hacer p.Edad++, estamos modificando directamente el objeto original.
func (p *Persona) CumplirAnios() {
	p.Edad++
	fmt.Printf("¡Feliz cumpleaños, %s! Ahora tienes %d años. (receptor por puntero)\n",
		p.Nombre, p.Edad)
}

// ------------------------------------------------------------
// 4. main: demostración de ambos métodos
// ------------------------------------------------------------

func main() {
	// 4.1 Instanciar una Persona por valor
	p1 := Persona{Nombre: "Priscila", Edad: 32}

	// Llamamos al método Saludar (receptor por valor)
	p1.Saludar()
	// Intentamos cumplir años: como Saludar es por valor, no afecta p1
	//p1.Ed // (sin CumplirAnios) sigue con la misma edad
	fmt.Printf("Edad antes de CumplirAnios (sin puntero): %d\n", p1.Edad)

	// 4.2 Llamar al método de puntero para mutar la instancia
	p1.CumplirAnios()
	// El método por puntero sí cambió p1.Edad
	fmt.Printf("Edad después de CumplirAnios: %d\n", p1.Edad)

	fmt.Println()

	// 4.3 Instanciar directamente un puntero a Persona
	p2 := &Persona{Nombre: "Carlos", Edad: 28}

	// Go permite llamar métodos con receptor por valor o puntero
	// tanto si tienes un valor como un puntero: la conversión es automática.
	p2.Saludar()      // receptor por valor (Go desreferencia el puntero)
	p2.CumplirAnios() // receptor por puntero

	// Confirmamos la edad actualizada
	fmt.Printf("Ahora %s tiene %d años\n", p2.Nombre, p2.Edad)
}
