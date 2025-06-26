// validator/validate.go
// Paquete validator: valida contraseñas según reglas de seguridad.

package validator

import (
	"errors"  // Para crear errores sentinel
	"unicode" // Para analizar caracteres Unicode (mayúsculas, dígitos, símbolos)
)

// Errores sentinel exportados para comparar con errors.Is()
// --------------------------------------------------------
var (
	ErrTooShort      = errors.New("la contraseña es demasiado corta (mínimo 8 caracteres)")
	ErrNoUpper       = errors.New("debe contener al menos una letra mayúscula")
	ErrNoLower       = errors.New("debe contener al menos una letra minúscula")
	ErrNoDigit       = errors.New("debe contener al menos un dígito")
	ErrNoSpecialChar = errors.New("debe contener al menos un carácter especial (!@#$...)")
)

// Validate revisa la cadena pw y devuelve un slice de errores
// correspondientes a cada regla que no cumpla.
// ------------------------------------------------------------
func Validate(pw string) []error {
	var errs []error

	// 1) Regla de longitud mínima
	if len(pw) < 8 {
		errs = append(errs, ErrTooShort)
	}

	// 2) Variables para rastrear cada tipo de carácter
	var hasUpper, hasLower, hasDigit, hasSpecial bool

	// 3) Iteramos cada rune de la contraseña
	for _, r := range pw {
		switch {
		case unicode.IsUpper(r):
			hasUpper = true
		case unicode.IsLower(r):
			hasLower = true
		case unicode.IsDigit(r):
			hasDigit = true
		case unicode.IsPunct(r) || unicode.IsSymbol(r):
			hasSpecial = true
		}
	}

	// 4) Si falta algún tipo, añadimos el error correspondiente
	if !hasUpper {
		errs = append(errs, ErrNoUpper)
	}
	if !hasLower {
		errs = append(errs, ErrNoLower)
	}
	if !hasDigit {
		errs = append(errs, ErrNoDigit)
	}
	if !hasSpecial {
		errs = append(errs, ErrNoSpecialChar)
	}

	return errs
}
