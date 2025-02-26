package types

import (
	"io/ioutil"
	"path/filepath"

	"github.com/resend/resend-go/v2"
)

type Listable[T any] []T

type EmailEntity struct {
	From        string
	To          Listable[string]
	Subject     string
	BCC         Listable[string]
	CC          Listable[string]
	ScheduledAt string
	ReplyTo     Listable[string]
	Content     string
	UseHTML     bool
	Files       Listable[string]
}

func Email2Request(email *EmailEntity) (*resend.SendEmailRequest, error) {
	params := &resend.SendEmailRequest{
		From:    email.From,
		To:      email.To,
		Subject: email.Subject,
	}
	if len(email.BCC) > 0 {
		params.Bcc = email.BCC
	}
	if len(email.CC) > 0 {
		params.Cc = email.CC
	}
	if email.ScheduledAt != "" {
		params.ScheduledAt = email.ScheduledAt
	}
	if len(email.ReplyTo) > 0 {
		params.ReplyTo = email.ReplyTo[0]
	}
	if email.UseHTML {
		params.Html = email.Content
	} else {
		params.Text = email.Content
	}
	if len(email.Files) > 0 {
		for _, v := range email.Files {
			name := filepath.Base(v)
			file, err := ioutil.ReadFile(v)
			if err != nil {
				return nil, err
			}
			params.Attachments = append(params.Attachments, &resend.Attachment{
				Filename: name,
				Content:  file,
			})
		}
	}
	return params, nil
}

func Emails2Request(emails []*EmailEntity) ([]*resend.SendEmailRequest, error) {
	batchEmails := []*resend.SendEmailRequest{}
	for _, v := range emails {
		request, err := Email2Request(v)
		if err != nil {
			return nil, err
		}
		batchEmails = append(batchEmails, request)
	}
	return batchEmails, nil
}
