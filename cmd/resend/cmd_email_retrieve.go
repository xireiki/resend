package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/resend/resend-go/v2"
	"github.com/spf13/cobra"
	"github.com/xireiki/resend/log"
	M "github.com/xireiki/resend/email"
)

var (
	emailId   string
	printJson bool
)

var emailRetrieveCommand = &cobra.Command{
	Use:   "retrieve",
	Short: "Get the specified mail by ID",
	Run:   func(cmd *cobra.Command, args []string) {
		if emailId == "" {
			log.Fatal("No email ID set, use \"--id\" to set")
		}
		client := resend.NewClient(key)
		email, err := M.RetrieveEmail(client, emailId)
		if err != nil {
			log.Error(err)
			return
		}
		if printJson {
			jsonData, err := json.MarshalIndent(email, "", "  ")
			if err != nil {
				log.Error("Error occurred while marshaling JSON: ", err)
				return
			}
			fmt.Println(string(jsonData))
			return
		}
		plainText := parseEmail(email)
		fmt.Println(plainText)
	},
}

func init() {
	emailCommand.AddCommand(emailRetrieveCommand)
	emailRetrieveCommand.PersistentFlags().StringVarP(&emailId, "id", "i", "", "Set email id")
	emailRetrieveCommand.PersistentFlags().BoolVarP(&printJson, "json", "j", false, "Output JSON text directly")
}

func parseEmail(email *resend.Email) string {
	var emailString string
	emailString = "From:    " + email.From
	emailString += "\nTo:      " + strings.Join(email.To, ", ")
	emailString += "\nSubject: " + email.Subject + "\n"
	var emailContext string
	if email.Text != "" {
		emailContext = email.Text
	}
	if email.Html != "" {
		emailContext = email.Html
	}
	if emailContext != "" {
		emailString += "\nText:\n\n" + email.Text
	}
	return emailString
}
