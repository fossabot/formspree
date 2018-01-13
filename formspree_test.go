package formspree_test

import (
	"fmt"
	"testing"

	"github.com/rs/xid"

	"github.com/hourglassdesign/formspree"
	"github.com/stretchr/testify/assert"
)

func TestFormspree_Form(t *testing.T) {
	form := formspree.Form{}

	form.SetCC("test@example.com")
	form.SetEmail("test@example.com")
	form.SetFormat("plain")
	form.SetSubject("email subject")
	form.Set("name", "testing")

	assert.Equal(t, "test@example.com", form["_cc"])
	assert.Equal(t, "email subject", form["_subject"])
	assert.Equal(t, "plain", form["_format"])
	assert.Equal(t, "test@example.com", form["_replyto"])
	assert.Equal(t, "testing", form["name"])
}

func TestFormspree_Submit(t *testing.T) {
	tt := []struct {
		Email         string
		Format        string
		Subject       string
		CC            string
		User          string
		Referer       string
		ExpectedError string
		Data          map[string]interface{}
	}{
		{
			Email:   "test@example.com",
			Format:  "plain",
			Subject: "test email",
			CC:      "test@example.com",
			User:    fmt.Sprintf("%v@example.com", xid.New().String()),
			Referer: "testing",
			Data:    map[string]interface{}{"name": "testing"},
		},
		{
			User:          "invalid-email",
			Referer:       "testing",
			ExpectedError: "invalid email address",
			Data:          map[string]interface{}{"name": "testing"},
		},
		{
			ExpectedError: "json: unsupported type: chan int",
			Data:          map[string]interface{}{"name": make(chan int)},
		},
	}

	for _, tc := range tt {
		form := formspree.Form{}

		form.SetCC(tc.CC)
		form.SetEmail(tc.Email)
		form.SetFormat(tc.Format)
		form.SetSubject(tc.Format)

		for key, value := range tc.Data {
			form.Set(key, value)
		}

		if err := formspree.Submit(form, tc.Referer, tc.User); err != nil {
			assert.Equal(t, tc.ExpectedError, err.Error())
		}
	}
}
