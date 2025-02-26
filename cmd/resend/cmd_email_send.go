package main

import (
	"fmt"

	"github.com/resend/resend-go/v2"
	"github.com/spf13/cobra"
	"github.com/xireiki/resend/log"
	M "github.com/xireiki/resend/email"
	"github.com/xireiki/resend/types"
)

var (
	from        string
	to          []string
	subject     string
	bcc         []string
	cc          []string
	reply_to    []string
	html        string
	text        string
	attachments []string
)

var emailSendCommand = &cobra.Command{
	Use:              "send",
	Short:            "Send an e-mail message",
	Run: func(cmd *cobra.Command, args []string) {
		if from == "" {
			log.Fatal("Missing from parameter")
		}
		if len(to) == 0 {
			log.Fatal("Missing to parameter(mail destination)")
		}
		if len(to) > 50 {
			log.Fatal("Too many mail destinations, it supports 50 addresses at most")
		}
		if subject == "" {
			log.Fatal("Confirm the subject of the message")
		}
		err := run()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	emailCommand.AddCommand(emailSendCommand)
	emailSendCommand.PersistentFlags().StringVarP(&from, "from", "f", "", "Set from email")
	emailSendCommand.PersistentFlags().StringArrayVarP(&to, "to", "t", nil, "Set to email")
	emailSendCommand.PersistentFlags().StringVarP(&subject, "subject", "s", "", "Set email subject")
	emailSendCommand.PersistentFlags().StringArrayVarP(&bcc, "bcc", "b", nil, "Set email bcc")
	emailSendCommand.PersistentFlags().StringArrayVarP(&cc, "cc", "c", nil, "Set email cc")
	emailSendCommand.PersistentFlags().StringArrayVarP(&reply_to, "reply-to", "r", nil, "Set reply_to email")
	emailSendCommand.PersistentFlags().StringVarP(&html, "html", "H", "", "Send a email in HTML format")
	emailSendCommand.PersistentFlags().StringVarP(&text, "text", "T", "", "Send email in plain text format")
	emailSendCommand.PersistentFlags().StringArrayVarP(&attachments, "attachments", "a", nil, "Add Attachment")
}

func run() error {
	client := resend.NewClient(key)
	email := &types.EmailEntity{
		From:      from,
		To:        to,
		Subject:   subject,
	}
	if len(bcc) > 0 {
		email.BCC = bcc
	}
	if len(cc) > 0 {
		email.CC = cc
	}
	if len(reply_to) > 0 {
		email.ReplyTo = reply_to
	}
	if text != "" {
		email.Content = text
	}
	if html != "" {
		email.Content = html
		email.UseHTML = true
	}
	if len(attachments) > 0 {
		email.Files = attachments
	}
	response, err := M.SendMail(globalCtx, client, email)
	if err != nil {
		return err
	}
	fmt.Println(response.Id)
	return nil
}
