package extractor

// Extractor is a contacts extractor.
type Extractor struct {
	sources        map[string]SourceExtractor
	merger         Merger
	phoneValidator Validator
	emailValidator Validator
	vkValidator    Validator
}

// Extract extracts contacts from the sources.
func (e *Extractor) Extract(
	sources []string, contacts []string,
) map[string][]string {
	// TODO: write validators
	// TODO: write merger
	// TODO: merge sources in loop -> validate sources -> return
	return map[string][]string{
		"phone":  {"111", "222"},
		"emails": {"sd", "adf"},
	}
}

// AddSource adds a source.
func (e *Extractor) AddSource(id string, source SourceExtractor) {
	e.sources[id] = source
}

// NewCSV returns the extractor object.
func NewExtractor(
	phoneValidator, emailValidator, vkValidator Validator,
) *Extractor {
	return &Extractor{
		phoneValidator: phoneValidator,
		emailValidator: emailValidator,
		vkValidator:    vkValidator,
	}
}
