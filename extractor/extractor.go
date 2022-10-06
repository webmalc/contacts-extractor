package extractor

import (
	"github.com/webmalc/contacts-extractor/extractor/extractors"
	"golang.org/x/exp/slices"
)

// Extractor is a contacts extractor.
type Extractor struct {
	sources map[string]sourceExtractor
	merger  mergerer
}

// Extract extracts contacts from the sources.
func (e *Extractor) Extract(
	sources []string, contacts []string,
) map[string][]string {
	result := map[string][]string{}

	for id, source := range e.sources {
		if !slices.Contains(sources, id) {
			continue
		}
		result = e.merger.merge(result, source.Extract(contacts))
	}

	return result
}

// AddSource adds a source.
func (e *Extractor) AddSource(id string, source sourceExtractor) {
	e.sources[id] = source
}

// NewExtractor returns the extractor object.
func NewExtractor(
	phoneValidator, emailValidator, vkValidator validator,
) *Extractor {
	e := &Extractor{
		merger: newMerger(map[string]validator{
			"email": emailValidator,
			"phone": phoneValidator,
			"vk":    vkValidator,
		}),
		sources: map[string]sourceExtractor{},
	}
	// e.AddSource("yeahdesk", extractors.NewYeahdesk())
	e.AddSource("email", extractors.NewEmailsExtractor())

	return e
}
