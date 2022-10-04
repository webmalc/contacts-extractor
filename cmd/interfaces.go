package cmd

// ErrorLogger logs errors.
type ErrorLogger interface {
	Error(args ...interface{})
}

// Runner runs the command.
type Runner interface {
	Run(names []string)
}

// Extractor extracts contacts.
type Extractor interface {
	Extract(sources []string, contacts []string) map[string][]string
}

// Formatter format the results.
type Formatter interface {
	Format(results map[string][]string) string
}
