// validator/validate_test.go
// Tests unitarios para el paquete validator.

package validator

import "testing"

// TestValidate cubre varios escenarios de validación de contraseña.
// ---------------------------------------------------------------
func TestValidate(t *testing.T) {
	// Casos de prueba: contraseña y slice esperado de errores
	cases := []struct {
		pw       string
		wantErrs []error
	}{
		// Muy corta: sólo ErrTooShort
		{"Ab1!", []error{ErrTooShort}},

		// Larga pero sin mayúsculas, sin dígitos ni especiales
		{"abcdefgh", []error{ErrNoUpper, ErrNoDigit, ErrNoSpecialChar}},

		// Con dígitos y especiales, pero sin minúsculas
		{"ABCDEF1!", []error{ErrNoLower}},

		// Con mayúscula y minúscula, pero sin dígitos ni especiales
		{"Abcdefgh", []error{ErrNoDigit, ErrNoSpecialChar}},

		// Cumple todas las reglas → ningún error
		{"Abcdef1!", nil},
	}

	for _, c := range cases {
		errs := Validate(c.pw)
		// Verificamos número de errores devueltos
		if len(errs) != len(c.wantErrs) {
			t.Errorf("Validate(%q) = %v; want %v", c.pw, errs, c.wantErrs)
			continue
		}
		// Comparamos cada error por igualdad de la variable sentinel
		for i, err := range errs {
			if err != c.wantErrs[i] {
				t.Errorf("Validate(%q)[%d] = %v; want %v", c.pw, i, err, c.wantErrs[i])
			}
		}
	}
}
