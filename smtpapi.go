package smtpapi

import (
	"encoding/json"
)

// SMTPAPIHeader will be used to set up X-SMTPAPI params
type SMTPAPIHeader struct {
	To         []string            `json:"to,omitempty"`
	Sub        map[string][]string `json:"sub,omitempty"`
	Section    map[string]string   `json:"section,omitempty"`
	Category   []string            `json:"category,omitempty"`
	UniqueArgs map[string]string   `json:"unique_args,omitempty"`
	Filters    map[string]Filter   `json:"filters,omitempty"`
}

type Filter struct {
	Settings map[string]string `json:"settings,omitempty"`
}

func NewSMTPAPIHeader() *SMTPAPIHeader {
	return &SMTPAPIHeader{}
}

func (h *SMTPAPIHeader) AddTo(email string) {
	h.To = append(h.To, email)
}

func (h *SMTPAPIHeader) AddTos(emails []string) {
	for i := 0; i < len(emails); i++ {
		h.AddTo(emails[i])
	}
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

func (h *SMTPAPIHeader) AddSubstitutions(key string, subs []string) {
	for i := 0; i < len(subs); i++ {
		h.AddSubstitution(key, subs[i])
	}
}

func (h *SMTPAPIHeader) SetSubstitutions(sub map[string][]string) {
	h.Sub = sub
}

func (h *SMTPAPIHeader) AddSection(section, value string) {
	if h.Section == nil {
		h.Section = make(map[string]string)
	}
	h.Section[section] = value
}

func (h *SMTPAPIHeader) SetSections(sections map[string]string) {
	h.Section = sections
}

func (h *SMTPAPIHeader) AddCategory(category string) {
	h.Category = append(h.Category, category)
}

func (h *SMTPAPIHeader) AddCategories(categories []string) {
	for i := 0; i < len(categories); i++ {
		h.AddCategory(categories[i])
	}
}

func (h *SMTPAPIHeader) SetCategories(categories []string) {
	h.Category = categories
}

func (h *SMTPAPIHeader) AddUniqueArg(arg, value string) {
	if h.UniqueArgs == nil {
		h.UniqueArgs = make(map[string]string)
	}
	h.UniqueArgs[arg] = value
}

func (h *SMTPAPIHeader) SetUniqueArgs(args map[string]string) {
	h.UniqueArgs = args
}

func (h *SMTPAPIHeader) AddFilter(filter, setting, value string) {
	if h.Filters == nil {
		h.Filters = make(map[string]Filter)
	}
	if _, ok := h.Filters[filter]; !ok {
		h.Filters[filter] = Filter{
			Settings: make(map[string]string),
		}
	}
	h.Filters[filter].Settings[setting] = value
}

func (h *SMTPAPIHeader) SetFilter(filter string, value *Filter) {
	if h.Filters == nil {
		h.Filters = make(map[string]Filter)
	}
	h.Filters[filter] = *value
}

func (h *SMTPAPIHeader) JSONString() (string, error) {
	headers, e := json.Marshal(h)
	return string(headers), e
}
