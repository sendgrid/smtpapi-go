package smtpapi

import (
	"encoding/json"
	"reflect"
	"testing"
)

var testEmail = "email@email.com"

func TestNewSMTPIAPIHeader(t *testing.T) {
	headers := NewSMTPAPIHeader()
	if headers == nil {
		t.Error("NewSMTPAPIHeader() should never return nil")
	}
}

func TestAddTo(t *testing.T) {
	headers := NewSMTPAPIHeader()
	headers.AddTo(testEmail)
	if len(headers.To) == 0 {
		t.Errorf("AddTo Failed - %v", headers.To)
	}
}

func TestAddTos(t *testing.T) {
	headers := NewSMTPAPIHeader()
	tos := []string{testEmail, testEmail}
	headers.AddTos(tos)
	if len(headers.To) == 0 {
		t.Errorf("AddTos Failed - %v", headers.To)
	}
}

func TestSetTos(t *testing.T) {
	headers := NewSMTPAPIHeader()
	tos := []string{testEmail, testEmail}
	headers.SetTos(tos)
	if len(headers.To) == 0 {
		t.Errorf("SetTos Failed - %v", headers.To)
	}
}

func TestAddSubstitution(t *testing.T) {
	headers := NewSMTPAPIHeader()
	headers.AddSubstitution("sub", "val")
	if len(headers.Sub) == 0 {
		t.Errorf("AddSubstitution Failed - %v", headers.Sub)
	}
}

func TestAddSubstitutions(t *testing.T) {
	headers := NewSMTPAPIHeader()
	headers.AddSubstitutions("sub", []string{"val", "val2"})
	if len(headers.Sub) == 0 && len(headers.Sub["sub"]) != 2 {
		t.Errorf("AddSubstitutions Failed - %v", headers.Sub)
	}
}

func TestSetSubstitutions(t *testing.T) {
	headers := NewSMTPAPIHeader()
	sub := make(map[string][]string)
	sub["sub"] = []string{"val", "val2"}
	headers.SetSubstitutions(sub)
	if len(headers.Sub) == 0 && len(headers.Sub["sub"]) != 2 {
		t.Errorf("SetSubstitutions Failed - %v", headers.Sub)
	}
}

func Test_Adds(t *testing.T) {
	validHeader, _ := json.Marshal([]byte(`{"to":["test@email.com"],"sub":{"subKey":["subValue"]},"section":{"testSection":"sectionValue"},"category":["testCategory"],"unique_args":{"testUnique":"uniqueValue"},"filters":{"testFilter":{"settings":{"filter":"filterValue"}}}}`))
	headers := NewSMTPAPIHeader()
	headers.AddTo("test@email.com")
	headers.AddSubstitution("subKey", "subValue")
	headers.AddSection("testSection", "sectionValue")
	headers.AddCategory("testCategory")
	headers.AddUniqueArg("testUnique", "uniqueValue")
	headers.AddFilter("testFilter", "filter", "filterValue")
	if h, e := headers.JsonString(); e != nil {
		t.Errorf("Error! %s", e)
	} else {
		testHeader, _ := json.Marshal([]byte(h))
		if reflect.DeepEqual(testHeader, validHeader) {
			t.Logf("Success")
		} else {
			t.Errorf("Invalid headers")
		}
	}
}

func Test_Sets(t *testing.T) {
	validHeader, _ := json.Marshal([]byte(`{"to":["test@email.com"],"sub":{"subKey":["subValue"]},"section":{"testSection":"sectionValue"},"category":["testCategory"],"unique_args":{"testUnique":"uniqueValue"},"filters":{"testFilter":{"settings":{"filter":"filterValue"}}}}`))
	headers := NewSMTPAPIHeader()
	headers.SetTos([]string{"test@email.com"})
	sub := make(map[string][]string)
	sub["subKey"] = []string{"subValue"}
	headers.SetSubstitutions(sub)
	sections := make(map[string]string)
	sections["testSection"] = "sectionValue"
	headers.SetSections(sections)
	headers.SetCategories([]string{"testCategory"})
	unique := make(map[string]string)
	unique["testUnique"] = "uniqueValue"
	headers.SetUniqueArgs(unique)
	headers.AddFilter("testFilter", "filter", "filterValue")
	if h, e := headers.JsonString(); e != nil {
		t.Errorf("Error! %s", e)
	} else {
		testHeader, _ := json.Marshal([]byte(h))
		if reflect.DeepEqual(testHeader, validHeader) {
			t.Logf("Success")
		} else {
			t.Errorf("Invalid headers")
		}
	}
}
