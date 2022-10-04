package extractors

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"golang.org/x/exp/slices"
)

// yeahdeskPerson is the response schema from the remote server.
type yeahdeskPerson struct {
	Contacts []struct {
		Value  string
		Type   string
		Origin string
	}
}

// Yeahdesk is a source.
type Yeahdesk struct {
	config *Config
}

// getRemoteData gets the remote persons data.
func (s *Yeahdesk) getRemoteData() []yeahdeskPerson {
	client := &http.Client{Timeout: 10 * time.Second}
	url := fmt.Sprintf(
		"%s?token=%s&limit=1000000000000",
		s.config.YeahdeskURL,
		s.config.YeahdeskToken,
	)
	r, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	var persons []yeahdeskPerson

	err = json.Unmarshal(data, &persons)

	if err != nil {
		panic(err)
	}

	return persons
}

// Extract extracts contacts.
func (s *Yeahdesk) Extract(contacts []string) map[string][]string {
	persons := s.getRemoteData()
	results := map[string][]string{}

	for _, p := range persons {
		for _, c := range p.Contacts {
			if slices.Contains(contacts, "email") && c.Type == "mail" {
				results["email"] = append(
					results["email"], strings.TrimSpace(c.Value),
				)
			}
			if slices.Contains(contacts, "phone") && c.Type == "phone" {
				results["phone"] = append(
					results["phone"], strings.TrimSpace(c.Value),
				)
			}
			if slices.Contains(contacts, "vk") && c.Type == "id" && c.Origin == "vk" {
				results["vk"] = append(
					results["vk"], strings.TrimSpace(c.Value),
				)
			}
		}
	}

	return results
}

// NewYeahdesk returns the object.
func NewYeahdesk() *Yeahdesk {
	return &Yeahdesk{config: NewConfig()}
}
