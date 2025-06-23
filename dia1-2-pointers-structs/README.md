# Día 1–2: Punteros y Structs

Este directorio contiene ejemplos prácticos para entender la diferencia entre **receptores por valor** y **receptores por puntero** en Go.

## Archivos

- `pointers_structs.go`  
  - Define el struct `Persona`  
  - Método `Saludar()` con **receptor por valor** (trabaja sobre copia)  
  - Método `CumplirAnios()` con **receptor por puntero** (modifica el original)  

- `bank_account.go`  
  - Define el struct `CuentaBancaria`  
  - Método `MostrarSaldo()` con **receptor por valor** (solo lee)  
  - Métodos `Depositar()` y `Retirar()` con **receptor por puntero** (mutan el saldo)  

## Cómo ejecutar

Abre una terminal y, desde este directorio, corre:

```bash
go run pointers_structs.go
go run bank_account.go
