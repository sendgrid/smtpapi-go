[![BuildStatus](https://travis-ci.org/sendgrid/smtpapi-go.svg?branch=master)](https://travis-ci.org/sendgrid/smtpapi-go)


**This is a simple library to simplify the process of using [SendGrid's](https://sendgrid.com) [X-SMTPAPI](http://sendgrid.com/docs/API_Reference/SMTP_API/index.html).**

# Announcements

All updates to this library is documented in our [CHANGELOG](https://github.com/sendgrid/smtpapi-go/blob/master/CHANGELOG.md).

# Installation

## Prerequisites

- Go version 1.6
- The SendGrid service, starting at the [free level](https://sendgrid.com/free?source=smtpapi-go)

## Install Package

```bash
go get github.com/sendgrid/smtpapi/go
```

# Quick Start

```go
package main

import (
  "github.com/sendgrid/smtpapi-go"
  "fmt"
)

func main() {
  header := smtpapi.NewSMTPAPIHeader()
  header.AddTo("test@example.com")
  fmt.Println(header.JSONString())
}
```

# Usage

- [SendGrid Docs](https://sendgrid.com/docs/API_Reference/SMTP_API/index.html)
- [Example Code](https://github.com/sendgrid/smtpapi-go/tree/master/examples)

## Roadmap

If you are interested in the future direction of this project, please take a look at our [milestones](https://github.com/sendgrid/smtpapi-go/milestones). We would love to hear your feedback.

## How to Contribute

We encourage contribution to our libraries, please see our [CONTRIBUTING](https://github.com/sendgrid/smtpapi-go/blob/master/CONTRIBUTING.md) guide for details.

Quick links:

- [Feature Request](https://github.com/sendgrid/smtpapi-go/blob/master/CONTRIBUTING.md#feature_request)
- [Bug Reports](https://github.com/sendgrid/smtpapi-go/blob/master/CONTRIBUTING.md#submit_a_bug_report)
- [Sign the CLA to Create a Pull Request](https://github.com/sendgrid/smtpapi-go/blob/master/CONTRIBUTING.md#cla)
- [Improvements to the Codebase](https://github.com/sendgrid/smtpapi-go/blob/master/CONTRIBUTING.md#improvements_to_the_codebase)

# About

smtpapi-go is guided and supported by the SendGrid [Developer Experience Team](mailto:dx@sendgrid.com).

smtpapi-go is maintained and funded by SendGrid, Inc. The names and logos for smtpapi-go are trademarks of SendGrid, Inc.

![SendGrid Logo]
(https://uiux.s3.amazonaws.com/2016-logos/email-logo%402x.png)
