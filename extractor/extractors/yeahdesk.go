package extractors

// Yeahdesk is a source.
type Yeahdesk struct{}

// Extract extracts contacts.
func (s *Yeahdesk) Extract(contacts []string) map[string][]string {
	return map[string][]string{
		"phone": {"+79253172878", "222"},
		"email": {"sd", "adf", "m@webmalc.pw"},
	}
}

// NewYeahdesk returns the object.
func NewYeahdesk() *Yeahdesk {
	return &Yeahdesk{}
}
