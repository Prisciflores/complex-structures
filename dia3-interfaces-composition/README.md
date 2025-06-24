# Día 3: Interfaces y Composición

Este directorio contiene ejemplos para comprender dos pilares de la programación en Go:

- **Interfaces**: Definir comportamientos abstractos y lograr polimorfismo.  
- **Composición**: Reutilizar código embebiendo structs (en lugar de herencia).

---

## 📁 Contenido

- `interfaces.go`           : Ejemplo de interfaz `Pagable` y dos implementaciones (`Tarjeta` y `Transferencia`).  
- `composition.go`          : Demostración de *embedding* con `PersonaSimple`, `Direccion` y `PersonaCompleta`.  
- `interfaces_simple.go`    : Ejemplo minimalista de interfaz `Animal` con `Dog` y `Cat`.  

---

## 🚀 Cómo ejecutar

1. Abre la terminal en este directorio:  
   ```bash
   cd complex-structures/dia3-interfaces-composition

2. Ejecuta
 ```bash
 go run interfaces.go
 go run composition.go
 go run interfaces_simple.go
