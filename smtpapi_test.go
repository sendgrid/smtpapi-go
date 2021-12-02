package smtpapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"regexp"
	"testing"
	"time"
)

func exampleJson() map[string]interface{} {
	data, _ := ioutil.ReadFile("smtpapi_test_strings.json")
	var f interface{}
	json.Unmarshal(data, &f)
	json := f.(map[string]interface{})
	return json
}

func TestNewSMTPIAPIHeader(t *testing.T) {
	t.Parallel()
	header := NewSMTPAPIHeader()
	if header == nil {
		t.Error("NewSMTPAPIHeader() should never return nil")
	}
}

func TestAddTo(t *testing.T) {
	t.Parallel()
	header := NewSMTPAPIHeader()
	header.AddTo("addTo@mailinator.com")
	result, _ := header.JSONString()
	if result != exampleJson()["add_to"] {
		t.Errorf("Result did not match")
	}
}

func TestAddTos(t *testing.T) {
	t.Parallel()
	header := NewSMTPAPIHeader()
	tos := []string{"addTo@mailinator.com"}
	header.AddTos(tos)
	result, _ := header.JSONString()
	if result != exampleJson()["add_to"] {
		t.Errorf("Result did not match")
	}
}

func TestSetTos(t *testing.T) {
	t.Parallel()
	header := NewSMTPAPIHeader()
	header.SetTos([]string{"setTos@mailinator.com"})
	result, _ := header.JSONString()
	if result != exampleJson()["set_tos"] {
		t.Errorf("Result did not match")
	}
}

func TestAddSubstitution(t *testing.T) {
	t.Parallel()
	header := NewSMTPAPIHeader()
	header.AddSubstitution("sub", "val")
	result, _ := header.JSONString()
	if result != exampleJson()["add_substitution"] {
		t.Errorf("Result did not match")
	}
}

func TestAddSubstitutions(t *testing.T) {
	t.Parallel()
	header := NewSMTPAPIHeader()
	header.AddSubstitutions("sub", []string{"val"})
	result, _ := header.JSONString()
	if result != exampleJson()["add_substitution"] {
		t.Errorf("Result did not match")
	}
}

func TestSetSubstitutions(t *testing.T) {
	t.Parallel()
	header := NewSMTPAPIHeader()
	sub := make(map[string][]string)
	sub["sub"] = []string{"val"}
	header.SetSubstitutions(sub)
	result, _ := header.JSONString()
	if result != exampleJson()["set_substitutions"] {
		t.Errorf("Result did not match")
	}
}

func TestAddSection(t *testing.T) {
	t.Parallel()
	header := NewSMTPAPIHeader()
	header.AddSection("set_section_key", "set_section_value")
	header.AddSection("set_section_key_2", "set_section_value_2")
	result, _ := header.JSONString()
	if result != exampleJson()["add_section"] {
		t.Errorf("Result did not match")
	}
}

func TestSetSections(t *testing.T) {
	t.Parallel()
	header := NewSMTPAPIHeader()
	sections := make(map[string]string)
	sections["set_section_key"] = "set_section_value"
	header.SetSections(sections)
	result, _ := header.JSONString()
	if result != exampleJson()["set_sections"] {
		t.Errorf("Result did not match")
	}
}

func TestAddCategory(t *testing.T) {
	t.Parallel()
	header := NewSMTPAPIHeader()
	header.AddCategory("addCategory")
	header.AddCategory("addCategory2")
	result, _ := header.JSONString()
	if result != exampleJson()["add_category"] {
		t.Errorf("Result did not match")
	}
}

func TestAddCategoryUnicode(t *testing.T) {
	t.Parallel()
	header := NewSMTPAPIHeader()
	header.AddCategory("カテゴリUnicode")
	header.AddCategory("カテゴリ2Unicode")
	header.AddCategory("鼖")
	result, _ := header.JSONString()
	if result != exampleJson()["add_category_unicode"] {
		t.Errorf("Result did not match")
	}
}

func TestAddCategories(t *testing.T) {
	t.Parallel()
	header := NewSMTPAPIHeader()
	categories := []string{"addCategory", "addCategory2"}
	header.AddCategories(categories)
	result, _ := header.JSONString()
	if result != exampleJson()["add_category"] {
		t.Errorf("Result did not match")
	}
}

func TestSetCategories(t *testing.T) {
	t.Parallel()
	header := NewSMTPAPIHeader()
	header.SetCategories([]string{"setCategories"})
	result, _ := header.JSONString()
	if result != exampleJson()["set_categories"] {
		t.Errorf("Result did not match")
	}
}

func TestAddUniqueArg(t *testing.T) {
	t.Parallel()
	header := NewSMTPAPIHeader()
	header.AddUniqueArg("add_unique_argument_key", "add_unique_argument_value")
	header.AddUniqueArg("add_unique_argument_key_2", "add_unique_argument_value_2")
	result, _ := header.JSONString()
	if result != exampleJson()["add_unique_arg"] {
		t.Errorf("Result did not match")
	}
}

func TestSetUniqueArgs(t *testing.T) {
	t.Parallel()
	header := NewSMTPAPIHeader()
	args := make(map[string]string)
	args["set_unique_argument_key"] = "set_unique_argument_value"
	header.SetUniqueArgs(args)
	result, _ := header.JSONString()
	if result != exampleJson()["set_unique_args"] {
		t.Errorf("Result did not match")
	}
}

func TestAddFilter(t *testing.T) {
	t.Parallel()
	header := NewSMTPAPIHeader()
	header.AddFilter("footer", "text/html", "<strong>boo</strong>")
	if len(header.Filters) != 1 {
		t.Error("AddFilter failed")
	}
}

func TestSetFilter(t *testing.T) {
	t.Parallel()
	header := NewSMTPAPIHeader()
	filter := &Filter{
		Settings: make(map[string]interface{}),
	}
	filter.Settings["enable"] = 1
	filter.Settings["text/plain"] = "You can haz footers!"
	header.SetFilter("footer", filter)
	result, _ := header.JSONString()
	if result != exampleJson()["set_filters"] {
		t.Errorf("Result did not match")
	}
}

func TestSetSendAt(t *testing.T) {
	t.Parallel()
	header := NewSMTPAPIHeader()
	header.SetSendAt(1428611024)
	result, _ := header.JSONString()
	if result != exampleJson()["set_send_at"] {
		t.Errorf("Result did not match")
	}
}

func TestAddSendEachAt(t *testing.T) {
	t.Parallel()
	header := NewSMTPAPIHeader()
	header.AddSendEachAt(1428611024)
	header.AddSendEachAt(1428611025)
	result, _ := header.JSONString()
	if result != exampleJson()["add_send_each_at"] {
		t.Errorf("Result did not match")
	}
}

func TestSetSendEachAt(t *testing.T) {
	t.Parallel()
	header := NewSMTPAPIHeader()
	sendEachAt := []int64{1428611024, 1428611025}
	header.SetSendEachAt(sendEachAt)
	result, _ := header.JSONString()
	if result != exampleJson()["set_send_each_at"] {
		t.Errorf("Result did not match")
	}
}

func TestSetASMGroupID(t *testing.T) {
	t.Parallel()
	header := NewSMTPAPIHeader()
	header.SetASMGroupID(1)
	result, _ := header.JSONString()
	if result != exampleJson()["set_asm_group_id"] {
		t.Errorf("Result did not match")
	}
}

func TestSetIpPool(t *testing.T) {
	t.Parallel()
	header := NewSMTPAPIHeader()
	header.SetIpPool("testPool")
	result, _ := header.JSONString()
	if result != exampleJson()["set_ip_pool"] {
		t.Errorf("Result did not match")
	}
}

func TestSAddASMGroupToDisplay(t *testing.T) {
	t.Parallel()
	header := NewSMTPAPIHeader()
	header.AddASMGroupToDisplay(671332)
	result, _ := header.JSONString()
	if result != exampleJson()["add_asm_group"] {
		t.Errorf("Result did not match")
	}
}

func TestSAddASMGroupsToDisplay(t *testing.T) {
	t.Parallel()
	header := NewSMTPAPIHeader()
	header.AddASMGroupsToDisplay([]int{45, 23})
	result, _ := header.JSONString()
	if result != exampleJson()["add_asm_groups"] {
		t.Errorf("Result did not match")
	}
}

func TestJSONString(t *testing.T) {
	t.Parallel()
	header := NewSMTPAPIHeader()
	result, _ := header.JSONString()
	if result != exampleJson()["json_string"] {
		t.Errorf("Result did not match")
	}
}

func TestJSONStringWithAdds(t *testing.T) {
	t.Parallel()
	validHeader, _ := json.Marshal([]byte(`{"to":["test@email.com"],"sub":{"subKey":["subValue"]},"section":{"testSection":"sectionValue"},"category":["testCategory"],"unique_args":{"testUnique":"uniqueValue"},"filters":{"testFilter":{"settings":{"filter":"filterValue"}}}}`))
	header := NewSMTPAPIHeader()
	header.AddTo("test@email.com")
	header.AddSubstitution("subKey", "subValue")
	header.AddSection("testSection", "sectionValue")
	header.AddCategory("testCategory")
	header.AddUniqueArg("testUnique", "uniqueValue")
	header.AddFilter("testFilter", "filter", "filterValue")
	if h, e := header.JSONString(); e != nil {
		t.Errorf("Error! %s", e)
	} else {
		testHeader, _ := json.Marshal([]byte(h))
		if reflect.DeepEqual(testHeader, validHeader) {
			t.Logf("Success")
		} else {
			t.Errorf("Invalid header")
		}
	}
}

func TestJSONStringWithSets(t *testing.T) {
	t.Parallel()
	validHeader, _ := json.Marshal([]byte(`{"to":["test@email.com"],"sub":{"subKey":["subValue"]},"section":{"testSection":"sectionValue"},"category":["testCategory"],"unique_args":{"testUnique":"uniqueValue"},"filters":{"testFilter":{"settings":{"filter":"filterValue"}}},"asm_group_id":1,"ip_pool":"testPool"}`))
	header := NewSMTPAPIHeader()
	header.SetTos([]string{"test@email.com"})
	sub := make(map[string][]string)
	sub["subKey"] = []string{"subValue"}
	header.SetSubstitutions(sub)
	sections := make(map[string]string)
	sections["testSection"] = "sectionValue"
	header.SetSections(sections)
	header.SetCategories([]string{"testCategory"})
	unique := make(map[string]string)
	unique["testUnique"] = "uniqueValue"
	header.SetUniqueArgs(unique)
	header.AddFilter("testFilter", "filter", "filterValue")
	header.SetASMGroupID(1)
	header.SetIpPool("testPool")
	if h, e := header.JSONString(); e != nil {
		t.Errorf("Error! %s", e)
	} else {
		testHeader, _ := json.Marshal([]byte(h))
		if reflect.DeepEqual(testHeader, validHeader) {
			t.Logf("Success")
		} else {
			t.Errorf("Invalid header")
		}
	}
}

func TestMarshalUnmarshall(t *testing.T) {
	t.Parallel()
	header := NewSMTPAPIHeader()
	header.SetTos([]string{"test@email.com"})
	sub := make(map[string][]string)
	sub["subKey"] = []string{"subValue"}
	header.SetSubstitutions(sub)
	sections := make(map[string]string)
	sections["testSection"] = "sectionValue"
	header.SetSections(sections)
	header.SetCategories([]string{"testCategory"})
	unique := make(map[string]string)
	unique["testUnique"] = "uniqueValue"
	header.SetUniqueArgs(unique)
	header.AddFilter("testFilter", "filter", "filterValue")
	header.SetASMGroupID(1)
	header.SetIpPool("testPool")
	header.SetASMGroupsToDisplay([]int{32, 12})

	newHeader := NewSMTPAPIHeader()
	b, err := header.JSONString()
	if err != nil {
		t.Errorf("Error in JSONString %v", err)
	}
	err = newHeader.Load([]byte(b))
	if err != nil {
		t.Errorf("Could not load newHeader %v", err)
	}
	if !reflect.DeepEqual(header, newHeader) {
		t.Errorf("Expected %v, but got %v", header, newHeader)
	}
}

func TestRepoFiles(t *testing.T) {
	files := []string{".env_sample", ".gitignore", ".github/workflows/test.yml", "CHANGELOG.md", "CODE_OF_CONDUCT.md",
		"CONTRIBUTING.md", "ISSUE_TEMPLATE.md", "LICENSE", "PULL_REQUEST_TEMPLATE.md",
		"README.md", "TROUBLESHOOTING.md", "USAGE.md"}

	for _, file := range files {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			t.Errorf("Repo file does not exist: %v", file)
		}
	}
}

func TestLicenseYear(t *testing.T) {
	t.Parallel()
	dat, err := ioutil.ReadFile("LICENSE")

	currentYear := time.Now().Year()
	r := fmt.Sprintf("%d", currentYear)
	match, _ := regexp.MatchString(r, string(dat))

	if err != nil {
		t.Error("License File Not Found")
	}
	if !match {
		t.Error("Incorrect Year in License Copyright")
	}
}
