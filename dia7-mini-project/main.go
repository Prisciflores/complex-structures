// main.go
// CLI minimalista que pide una contraseña, la valida y muestra resultados.

package main

import (
	"bufio"                       // Para leer entrada por consola
	"dia7-mini-project/validator" // Importa el paquete de validación de contraseñas
	"fmt"                         // Para imprimir mensajes
	"os"                          // Para acceder a stdin
	"strings"                     // Para manipular cadenas
	// Nota: Asegúrate de que el paquete validator esté en la ruta correcta
)

func main() {
	// 1. Pedir contraseña al usuario
	fmt.Print("Ingresa una contraseña para validar: ")
	reader := bufio.NewReader(os.Stdin)
	pw, _ := reader.ReadString('\n') // Leer hasta salto de línea
	pw = strings.TrimSpace(pw)       // Quitar espacios y '\n'

	// 2. Validar usando el paquete validator
	errs := validator.Validate(pw)

	// 3. Mostrar resultados
	if len(errs) == 0 {
		fmt.Println("✅ Contraseña válida")
		return
	}

	fmt.Println("❌ La contraseña NO es válida:")
	for _, err := range errs {
		// err.Error() devuelve el mensaje definido en los sentinels
		fmt.Printf("  - %s\n", err.Error())
	}
}
