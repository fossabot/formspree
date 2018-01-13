# formspree
[![CircleCI](https://circleci.com/gh/hourglassdesign/formspree.svg?style=svg)](https://circleci.com/gh/hourglassdesign/formspree)
[![GoDoc](https://godoc.org/github.com/hourglassdesign/formspree?status.svg)](http://godoc.org/github.com/hourglassdesign/formspree)

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
