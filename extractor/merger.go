package extractor

import (
	mapset "github.com/deckarep/golang-set/v2"
)

// merger is a merger.
type merger struct {
	validators map[string]validator
}

// merge mergers results.
func (m *merger) merge(
	results, contacts map[string][]string,
) map[string][]string {
	for key, values := range contacts {
		unique := mapset.NewSet[string]()
		for _, c := range values {
			// unique.Add(c)
			if val, ok := m.validators[key]; !ok || (ok && val.Validate(c)) {
				unique.Add(c)
			}
		}
		for _, c := range results[key] {
			// unique.Add(c)
			if val, ok := m.validators[key]; !ok || (ok && val.Validate(c)) {
				unique.Add(c)
			}
		}
		results[key] = unique.ToSlice()
	}

	return results
}

// newMerger return a new merger.
func newMerger(validators map[string]validator) *merger {
	return &merger{
		validators: validators,
	}
}
