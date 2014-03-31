# SMTP-API

This is a simple library to simplify the process of using [SendGrid's](https://sendgrid.com) [X-SMTPAPI](http://sendgrid.com/docs/API_Reference/SMTP_API/index.html).

[![BuildStatus](https://travis-ci.org/sendgrid/sendgrid-go.png?branch=master)](https://travis-ci.org/sendgrid/sendgrid-go)

## Examples

### [AddSubstitution](http://sendgrid.com/docs/API_Reference/SMTP_API/substitution_tags.html)

```Go
header := smtpapi.NewSMTPAPIHeader()

header.AddSubstitution("key", "value")
```

### SetSubsitutions

```Go
header := smtpapi.NewSMTPAPIHeader()

sub := make(map[string][]string)
sub["sub"] = []string{"val"}

header.SetSubstitutions(sub)
```

### [Section](http://sendgrid.com/docs/API_Reference/SMTP_API/section_tags.html)

```Go
header := smtpapi.NewSMTPAPIHeader()

header.AddSection("section", "value")
```

### SetSections

```Go
header := smtpapi.NewSMTPAPIHeader()

sections := make(map[string]string)
sections["set_section_key"] = "set_section_value"

header.SetSections(sections)
```

### [Category](http://sendgrid.com/docs/Delivery_Metrics/categories.html)

```Go
header := smtpapi.NewSMTPAPIHeader()

header.AddCategory("category")
```

### SetCategories

```Go
header := smtpapi.NewSMTPAPIHeader()

header.SetCategories([]string{"setCategories"})
```

### [Unique Arguments](http://sendgrid.com/docs/API_Reference/SMTP_API/unique_arguments.html)

```Go
header := smtpapi.NewSMTPAPIHeader()

header.AddUniqueArg("key", "value")
```

### SetUniqueArgs

```Go
header := smtpapi.NewSMTPAPIHeader()

args := make(map[string]string)
args["set_unique_argument_key"] = "set_unique_argument_value"

header.SetUniqueArgs(args)
```

### [Filter](http://sendgrid.com/docs/API_Reference/SMTP_API/apps.html)

```Go
header := smtpapi.NewSMTPAPIHeader()

header.AddFilter("filter", "setting", "value")
```

### JsonString

```Go
header.JsonString() //returns a JSON string representation of the headers
```

## MIT License

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Added some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

## Running Tests

````bash
go test -v
```
