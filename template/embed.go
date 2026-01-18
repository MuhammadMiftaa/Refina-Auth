package template

import _ "embed"

//go:embed otp-email-template.html
var otpEmailTemplate string

var Template = map[string]string{
	"otp-email-template.html": otpEmailTemplate,
}
