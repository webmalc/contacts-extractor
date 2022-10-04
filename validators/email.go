package validators

import "github.com/go-playground/validator/v10"

// EmailValidator struct.
type EmailValidator struct {
}

// Validate validates contacts.
func (v *EmailValidator) Validate(contact string) bool {
	validate := validator.New()
	if err := validate.Var(contact, "required,email"); err != nil {
		return false
	}

	return true
}

// NewEmailValidator returns the validator object.
func NewEmailValidator() *EmailValidator {
	return &EmailValidator{}
}
