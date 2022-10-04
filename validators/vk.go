package validators

import "regexp"

// VKValidator struct.
type VKValidator struct {
}

// Validate validates contacts.
func (v *VKValidator) Validate(contact string) bool {
	matched, err := regexp.MatchString(
		`^((?:id)|(?:ID))?\d{5,50}$`,
		contact,
	)
	if err != nil {
		return false
	}

	return matched
}

// NewVKValidator returns the validator object.
func NewVKValidator() *VKValidator {
	return &VKValidator{}
}
