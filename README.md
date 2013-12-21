# SMTP-API

This is a simple library to simplify the process of using [SendGrid's](https://sendgrid.com)[X-SMTPAPI](http://sendgrid.com/docs/API_Reference/SMTP_API/index.html).

# Examples

### [Substitution](http://sendgrid.com/docs/API_Reference/SMTP_API/substitution_tags.html)

```Go
header := smtpapi.NewSMTPAPIHeader()

header.AddSubstitution("key", "value")
```

### [Section](http://sendgrid.com/docs/API_Reference/SMTP_API/section_tags.html)

```Go
header := smtpapi.NewSMTPAPIHeader()

header.AddSection("section", "value")
```

### [Category](http://sendgrid.com/docs/Delivery_Metrics/categories.html)

```Go
header := smtpapi.NewSMTPAPIHeader()

header.AddSubstitution("category")
```

### [Unique Arguments](http://sendgrid.com/docs/API_Reference/SMTP_API/unique_arguments.html)

```Go
header := smtpapi.NewSMTPAPIHeader()

header.AddUniqueArg("key", "value")
```

### [Filter](http://sendgrid.com/docs/API_Reference/SMTP_API/apps.html)

```Go
header := smtpapi.NewSMTPAPIHeader()

header.AddFilter("filter", "setting", "value")
```

### GetHeaders

```Go
header.GetHeaders() //returns a JSON string representation of the headers
```

## Todo

Write tests and create Set Methods.

## MIT License
