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

func TestAdd(t *testing.T) {
	headers := NewSMTPAPIHeader()
	headers.AddTo(testEmail)
	if len(headers.To) == 0 {
		t.Errorf("AddTo Failed - %v", headers.To)
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
