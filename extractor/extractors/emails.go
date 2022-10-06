package extractors

import (
	"bufio"
	"os"
	"regexp"
	"strings"

	"github.com/microcosm-cc/bluemonday"
	"golang.org/x/exp/slices"
)

// EmailsExtractor is emails extractor.
type EmailsExtractor struct {
	config           *Config
	excludeFilenames []string
}

func (s *EmailsExtractor) processFile(filename string) []string {
	policy := bluemonday.StrictPolicy()
	replacer := strings.NewReplacer("(", "", ")", "", "-", "")
	results := []string{}

	f, err := os.Open(s.config.EmailsPath + filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewReader(f)
	for {
		bytes, _, err := scanner.ReadLine()
		emailRegexp := regexp.MustCompile(`[78]\d{10}`)
		text := replacer.Replace(policy.Sanitize(string(bytes)))
		all := emailRegexp.FindAllString(text, -1)
		results = append(results, all...)

		if err != nil {
			break
		}
	}

	return results
}

// Extract extracts contacts.
func (s *EmailsExtractor) Extract(contacts []string) map[string][]string {
	results := map[string][]string{}
	if !slices.Contains(contacts, "phone") {
		return results
	}

	files, err := os.ReadDir(s.config.EmailsPath)
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		if slices.Contains(s.excludeFilenames, f.Name()) {
			continue
		}
		results["phone"] = append(results["phone"], s.processFile(f.Name())...)
	}

	return results
}

// NewEmails returns the object.
func NewEmailsExtractor() *EmailsExtractor {
	return &EmailsExtractor{
		config:           NewConfig(),
		excludeFilenames: []string{".gitkeep"},
	}
}
