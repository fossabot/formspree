package formspree

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type (
	// Form represents an HTML form that is to be submitted via Formspree
	Form map[string]interface{}

	// Response represents the HTTP response given by Formspree after submission.
	Response struct {
		Error string `json:"error"`
	}
)

const urlFormat = "https://formspree.io/%v"

// SetSubject is used to set the email's subject, so that you can quickly reply
// to submissions without having to edit the subject line each time.
func (f Form) SetSubject(subject string) {
	f["_subject"] = subject
}

// SetCC is used to set the email's CC Field. This lets you send a copy of
// each submission to another email address.
func (f Form) SetCC(cc string) {
	f["_cc"] = cc
}

// SetFormat sets the desired format for the resulting email. Adding this to
// your form will allow for you to receive plain text versions of emails
// for form submissions.
func (f Form) SetFormat(format string) {
	f["_format"] = format
}

// SetEmail sets the email's Reply-To field. This way you can directly "Reply"
// to the email to respond to the person who originally submitted the form.
func (f Form) SetEmail(email string) {
	f["_replyto"] = email
}

// Set adds an attribute to the form with a given value. Values must be JSON-encodable.
func (f Form) Set(key string, value interface{}) {
	f[key] = value
}

// Submit attempts to post the given form to Formspree.
func Submit(form Form, referer, user string) error {
	var client http.Client
	var response Response

	// Generate url.
	url := fmt.Sprintf(urlFormat, user)

	data, err := json.Marshal(form)

	if err != nil {
		return err
	}

	// Build request.
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))

	if err != nil {
		return err
	}

	req.Header.Add("content-type", "application/json")
	req.Header.Add("referer", referer)

	// Perform request.
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	// If successful, return nil.
	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated {
		return nil
	}

	decoder := json.NewDecoder(resp.Body)

	// Otherwise, parse error response.
	if err := decoder.Decode(&response); err != nil {
		return err
	}

	// Convert Formspree error into error type.
	msg := strings.ToLower(response.Error)
	return errors.New(msg)
}
