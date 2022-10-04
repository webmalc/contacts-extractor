package formatters

import (
	"encoding/csv"
	"fmt"
	"strings"
)

// CSV is a csv formatter.
type CSV struct{}

// Format formats the results.
func (c *CSV) Format(results map[string][]string) string {
	b := new(strings.Builder)
	w := csv.NewWriter(b)
	for id, contacts := range results {
		if len(results) > 1 {
			b.WriteString(fmt.Sprintf("\n\n%s\n\n", id))
		}
		for _, contact := range contacts {
			if err := w.Write([]string{contact}); err != nil {
				panic(err)
			}
			w.Flush()
		}
	}
	w.Flush()

	return b.String()
}

// NewCSV returns the CSV object.
func NewCSV() *CSV {
	return &CSV{}
}
