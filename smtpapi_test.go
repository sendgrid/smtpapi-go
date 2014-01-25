package smtpapi

import (
	"encoding/json"
	"reflect"
	"testing"
)

func Test_Adds(t *testing.T) {
	validHeader, _ := json.Marshal([]byte(`{"to":["test@email.com"],
	"sub":{"subKey":["subValue"]},
	"section":{"testSection":"sectionValue"},
	"category":["testCategory"],
	"unique_args":{"testUnique":"uniqueValue"},
	"filters":{"testFilter":{"settings":{"filter":"filterValue"}}}}`))

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
			t.Errorf("Invalid output")
		}
	}
}
