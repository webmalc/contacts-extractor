package extractor

// SourceExtractor extracts contacts from the source.
type sourceExtractor interface {
	Extract(contacts []string) map[string][]string
}

// Validator validates contacts.
type validator interface {
	Validate(contact string) bool
}

// Merger merge the results.
type mergerer interface {
	merge(result, contacts map[string][]string) map[string][]string
}
