package main

import (
  "github.com/sendgrid/smtpapi-go"
  "fmt"
)

func main() {
  header := smtpapi.NewSMTPAPIHeader()

  // Recipients
  header.AddTo("test@example.com")
  // or
  //tos := []string{"test1@example.com", "test2@example.com"}
  //header.AddTos(tos)
  // or
  //header.SetTos(tos)

  //[Substitutions](http://sendgrid.com/docs/API_Reference/SMTP_API/substitution_tags.html)
  header.AddSubstitution("key", "value")
  // or
  //values := []string{"value1", "value2"}
  //header.AddSubstitutions("key", values)
  // or
  //sub := make(map[string][]string)
  //sub["key"] = values
  //header.SetSubstitutions(sub)

  //[Section](http://sendgrid.com/docs/API_Reference/SMTP_API/section_tags.html)
  header.AddSection("section", "value")
  // or
  //sections := make(map[string]string)
  //sections["section"] = "value"
  //header.SetSections(sections)

  //[Category](http://sendgrid.com/docs/Delivery_Metrics/categories.html)
  header.AddCategory("category")
  // or
  //categories := []string{"setCategories"}
  //header.AddCategories(categories)
  // or
  //header.SetCategories(categories)

  //[Unique Arguments](http://sendgrid.com/docs/API_Reference/SMTP_API/unique_arguments.html)
  header.AddUniqueArg("key", "value")
  // or
  //args := make(map[string]string)
  //args["key"] = "value"
  //header.SetUniqueArgs(args)

  //[Filters](http://sendgrid.com/docs/API_Reference/SMTP_API/apps.html)
  header.AddFilter("filter", "setting", "value")
  // or
  /*
  filter := &Filter{
    Settings: make(map[string]string),
  }
  filter.Settings["enable"] = "1"
  filter.Settings["text/plain"] = "You can haz footers!"
  header.SetFilter("footer", filter)
  */

  //[Send At](https://sendgrid.com/docs/API_Reference/SMTP_API/scheduling_parameters.html)
  header.SetSendAt(1428611024)
  // or
  //sendEachAt := []int64{1428611024, 1428611025}
  //header.SetSendEachAt(sendEachAt)
  // or
  //header.AddSendEachAt(1428611024)
  //header.AddSendEachAt(1428611025)

  //[ASM Group ID](https://sendgrid.com/docs/User_Guide/advanced_suppression_manager.html)
  asmGroupID := 1
  header.SetASMGroupID(asmGroupID)

  // [IP Pools](https://sendgrid.com/docs/API_Reference/Web_API_v3/IP_Management/ip_pools.html)
  header.SetIpPool("testPool")

  fmt.Println(header.JSONString())
}