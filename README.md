# formspree
[![CircleCI](https://img.shields.io/circleci/project/github/hourglassdesign/formspree.svg)](https://circleci.com/gh/hourglassdesign/formspree)
[![GoDoc](https://godoc.org/github.com/hourglassdesign/formspree?status.svg)](http://godoc.org/github.com/hourglassdesign/formspree)
[![Go Report Card](https://goreportcard.com/badge/github.com/hourglassdesign/formspree)](https://goreportcard.com/report/github.com/hourglassdesign/formspree)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/hourglassdesign/formspree/release/LICENSE)

A simple golang client for submitting forms via formspree.io

## usage:
```go
    // Construct a form
    form := formspree.Form{}

    // Add attributes
    form.SetEmail("example@example.com")
    form.SetSubject("custom email subject")
    form.Set("custom attribute", "custom value")

    // Submit your form using your referer name 
    // and desired destination
    err := formspree.Submit(form, "referer-id", "me@example.com")
```
