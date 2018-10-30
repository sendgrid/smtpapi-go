![SendGrid Logo](https://uiux.s3.amazonaws.com/2016-logos/email-logo%402x.png)

[![BuildStatus](https://travis-ci.org/sendgrid/smtpapi-go.svg?branch=master)](https://travis-ci.org/sendgrid/smtpapi-go)
[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE.txt)
[![Email Notifications Badge](https://dx.sendgrid.com/badge/php)](https://dx.sendgrid.com/newsletter/php)
[![Twitter Follow](https://img.shields.io/twitter/follow/sendgrid.svg?style=social&label=Follow)](https://twitter.com/sendgrid)
[![GitHub contributors](https://img.shields.io/github/contributors/sendgrid/smtpapi-go.svg)](https://github.com/sendgrid/smtpapi-go/graphs/contributors)
[![Go Report Card](https://goreportcard.com/badge/github.com/sendgrid/smtpapi-go)](https://goreportcard.com/report/github.com/sendgrid/smtpapi-go)
[![GoDoc](https://godoc.org/github.com/sendgrid/smtpapi-go?status.svg)](https://godoc.org/github.com/sendgrid/smtpapi-go)

**This is a simple library to simplify the process of using [SendGrid's](https://sendgrid.com) [X-SMTPAPI](http://sendgrid.com/docs/API_Reference/SMTP_API/index.html) with the Go programming language**

# Table of Contents

* [Announcements](#announcements)
* [Installation](#installation)
* [Quick Start](#quick-start)
* [Usage](#usage)
* [Roadmap](#roadmap)
* [How to Contribute](#how-to-contribute)
* [About](#about)
* [License](#license)

# Announcements

All updates to this library is documented in our [CHANGELOG](https://github.com/sendgrid/smtpapi-go/blob/master/CHANGELOG.md).

If you're a software engineer who is passionate about #DeveloperExperience and/or #OpenSource, this is an incredible opportunity to join our #DX team as a Developer Experience Engineer and work with @thinkingserious and @aroach! Tell your friends :)

<a name="installation"></a>
# Installation

## Prerequisites

* Go version 1.6
* The SendGrid service, starting at the [free level](https://sendgrid.com/free?source=smtpapi-go)

## Install Package

```bash
go get github.com/sendgrid/smtpapi-go
```

## Setup Environment Variables

### Environment Variable

Update the development environment with your [SENDGRID_API_KEY](https://app.sendgrid.com/settings/api_keys), for example:

```bash
echo "export SENDGRID_API_KEY='YOUR_API_KEY'" > sendgrid.env
echo "sendgrid.env" >> .gitignore
source ./sendgrid.env
```

<a name="quick-start"></a>
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
<a name="usage"></a>
# Usage

* [SendGrid Docs](https://sendgrid.com/docs/API_Reference/SMTP_API/index.html)
* [Example Code](https://github.com/sendgrid/smtpapi-go/tree/master/examples)

# Roadmap

If you are interested in the future direction of this project, please take a look at our [milestones](https://github.com/sendgrid/smtpapi-go/milestones). We would love to hear your feedback.

# How to Contribute

We encourage contribution to our libraries, please see our [CONTRIBUTING](https://github.com/sendgrid/smtpapi-go/blob/master/CONTRIBUTING.md) guide for details.

Quick links:

* [Feature Request](https://github.com/sendgrid/smtpapi-go/blob/master/CONTRIBUTING.md#feature-request)
* [Bug Reports](https://github.com/sendgrid/smtpapi-go/blob/master/CONTRIBUTING.md#submit-a-bug-report)
* [Sign the CLA to Create a Pull Request](https://github.com/sendgrid/smtpapi-go/blob/master/CONTRIBUTING.md#cla)
* [Improvements to the Codebase](https://github.com/sendgrid/smtpapi-go/blob/master/CONTRIBUTING.md#improvements-to-the-codebase)

<a name="about"></a>
# About

smtpapi-go is guided and supported by the SendGrid [Developer Experience Team](mailto:dx@sendgrid.com).

smtpapi-go is maintained and funded by SendGrid, Inc. The names and logos for smtpapi-go are trademarks of SendGrid, Inc.

# License

[The MIT License (MIT)](LICENSE.txt)
