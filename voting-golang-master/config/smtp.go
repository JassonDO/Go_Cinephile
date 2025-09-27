package config

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var (
	SEND_FROM_NAME    = "Voting App"
	SEND_FROM_ADDRESS = "garrydavid.gl@gmail.com"
	SENDGRID_API_KEY  = "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"
)

var SendGridClient *sendgrid.Client
var From *mail.Email

func init() {
	SendGridClient = sendgrid.NewSendClient(SENDGRID_API_KEY)
	From = mail.NewEmail(SEND_FROM_NAME, SEND_FROM_ADDRESS)
}
