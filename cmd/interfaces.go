package cmd

// errorLogger logs errors.
type errorLogger interface {
	Error(args ...interface{})
}

// extractor extracts contacts.
type extractor interface {
	Extract(sources []string, contacts []string) map[string][]string
}

// formatter format the results.
type formatter interface {
	Format(results map[string][]string) string
}
