package formspree_test

import (
	"fmt"

	"github.com/hourglassdesign/formspree"
	"github.com/rs/xid"
)

var (
	user     = fmt.Sprintf("%s@example.com", xid.New())
	referrer = "1234"
)

func ExampleForm_Submit() {
	form := formspree.Form{}

	form.SetEmail("sender@example.com")
	form.SetFormat("plain")
	form.SetSubject("an example subject")

	err := formspree.Submit(form, referrer, user)

	if err != nil {
		// handle errors
	}
}
