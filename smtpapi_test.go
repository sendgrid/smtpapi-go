package smtpapi

import (
	"encoding/json"
	"reflect"
	"testing"
)

func Test_JsonString(t *testing.T) {
  header := NewSMTPAPIHeader()

  result, _ := header.JsonString()

  if result != "{}" {
    t.Errorf("Result did not match")
  }
}

func Test_AddTo(t *testing.T) {
  header := NewSMTPAPIHeader()

  header.AddTo("addTo@mailinator.com")
  result, _ := header.JsonString()

  if result != "{\"to\":[\"addTo@mailinator.com\"]}" {
    t.Errorf("Result did not match")
  }
}

func Test_SetTos(t *testing.T) {
  header := NewSMTPAPIHeader()

  header.SetTos([]string{"setTos@mailinator.com"})
  result, _ := header.JsonString()

  if result != "{\"to\":[\"setTos@mailinator.com\"]}" {
    t.Errorf("Result did not match")
  }
}

func Test_AddSubstitution(t *testing.T) {
  header := NewSMTPAPIHeader()

  header.AddSubstitution("sub", "val")
  result, _ := header.JsonString()

  if result != "{\"sub\":{\"sub\":[\"val\"]}}" {
    t.Errorf("Result did not match")
  }
}

func Test_SetSubstitutions(t *testing.T) {
  header := NewSMTPAPIHeader()

	sub := make(map[string][]string)
	sub["sub"] = []string{"val"}

  header.SetSubstitutions(sub)
  result, _ := header.JsonString()

  if result != "{\"sub\":{\"sub\":[\"val\"]}}" {
    t.Errorf("Result did not match")
  }
}

func Test_AddUniqueArg(t *testing.T) {
  header := NewSMTPAPIHeader()

  header.AddUniqueArg("add_unique_argument_key", "add_unique_argument_value")
  header.AddUniqueArg("add_unique_argument_key_2", "add_unique_argument_value_2")
  result, _ := header.JsonString()

  if result != "{\"unique_args\":{\"add_unique_argument_key\":\"add_unique_argument_value\",\"add_unique_argument_key_2\":\"add_unique_argument_value_2\"}}" {
    t.Errorf("Result did not match")
  }
}

func Test_SetUniqueArgs(t *testing.T) {
  header := NewSMTPAPIHeader()

	args := make(map[string]string)
	args["set_unique_argument_key"] = "set_unique_argument_value"

  header.SetUniqueArgs(args)
  result, _ := header.JsonString()

  if result != "{\"unique_args\":{\"set_unique_argument_key\":\"set_unique_argument_value\"}}" {
    t.Errorf("Result did not match")
  }
}

func Test_AddCategory(t *testing.T) {
  header := NewSMTPAPIHeader()

  header.AddCategory("addCategory")
  header.AddCategory("addCategory2")
  result, _ := header.JsonString()

  if result != "{\"category\":[\"addCategory\",\"addCategory2\"]}" {
    t.Errorf("Result did not match")
  }
}

func Test_SetCategories(t *testing.T) {
  header := NewSMTPAPIHeader()

  header.SetCategories([]string{"setCategories"})
  result, _ := header.JsonString()

  if result != "{\"category\":[\"setCategories\"]}" {
    t.Errorf("Result did not match")
  }
}

func Test_AddSection(t *testing.T) {
  header := NewSMTPAPIHeader()

  header.AddSection("set_section_key", "set_section_value")
  header.AddSection("set_section_key_2", "set_section_value_2")
  result, _ := header.JsonString()

  if result != "{\"section\":{\"set_section_key\":\"set_section_value\",\"set_section_key_2\":\"set_section_value_2\"}}" {
    t.Errorf("Result did not match")
  }
}

func Test_SetSections(t *testing.T) {
  header := NewSMTPAPIHeader()

	sections := make(map[string]string)
	sections["set_section_key"] = "set_section_value"

  header.SetSections(sections)
  result, _ := header.JsonString()

  if result != "{\"section\":{\"set_section_key\":\"set_section_value\"}}" {
    t.Errorf("Result did not match")
  }
}

//func Test_AddFilter(t *testing.T) {
//  header := NewSMTPAPIHeader()
//
//  header.AddFilter("footer", "text/html", "<strong>boo</strong>")
//  result, _ := header.JsonString()
//
//  t.Logf(result);
//  if result != "{\"filters\":{\"footer\":{\"settings\":{\"text/html\":\"<strong>boo</strong>\"}}}}" {
//    t.Errorf("Result did not match")
//  }
//}
//
//func Test_SetFilters(t *testing.T) {
//  header := NewSMTPAPIHeader()
//
//  header.SetFilters("footer", "text/html", "<strong>boo</strong>")
//  result, _ := header.JsonString()
//
//  t.Logf(result);
//  if result != "{\"filters\":{\"footer\":{\"setting\":{\"enable\":1,\"text/plain\":\"You can haz footers!\"}}}}" {
//    t.Errorf("Result did not match")
//  }
//}

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
