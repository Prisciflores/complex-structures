// interfaces_simple.go
package main

import "fmt"

// Animal es una interfaz que define el comportamiento Speak,
// que devuelve el sonido que hace el animal.
type Animal interface {
	Speak() string
}

// Dog implementa Animal con receptor por valor.
type Dog struct{}

// Speak de Dog devuelve "Woof!".
func (d Dog) Speak() string {
	return "Woof!"
}

// Cat implementa Animal con receptor por puntero.
type Cat struct {
	Name string
}

// Speak de Cat devuelve "Meow!" incluyendo su nombre.
func (c *Cat) Speak() string {
	return "Meow! Soy " + c.Name
}

// greet recibe cualquier Animal e imprime su sonido.
func greet(a Animal) {
	fmt.Println("El animal dice:", a.Speak())
}

func main() {
	// Instanciamos un Dog (valor) y un Cat (puntero)
	dog := Dog{}
	cat := &Cat{Name: "Misu"}

	// Ambos satisfacen la interfaz Animal
	greet(dog) // El animal dice: Woof!
	greet(cat) // El animal dice: Meow! Soy Misu

	// También podemos guardarlos en un slice de Animal
	animales := []Animal{dog, cat}
	fmt.Println("\nRecorriendo slice de Animal:")
	for i, a := range animales {
		fmt.Printf("Animal #%d: %s\n", i+1, a.Speak())
	}
}
