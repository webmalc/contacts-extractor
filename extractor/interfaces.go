package extractor

// InfoLogger logs info.
type InfoLogger interface {
	Info(args ...interface{})
}

// SourceExtractor extracts contacts from the source.
type SourceExtractor interface {
	Extract(contacts []string) map[string][]string
}

// Validator validates contacts.
type Validator interface {
	Validate(contact string) bool
}

// Merger merge the results.
type Merger interface {
	merge(map[string][]string) map[string][]string
}
