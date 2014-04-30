package smtpapi

import (
	"encoding/json"
	"fmt"
)

// SMTPAPIHeader will be used to set up X-SMTPAPI params
type SMTPAPIHeader struct {
	To          []string            `json:"to,omitempty"`
	Sub         map[string][]string `json:"sub,omitempty"`
	Section     map[string]string   `json:"section,omitempty"`
	Category    []string            `json:"category,omitempty"`
	Unique_args map[string]string   `json:"unique_args,omitempty"`
	Filters     map[string]Filter   `json:"filters,omitempty"`
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

func (h *SMTPAPIHeader) SetTos(emails []string) {
	h.To = emails
}

func (h *SMTPAPIHeader) AddSubstitution(key, sub string) {
	if h.Sub == nil {
		h.Sub = make(map[string][]string)
	}
	h.Sub[key] = append(h.Sub[key], sub)
}

func (h *SMTPAPIHeader) SetSubstitutions(sub interface{}) error {
	var e bool
	if h.Sub, e = sub.(map[string][]string); e {
		return fmt.Errorf("smtpapi.go error: SetSubstitutions failed")
	} else {
		return nil
	}
}

func (h *SMTPAPIHeader) AddSection(section, value string) {
	if h.Section == nil {
		h.Section = make(map[string]string)
	}
	h.Section[section] = value
}

func (h *SMTPAPIHeader) SetSections(sections interface{}) error {
	var e bool
	if h.Section, e = sections.(map[string]string); e {
		return fmt.Errorf("smtpapi.go error: SetSections failed")
	} else {
		return nil
	}
}

func (h *SMTPAPIHeader) AddCategory(value string) {
	h.Category = append(h.Category, value)
}

func (h *SMTPAPIHeader) SetCategories(categories interface{}) error {
	var e bool
	if h.Category, e = categories.([]string); e {
		return fmt.Errorf("smtpapi.go ")
	} else {
		return nil
	}
}

func (h *SMTPAPIHeader) AddUniqueArg(arg, value string) {
	if h.Unique_args == nil {
		h.Unique_args = make(map[string]string)
	}
	h.Unique_args[arg] = value
}

func (h *SMTPAPIHeader) SetUniqueArgs(unique interface{}) error {
	var e bool
	if h.Unique_args, e = unique.(map[string]string); e {
		return fmt.Errorf("smtpapi.go error: SetUniqueArgs failed")
	} else {
		return nil
	}
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
