# Usage

## Initialisation

```go
package main

import (
    "fmt"
    "github.com/sendgrid/smtpapi-go"
)

header := smtpapi.NewSMTPAPIHeader()
```

## Table of Contents

- [Recipients](#recipients)
- [Categories](#categories)
- [Substitution](#substitution)
- [Sections](#sections)
- [Filters](#filters)
- [Advanced Suppression Manager (ASM)](#asm)
- [Send At](#send-at)
- [Batches](#batches)
- [IP Pools](#ip-pools)
- [Retrieving Data](#retrieving-data)

<a name="recipients"></a>
## Recipients

#### Adding a single recipient

The `header.AddTo` method allows you to add a new recipient to the `To` array.

To do this, you can use the function with an email address, and optionally include the recipient's name as a second parameter.

```go
header.SetTo("example@example.com") // An email address with no name provided
header.SetTo("example@example.com", "An Example User")  // An email address with a name provided as the second parameter
```

#### Setting multiple recipients

The `header.SetTos` method allows you to set multiple recipients by providing an array of email addresses.

This will set the `To` array to the array you provide. This will need to either be an array of emails, or if names are provided, they need to be formatted as `{{ name }} <{{ email }}>`.

```go
header.SetTos([]string{
    "example@example.com", // An email address with no name
    "An Example User <example@example.com>", // An email with a name field
})
```

<a name="categories"></a>
## Categories

Categories are a useful way to organise your email analytics by tagging your emails with a specific type or topic.

There are multiple methods available for setting your categories when sending emails.

#### Adding a single category

The `header.AddCategory` method can be used to add a category to your list.

```go
header.AddCategory('marketing') // Add a new category of 'marketing' to the array of categories
```

#### Setting multiple categories

The `header.SetCategories` method can be used to set your categories list to an array of strings.

This is useful when there are a set amount of categories required for the email you are sending.

This method will remove any categories that have been previously set.

```go
header.SetCategories([]string{
    "marketing",
    "sales",
}) // Sets the current categories to be 'marketing' and 'sales'
```

#### Setting a single category

The `header.SetCategory` method can be used to set a single specific category.

It is useful for removing previously set categories. And will create a new array with the string you provide.

This method will remove any categories that have been previously set.

```go
header.SetCategory("marketing") // Reset the categories to be 'marketing' only
```

<a name="substitution"></a>
## Substitution

Substitutions are a great way of writing short dynamic email content easily,

#### Adding a single substitution string

The `header.AddSubstitution` method can be used to replace content for recipients.

```go
header.AddSubstitution("-name-", "John") // Replace the -name- variable with John.
```

#### Setting substitution strings

The `header.SetSubstitutions` method can be used to replace content for any number of strings.

This method will reset any key pairs that have previously been set.

```go
header.SetSubstitutions(map[string][]string{
    "-name-":   {"John", "Jane"}, // Replace the -name- variable to John or Jane
    "-number-": {"555.555.5555", "777.777.7777"}, // Replace the -number- variable with the provided numbers
})
```

<a name="sections"></a>
## Sections

Sections are similar to substitutions, but are specific to the actual message rather than the recipient.

This is useful when you are sending multiple emails with the same style, but different content.

Note that substitution variables can also be included within a section, but section variables cannot

#### Adding a section

The `header.AddSection` method can be used to add a new section to the sections array. This is useful for building up a list of sections dynamically, perhaps based on a user's actions.

```go
header.AddSection("-event_details-", "The event will be held tomorrow.") // Replaces -event_details- with the event's string
```

#### Setting multiple sections

The `header.SetSections` allows you to set multiple sections in a single array.

This is good when sending out multiple emails where no dynamic variation is required.

This will reset any section key-pairs that have previously been set.

```go
header.SetSections(map[string]string{
    "-event_details-":   "The event will be held tomorrow.",
    "-event_open_time-": "'It will be open from 1am to 9pm.",
})
```

<a name="filters"></a>
## Filters

Filters allow you to dynamically toggle features such as click tracking, blind copying and DKIM domain validation.

#### Adding a single filter

Adding a filter is easy with the `header.AddFilter` method.

This method requires 3 values:
- The filter's name
- The parameter's name
- The value

```go
header.AddFilter("dkim", "use_from", true)
header.AddFilter("dkim", "domain", "example.com")
```

#### Adding pre-existing filters

Filters with predetermined settings can also be added using the `header.SetFilter` method.

```go
filter := &Filter{
    Settings: make(map[string]interface{}),
}
filter.Settings["enable"] = 1
filter.Settings["text/plain"] = "You can haz footers!"
header.SetFilter("footer", filter)
```

<a name="asm"></a>
## Advanced Suppression Management (ASM)

Advanced Suppression Management (or Unsubscribe Groups) are a good way of allowing recipients to unsubscribe from a specific set of emails.

You can

#### Setting the ASM group ID

The `header.SetASMGroupID` method is a quick way to set the type of email that you are sending.

All it requires is the ID of the ASM group, which can be found using the API.

```go
header.SetASMGroupID(42) // Sets the ASM ID to 42
```

<a name="send-at"></a>
## Send At

Scheduling the time of your email campaign can be done using a collection of quick and easy methods.

#### Adding a single 'send at' date

The `header.AddSendEachAt` method is a good way to add the time to send at.

This method requires a unix timestamp as the input.

```go
header.AddSendEachAt(1508694645)
```

#### Setting multiple 'send at' date

The `header.SetSendEachAt` method is useful for setting an array of times that recipients have their emails sent.

This method requires an array of unix timestamps as the input.

```go
header.SetSendEachAt([]int64{
    1508694645,
    1508694835,
})
```

#### Setting a single date to send all emails

The `header.SetSendAt` method is useful for setting a single time that all emails in the collection will be sent.

This method requires a unix timestamp as the input.

```go
header.SetSendAt(1508694645)
```

<a name="batches"></a>
## Batches

Batches are a great way to group a collection of scheduled items for sending. It allows you to cancel scheduled emails, and provides more control over the emails.

The batch ID can be set using the `header.AddBatchId` method. You must have generated the batch ID first through the API.

```go
header.AddBatchId("HkJ5yLYULb7Rj8GKSx7u025ouWVlMgAi") // Adds a previously generated batch ID to the emails
```

<a name="ip-pools"></a>
## IP Pools

IP Pools allow you to group SendGrid IP addresses together. For example, if you have a set of marketing IPs, you can assign them a pool ID of `marketing`.

The IP Pool name can be set using the `header.SetIpPool` method. You must have generated the IP Pool first through the API.

```go
header.SetIpPool("marketing") // Sets the IP Pool to be the marketing collection of IPs
```


<a name="retrieving-data"></a>
#### Retrieving data as a JSON string

The `header.JSONString` method allows the data from the header instance to be exported as a JSON string.

```go
headerString, _ := header.JSONString()
fmt.Println(headerString)

// {"to":["test@example.com"],"sub":{"key":["value"]},"section":{"section":"value"},"category":["category"],"unique_args":{"key":"value"},"filters":{"filter":{"settings":{"setting":"value"}}},"asm_group_id":1,"send_at":1428611024,"ip_pool":"testPool"}
```