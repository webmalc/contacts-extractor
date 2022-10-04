package validators

import "regexp"

// import "github.com/go-playground/validator/v10"

// EmailValidator struct.
type PhoneValidator struct {
}

// Validate validates contacts.
func (v *PhoneValidator) Validate(contact string) bool {
	matched, err := regexp.MatchString(
		`^\s*(?:\+?(\d{1,3}))?([-. (]*(\d{3})[-. )]*)?((\d{3})[-. ]*(\d{2,4})(?:[-.x ]*(\d+))?)$`,
		contact,
	)
	if err != nil {
		return false
	}

	return matched
}

// NewPhoneValidator returns the validator object.
func NewPhoneValidator() *PhoneValidator {
	return &PhoneValidator{}
}
