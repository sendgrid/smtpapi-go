package smtpapi

import "encoding/json"

// SMTPAPIHeader will be used to set up X-SMTPAPI params
type SMTPAPIHeader struct {
	To          []string                                `json:"to,omitempty"`
	Sub         map[string][]string                     `json:"sub,omitempty"`
	Section     map[string]string                       `json:"section,omitempty"`
	Category    []string                                `json:"category,omitempty"`
	Unique_args map[string]string                       `json:"unique_args,omitempty"`
	Filters     map[string]map[string]map[string]string `json:"filters,omitempty"`
}

func NewSMTPAPIHeader() SMTPAPIHeader {
	return SMTPAPIHeader{}
}

func (h *SMTPAPIHeader) AddTo(email string) {
	h.To = append(h.To, email)
}

func (h *SMTPAPIHeader) SetTos(emails []string) {
	h.To = emails
}

func (h *SMTPAPIHeader) AddSubstitution(key, sub string) {
	if h.Sub == nil {
		h.Sub = make(map[string][]string)
	}
	h.Sub[key] = append(h.Sub[key], sub)
}

func (h *SMTPAPIHeader) SetSubstitutions(sub interface{}) {
	h.Sub = sub
}

func (h *SMTPAPIHeader) AddSection(section, value string) {
	if h.Section == nil {
		h.Section = make(map[string]string)
	}
	h.Section[section] = value
}

func (h *SMTPAPIHeader) SetSections(sections interface{}) {
	h.Section = sections
}

func (h *SMTPAPIHeader) AddCategory(value string) {
	h.Category = append(h.Category, value)
}

func (h *SMTPAPIHeader) SetCategories(categories interface{}) {
	h.Category = categories
}

func (h *SMTPAPIHeader) AddUniqueArg(arg, value string) {
	if h.Unique_args == nil {
		h.Unique_args = make(map[string]string)
	}
	h.Unique_args[arg] = value
}

func (h *SMTPAPIHeader) SetUniqueArgs(unique interface{}) {
	h.Unique_args = unique
}

func (h *SMTPAPIHeader) AddFilter(filter, setting, value string) {
	if h.Filters == nil {
		h.Filters = make(map[string]map[string]map[string]string)
	}
	if h.Filters[filter] == nil {
		h.Filters[filter] = make(map[string]map[string]string)
	}
	if h.Filters[filter]["settings"] == nil {
		h.Filters[filter]["settings"] = make(map[string]string)
	}
	h.Filters[filter]["settings"][setting] = value
}

func (h *SMTPAPIHeader) JsonString() (string, error) {
	headers, e := json.Marshal(h)
	return string(headers), e
}
