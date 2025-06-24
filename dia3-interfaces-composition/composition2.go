// composition2.go
package main

import "fmt"

// ------------------------------------------------------------
// 1. Componentes independientes
// ------------------------------------------------------------

// Engine representa el motor de un vehículo.
type Engine struct {
	Horsepower int  // potencia en caballos de fuerza
	Running    bool // estado del motor
}

// Start enciende el motor.
func (e *Engine) Start() {
	if e.Running {
		fmt.Println("El motor ya está en marcha.")
	} else {
		e.Running = true
		fmt.Printf("Motor arrancado (%d CV).\n", e.Horsepower)
	}
}

// Stop apaga el motor.
func (e *Engine) Stop() {
	if !e.Running {
		fmt.Println("El motor ya está apagado.")
	} else {
		e.Running = false
		fmt.Println("Motor detenido.")
	}
}

// ------------------------------------------------------------
// 2. Otro componente: Wheels
// ------------------------------------------------------------

// Wheels representa las ruedas de un vehículo.
type Wheels struct {
	Count int // número de ruedas
}

// Roll simula el rodado de las ruedas.
func (w Wheels) Roll() {
	fmt.Printf("Las %d ruedas están rodando.\n", w.Count)
}

// ------------------------------------------------------------
// 3. Composición: Car embebe Engine y Wheels
// ------------------------------------------------------------

// Car combina un Engine y un Wheels, más un modelo y año.
type Car struct {
	Engine
	Wheels
	Model string
	Year  int
}

// Drive es un método propio de Car que coordina motor + ruedas.
func (c *Car) Drive() {
	// Asegurarse de que el motor está arrancado
	c.Start() // llama a Engine.Start()
	// Hacer rodar las ruedas
	c.Roll() // llama a Wheels.Roll()
	fmt.Printf("Conduciendo %s %d...\n", c.Model, c.Year)
}

// Park detiene el motor al estacionar.
func (c *Car) Park() {
	fmt.Printf("Estacionando %s %d...\n", c.Model, c.Year)
	c.Stop() // llama a Engine.Stop()
}

// ------------------------------------------------------------
// 4. main: demostración de composición en acción
// ------------------------------------------------------------
func main() {
	// Creamos un Car con motor de 150 CV y 4 ruedas
	myCar := Car{
		Engine: Engine{Horsepower: 150},
		Wheels: Wheels{Count: 4},
		Model:  "SedanGT",
		Year:   2025,
	}

	// Conducir el auto
	myCar.Drive()
	// Estacionar
	myCar.Park()

	fmt.Println()

	// Podemos detener el motor directamente
	myCar.Stop()
}
